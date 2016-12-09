package notifierc

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(msg string)
}

type NotifierController struct {
	report    chan string
	notifiers []Notifier
}

func New(report chan string, notifies []Notifier) *NotifierController {
	return &NotifierController{
		report:    report,
		notifiers: notifies,
	}
}

func (n *NotifierController) Run() error {
	if n.report == nil {
		return errors.New(fmt.Sprint("Notifier: items chan is nill"))
	}
	if len(n.notifiers) == 0 {
		return errors.New(fmt.Sprint("Notifier: doesn't have notifiers"))
	}

	n.run()
	return nil
}

func (n *NotifierController) run() {

	go func(n *NotifierController) {
		for {
			report := <-n.report

			for _, notifier := range n.notifiers {
				go func(notifier Notifier, report string) {
					notifier.Send(report)
				}(notifier, report)
			}
		}
	}(n)
}
