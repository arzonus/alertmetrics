package loader

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/arzonus/alertmetrics/pkg/interfaces/logger"
	"github.com/arzonus/alertmetrics/pkg/interfaces/storage"
	m "github.com/arzonus/alertmetrics/pkg/model"
	"time"
)

type Loader struct {
	db     *sql.DB
	items  chan *m.Items
	period uint

	sinceTime time.Time

	s   storage.Storage
	log logger.ILogger
}

func New(db *sql.DB, items chan *m.Items, period uint, s storage.Storage, log logger.ILogger) *Loader {
	return &Loader{
		db:        db,
		items:     items,
		period:    period,
		log:       log,
		s:         s,
		sinceTime: time.Now(),
	}
}

func (l *Loader) Run() error {
	if l.period < 1 {
		return errors.New(fmt.Sprintf("Loader: period value %d < 1", l.period))
	}
	if l.items == nil {
		return errors.New("Loader: items channel is empty!")
	}

	l.run()

	return nil
}

func (l *Loader) run() {
	go func(l *Loader) {
		for {
			time.Sleep(time.Duration(l.period) * time.Second)
			l.getItems()
		}
	}(l)
}

func (l *Loader) getItems() {
	var err error
	var items = new(m.Items)
	var nowTime = time.Now()

	items, err = l.s.Item.ListSinceTime(l.db, l.sinceTime)
	if err != nil {
		l.log.Error(err)
		return
	}

	if items != nil {
		l.sinceTime = nowTime
		l.items <- items
		return
	}
}
