package checkerc

import (
	"fmt"
	c "github.com/arzonus/alertmetrics/pkg/checker"
	m "github.com/arzonus/alertmetrics/pkg/model"
	"github.com/vulteam/api/libs/errors"
)

type CheckerController struct {
	checker *c.Checker

	report chan string
	items  chan *m.Items
}

func NewCheckerController(report chan string, items chan *m.Items) {
	return &CheckerController{
		report: report,
		items:  items,
	}
}

func (c *CheckerController) Run() error {
	if c.report == nil {
		return errors.New(fmt.Sprint("CheckerController: report chan is null"))
	}
	if c.items == nil {
		return errors.New(fmt.Sprint("CheckerController: items chan is null"))
	}

	c.run()
	return nil
}

func (c *CheckerController) run() {
	go func(c *CheckerController) {
		for {
			var items = new(m.Items)
			items <- c.items
			c.report <- c.checker.CheckItems(&items).Report()
		}
	}(c)
}
