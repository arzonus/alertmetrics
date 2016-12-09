package notifiers

import "github.com/arzonus/alertmetrics/pkg/interfaces/logger"

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
	n.log.Debug(msg)
}
