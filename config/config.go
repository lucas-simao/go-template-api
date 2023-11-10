package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/lucas-simao/golog"
)

type Env struct {
	DataSourceUrl string
}

func NewConfig() *Env {
	if err := godotenv.Load(); err != nil {
		log.Critical(err)
	}

	dataSourceUrl, ok := os.LookupEnv("DATASOURCE_URL")
	if !ok {
		log.Critical("failed to load env DATASOURCE_URL")
	}

	return &Env{
		DataSourceUrl: dataSourceUrl,
	}
}
