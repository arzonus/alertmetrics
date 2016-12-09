package appc

import (
	c "github.com/arzonus/alertmetrics/pkg/checker"
	cc "github.com/arzonus/alertmetrics/pkg/checkerc"
	"github.com/arzonus/alertmetrics/pkg/config"
	"github.com/arzonus/alertmetrics/pkg/context"
	l "github.com/arzonus/alertmetrics/pkg/loader"
	m "github.com/arzonus/alertmetrics/pkg/model"
	"github.com/arzonus/alertmetrics/pkg/notifier"
	nc "github.com/arzonus/alertmetrics/pkg/notifierc"
)

type Application struct {
	cfg *config.Config
	ctx *context.Context

	loader    *l.Loader
	checkerc  *cc.CheckerController
	notifierc *nc.NotifierController

	checker   *c.Checker
	notifiers []nc.Notifier

	items  chan *m.Items
	report chan string
}

func New() *Application {
	return &Application{
		cfg: config.Get(),
		ctx: context.Get(),
	}
}

func (a *Application) Init() (err error) {

	if err = a.initChecker(); err != nil {
		return
	}

	if err = a.initNotifiers(); err != nil {
		return
	}

	a.items = make(chan *m.Items)
	a.report = make(chan string)

	a.loader = l.New(
		a.ctx.Database,
		a.items,
		a.cfg.Period,
		a.ctx.Storage,
		a.ctx.Logger,
	)

	a.checkerc = cc.New(
		a.report,
		a.items,
		a.checker,
	)

	a.notifierc = nc.New(
		a.report,
		a.notifiers,
	)

	return
}

func (a *Application) initNotifiers() error {
	if a.cfg.Notifier.LogNotifier.Enable {
		a.notifiers = append(a.notifiers, notifier.NewLogNotifier(a.ctx.Logger))
	}

	return nil
}

func (a *Application) initChecker() (err error) {
	var lowerBound = make(map[string]uint)
	var upperBound = make(map[string]uint)

	for _, metric := range a.cfg.Metrics {
		lowerBound[metric.Name] = metric.LowerBound
		upperBound[metric.Name] = metric.UpperBound
	}

	a.checker = c.New(lowerBound, upperBound)
	if err = a.checker.Validate(); err != nil {
		return
	}

	return
}

func (a *Application) Run() {
	a.loader.Run()
	a.checkerc.Run()
	a.notifierc.Run()
	select {}
}
