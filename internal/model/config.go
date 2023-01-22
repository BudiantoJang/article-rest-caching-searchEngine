package model

type (
	Config struct {
		App              App              `json:"app"`
		RedisConfig      RedisConfig      `json:"redisConfig"`
		RedisearchConfig RedisearchConfig `json:"redisearchConfig"`
		PostgresConfig   PostgresConfig   `json:"postgresConfig"`
	}

	App struct {
		Port int `json:"port"`
	}

	RedisConfig struct {
		Host     string `env:"REDIS_HOST"`
		Password string `env:"REDIS_PASSWORD"`
		DB       string `env:"REDIS_DB"`
	}

	RedisearchConfig struct {
		Host  string `env:"REDISEARCH_HOST"`
		Index string `env:"REDISEARCH_INDEX"`
	}

	PostgresConfig struct {
		Host     string `env:"POSTGRES_HOST"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		DBName   string `env:"POSTGRES_DB_NAME"`
		Port     string `env:"POSTGRES_PORT"`
	}
)
