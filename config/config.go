package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	Host                  string
	Port                  int
	PostgresHost          string
	PostgresPort          int
	PostgresUser          string
	PostgresPass          string
	PostgresDB            string
	PostgresMigrationPath string
}

func Load() *Config {
	return &Config{
		Host:                  cast.ToString(getOrReturnDefault("HOST", "localhost")),
		Port:                  cast.ToInt(getOrReturnDefault("PORT", "8080")),
		PostgresHost:          cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:          cast.ToInt(getOrReturnDefault("POSTGRES_PORT", "5432")),
		PostgresUser:          cast.ToString(getOrReturnDefault("POSTGRES_USER", "akbarshoh")),
		PostgresPass:          cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1")),
		PostgresDB:            cast.ToString(getOrReturnDefault("POSTGRES_DB", "exam")),
		PostgresMigrationPath: cast.ToString(getOrReturnDefault("POSTGRES_MIGRATION_PATH", "migrations")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if _, ok := os.LookupEnv(key); ok {
		return os.Getenv(key)
	}

	return defaultValue
}
