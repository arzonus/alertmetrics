package notifier

import (
	"fmt"
	"github.com/vulteam/api/libs/errors"
)

type Notifier interface {
	Send(msg string)
}

type NotifierController struct {
	report    chan string
	notifiers []Notifier
}

func NewNotifierController(report chan string, notifies Notifier) *NotifierController {
	return &NotifierController{
		report:    report,
		notifiers: notifies,
	}
}

func (n *NotifierController) Run() error {
	if n.report == nil {
		return errors.New(fmt.Sprint("Notifier: items chan is nill"))
	}
	if !(len(n.notifiers)) {
		return errors.New(fmt.Sprint("Notifier: doesn't have notifiers"))
	}

	n.run()
	return nil
}

func (n *NotifierController) run() {

	go func(n *NotifierController) {
		for {
			var report string
			report <- n.report

			for _, notifier := range n.notifiers {
				go func(notifier Notifier, report string) {
					notifier.Send(report)
				}(notifier, report)
			}
		}
	}(n)
}
