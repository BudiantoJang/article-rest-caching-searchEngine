package env

import (
	"jang-article/internal/model"
	"jang-article/internal/port"

	"github.com/Netflix/go-env"
)

type envBasedReader struct{}

func NewEnvBasedReader() port.ConfigReader {
	return envBasedReader{}
}

func (r envBasedReader) Read() (cfg model.Config, err error) {
	var eCfg model.Config
	_, err = env.UnmarshalFromEnviron(&eCfg)
	if err != nil {
		return
	}
	cfg.RedisConfig.Host = eCfg.RedisConfig.Host
	cfg.RedisConfig.Password = eCfg.RedisConfig.Password
	cfg.RedisConfig.DB = eCfg.RedisConfig.DB

	cfg.RedisearchConfig.Host = eCfg.RedisearchConfig.Host
	cfg.RedisearchConfig.Index = eCfg.RedisearchConfig.Index

	cfg.PostgresConfig.Host = eCfg.PostgresConfig.Host
	cfg.PostgresConfig.User = eCfg.PostgresConfig.User
	cfg.PostgresConfig.Password = eCfg.PostgresConfig.Password
	cfg.PostgresConfig.DBName = eCfg.PostgresConfig.DBName
	cfg.PostgresConfig.Port = eCfg.PostgresConfig.Port

	return
}
