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
	period int

	sinceTime time.Time

	s   storage.Storage
	log logger.ILogger
}

func NewLoader(db *sql.DB, items chan *m.Items, period int, s storage.Storage, log logger.ILogger) *Loader {
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

	l.run()

	return nil
}

func (l *Loader) run() {
	go func(l *Loader) {
		for {
			time.Sleep(l.period * time.Second)
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
	}

	if l.items != nil {
		l.sinceTime = nowTime
		l.items <- items
		return
	}

	l.log.Error("Items channel is nil!")
}
