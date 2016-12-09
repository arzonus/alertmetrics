package notifier

import (
	"github.com/arzonus/alertmetrics/pkg/interfaces/logger"
	"strings"
)

type LogNotifier struct {
	Notifier
	log logger.ILogger
}

func NewLogNotifier(log logger.ILogger) *LogNotifier {
	return &LogNotifier{
		log: log,
	}
}

func (n *LogNotifier) Send(msg string) {
	n.log.Debug("LogNotifier send notify:")
	reports := strings.Split(msg, "\n")
	for _, report := range reports {
		n.log.Debug(report)
	}
}
