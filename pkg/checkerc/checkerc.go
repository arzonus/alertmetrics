package checkerc

import (
	"errors"
	"fmt"
	c "github.com/arzonus/alertmetrics/pkg/checker"
	m "github.com/arzonus/alertmetrics/pkg/model"
)

type CheckerController struct {
	checker *c.Checker

	report chan string
	items  chan *m.Items
}

func New(report chan string, items chan *m.Items, checker *c.Checker) *CheckerController {
	return &CheckerController{
		report:  report,
		items:   items,
		checker: checker,
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
			items := <-c.items
			checkedItems := c.checker.CheckItems(*items)
			if checkedItems != nil {
				c.report <- checkedItems.Report()
			}
		}
	}(c)
}
