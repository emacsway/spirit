package spirit

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/go-spirit/spirit/cache"
	"github.com/go-spirit/spirit/component"
	"github.com/go-spirit/spirit/doc"
	"github.com/go-spirit/spirit/mail"
	"github.com/go-spirit/spirit/mail/dispatcher"
	"github.com/go-spirit/spirit/worker"
	"github.com/gogap/config"
	"github.com/gogap/logrus_mate"

	_ "github.com/go-spirit/spirit/cache/gocache"
	_ "github.com/go-spirit/spirit/mail/dispatcher/async/goroutine"
	_ "github.com/go-spirit/spirit/mail/dispatcher/sync/synchronized"
	_ "github.com/go-spirit/spirit/mail/mailbox"
	_ "github.com/go-spirit/spirit/mail/postman/tiny"
	_ "github.com/go-spirit/spirit/mail/registry/tiny"
	_ "github.com/go-spirit/spirit/worker/fbp"
)

var (
	ErrNameIsEmpty = errors.New("name param is empty")
)

type Spirit struct {
	loc sync.Mutex

	postman mail.Postman
	reg     mail.Registry

	workers         map[string]worker.Worker
	actors          map[string]*Actor
	actorStartOrder []string

	conf config.Configuration
}

func New(opts ...Option) (s *Spirit, err error) {

	spiritOpts := &Options{}

	for _, o := range opts {
		o(spiritOpts)
	}

	loggerConf := spiritOpts.config.GetConfig("logger")

	if loggerConf != nil {
		logrus_mate.Hijack(logrus.StandardLogger(), logrus_mate.WithConfig(loggerConf))
	}

	reg, err := mail.NewRegistry("tiny")
	if err != nil {
		return
	}

	man, err := mail.NewPostman("tiny", mail.PostmanRegistry(reg))
	if err != nil {
		return
	}

	sp := &Spirit{
		reg:     reg,
		postman: man,
		workers: make(map[string]worker.Worker),
		actors:  make(map[string]*Actor),
		conf:    spiritOpts.config,
	}

	err = sp.generateActors()
	if err != nil {
		return
	}

	s = sp

	return
}

func (p *Spirit) generateActors() (err error) {
	componentsConf := p.conf.GetConfig("components")
	if componentsConf == nil {
		return
	}

	drivers := componentsConf.Keys()

	if len(drivers) == 0 {
		return
	}

	for _, driver := range drivers {

		driverConf := componentsConf.GetConfig(driver)

		if driverConf == nil {
			continue
		}

		actorNames := driverConf.Keys()

		if len(actorNames) == 0 {
			continue
		}

		for _, actName := range actorNames {
			_, err = p.NewActor(
				actName,
				ActorComponent(driver, component.Config(
					driverConf.GetConfig(actName),
				)),
			)

			if err != nil {
				return
			}
		}
	}

	return
}

func (p *Spirit) Run() (err error) {

	for _, actUrl := range p.actorStartOrder {
		act := p.actors[actUrl]
		err = act.Start()
		if err != nil {
			return
		}
	}

	err = p.postman.Start()
	if err != nil {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case <-ch:
	}

	err = p.Stop()
	if err != nil {
		return
	}

	return
}

func (p *Spirit) Stop() error {
	for _, act := range p.actors {
		e := act.Stop()
		if e != nil {
			logrus.WithError(e).Errorln("stop actor")
		}
	}

	return nil
}

func (p *Spirit) WithPostman(pm mail.Postman) {
	p.postman = pm
}

func (p *Spirit) newWorker(name, driver string, opts ...WorkerOption) (wk worker.Worker, err error) {

	if len(name) == 0 {
		err = ErrNameIsEmpty
		return
	}

	key := fmt.Sprintf("workers.%s", name)

	workerOptions := WorkerOptions{}

	for _, o := range opts {
		o(&workerOptions)
	}

	newWk, err := worker.New(
		driver,
		worker.Postman(p.postman),
		worker.Router(workerOptions.HandlerRouter),
	)

	if err != nil {
		return
	}

	mailboxDriver := p.conf.GetString(key+"mail.mailbox.driver", "unbounded")

	dispatcherDriver := p.conf.GetString(key+"mail.dispatcher.driver", "goroutine")
	throughput := p.conf.GetInt32(key+"mail.dispatcher.throughput", 300)

	dispatcher, err := dispatcher.NewDispatcher(dispatcherDriver, int(throughput))

	if err != nil {
		return
	}

	box, err := mail.NewMailbox(
		mailboxDriver,
		mail.MailboxUrl(workerOptions.Url),
		mail.MailboxMessageInvoker(newWk),
		mail.MailboxDispatcher(dispatcher),
	)

	if err != nil {
		return
	}

	err = p.reg.Register(box)
	if err != nil {
		return
	}

	wk = newWk

	return
}

func (p *Spirit) NewActor(name string, opts ...ActorOption) (act *Actor, err error) {

	actOpts := ActorOptions{}

	for _, o := range opts {
		o(&actOpts)
	}

	if len(actOpts.componentDriver) == 0 {
		err = errors.New("component driver name is empty")
		return
	}

	if len(actOpts.workerDriver) == 0 {
		actOpts.workerDriver = "fbp"
		logrus.Debugln("there is no worker driver specificed, use default worker of fbp")
	}

	if len(actOpts.url) == 0 {
		actOpts.url = fmt.Sprintf("spirit://actors/%s/%s/%s", actOpts.workerDriver, actOpts.componentDriver, name)
	}

	_, exist := p.actors[actOpts.url]
	if exist {
		err = fmt.Errorf("actor already registerd, url: %s", actOpts.url)
		return
	}

	componentConfKey := fmt.Sprintf("components.%s.%s", actOpts.componentDriver, name)

	componentConf := p.conf.GetConfig(componentConfKey)
	if componentConf == nil {
		componentConf = config.NewConfig()
	}

	cacheConf := componentConf.GetConfig("caches")

	caches, err := cache.NewCaches(cacheConf)
	if err != nil {
		return
	}

	compOptions := []component.Option{
		component.Postman(p.postman),
		component.Caches(caches),
		component.Config(componentConf),
	}

	compOptions = append(compOptions, actOpts.componentOptions...)

	comp, err := component.NewComponent(
		actOpts.componentDriver,
		name,
		compOptions...,
	)

	if err != nil {
		return
	}

	warnNoDocsComp(name, actOpts.componentDriver, comp)

	worker, err := p.newWorker(
		name,
		actOpts.workerDriver,
		WorkerUrl(actOpts.url),
		WorkerHandlerRouter(comp),
	)

	if err != nil {
		return
	}

	act = &Actor{
		worker:    worker,
		component: comp,
	}

	p.actors[actOpts.url] = act
	p.actorStartOrder = append(p.actorStartOrder, actOpts.url)

	logrus.WithField("url", act.Url()).
		WithField("name", name).
		WithField("worker", actOpts.workerDriver).
		WithField("componet", actOpts.componentDriver).Debugln("actor registered")

	return
}

func warnNoDocsComp(name, driver string, comp interface{}) {
	if _, ok := comp.(doc.Documenter); !ok {
		if driver == "function" {
			if _, exist := doc.GetDocumenter(name); !exist {
				logrus.WithField("function", name).Warnln("no document implement")
			}
		} else {
			logrus.WithField("component-driver", driver).Warnln("no document implement")
		}
	} else if _, exist := doc.GetDocumenter(driver); !exist {
		logrus.WithField("component-driver", driver).Warnln("document implemented, but not registered")
	}
}
