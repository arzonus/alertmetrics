package context

import (
	"github.com/arzonus/alertmetrics/pkg/config"
	"github.com/arzonus/alertmetrics/pkg/interfaces/logger"
	"github.com/arzonus/alertmetrics/pkg/interfaces/storage"

	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/arzonus/alertmetrics/pkg/pgsql"
	_ "github.com/lib/pq"
)

var context Context

func Get() *Context {
	return &context
}

type Context struct {
	Database *sql.DB
	Storage  storage.Storage
	Logger   logger.ILogger
}

func Init() (*Context, error) {
	cfg := config.Get()

	context.setLogger()
	context.setStorage()

	if err := context.setDatabase(cfg); err != nil {
		return nil, err
	}

	return &context, nil
}

func (ctx *Context) setLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	ctx.Logger = logrus.StandardLogger()
}

func (ctx *Context) setDatabase(cfg *config.Config) (err error) {
	ctx.Database, err = sql.Open("postgres", cfg.Database.Connection)
	return
}

func (ctx *Context) setStorage() {
	ctx.Storage.Item = pgsql.Item{}
}
