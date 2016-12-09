package storage

import (
	"github.com/arzonus/alertmetrics/pkg/interfaces/db"
	m "github.com/arzonus/alertmetrics/pkg/model"
	"time"
)

type Item interface {
	ListSinceTime(db.IDatabase, time.Time) (*m.Items, error)
}
