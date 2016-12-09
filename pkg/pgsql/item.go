package pgsql

import (
	"github.com/arzonus/alertmetrics/pkg/interfaces/db"
	m "github.com/arzonus/alertmetrics/pkg/model"
	"time"
)

type Item struct{}
type ItemModel struct {
	ID nullString

	Metrics map[string]nullInt64

	Created nullTime
}

func (im *ItemModel) convert() *m.Item {
	return &m.Item{
		ID:      im.ID.String,
		Metrics: im.Metrics,
		Created: im.Created.Time,
	}
}

func (Item) ListSinceTime(db db.IDatabase, sinceTime time.Time) (*m.Items, error) {
	var err error
	var items = new(m.Items)

	strsql := `
		SELECT *
		FROM item_metrics
		WHERE created_time > $1`

	rows, err := db.Query(strsql, newNullTime(sinceTime))
	if err != nil {
		return nil, err
	}

	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// support multiple metrcis count
	// metric name in sql always equal metric name in config
	for rows.Next() {
		var im = new(ItemModel)

		columns := make([]interface{}, len(columnNames))
		columnPointers := make([]interface{}, len(columnNames))
		for i := 0; i < len(columnNames); i++ {
			columnPointers[i] = &columns[i]
		}

		if err = rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// IN table ID field must be first
		// IN table Created field must be last
		im.ID = columns[0]
		im.Created = columns[len(columns)-1]

		for i := 1; i < len(columnNames)-1; i++ {
			im.Metrics[columnNames[i]] = columns[i]
		}

		*items = append(*items, *(im.convert()))

	}

	return items, nil
}
