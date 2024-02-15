package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type SystemConfig struct {
	Port              string `envconfig:"SERVER_PORT" required:"true"`
	ServiceName       string `envconfig:"SERVICE_NAME" required:"true"`
	OtelTraceEndpoint string `envconfig:"OTEL_TRACE_ENDPOINT" required:"true"`
}

type MongoConfig struct {
	URI    string `envconfig:"MONGO_URI" required:"true"`
	DBName string `envconfig:"MONGO_DB_NAME" required:"true"`
	DBTZ   string `envconfig:"MONGO_DB_TZ" required:"true"`
}

type AppConfig struct {
	System SystemConfig
	Mongo  MongoConfig
}

func (cfg *AppConfig) Init() {
	envconfig.MustProcess("", &cfg.System)
	envconfig.MustProcess("", &cfg.Mongo)
}

func LoadAppConfig() (appCfg *AppConfig) {
	env, ok := os.LookupEnv("ENV")
	if ok && env == "local" {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/../../.env.%v", basePath, env))
		if err != nil {
			err = godotenv.Load()
			if err != nil {
				panic(err)
			}
		}
	}
	appCfg = &AppConfig{}
	appCfg.Init()
	return
}
