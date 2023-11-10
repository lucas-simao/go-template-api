package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/lucas-simao/golog"
)

func NewDB(dataSource string) *sqlx.DB {
	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Critical(err)
	}

	return conn
}
