package main

import (
	"context"
	"fmt"
	"jang-article/internal/adapter/env"
	pg "jang-article/internal/adapter/postgres"
	rds "jang-article/internal/adapter/redis"
	"jang-article/internal/adapter/usecase"
	"jang-article/internal/adapter/validation"

	handler "jang-article/internal/adapter/http"

	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	envReader := env.NewEnvBasedReader()
	config, err := envReader.Read()
	if err != nil {
		log.Panic(err)
	}

	redisDBNum, err := strconv.Atoi(config.RedisConfig.DB)
	if err != nil {
		log.Panic(err)
	}

	vald := validation.New()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host,
		Password: config.RedisConfig.Password,
		DB:       redisDBNum,
	})

	rds := rds.New(redisClient)

	_, err = rds.CheckHealth(ctx)
	if err != nil {
		log.Panic(err)
	}

	pgDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.PostgresConfig.Host,
		config.PostgresConfig.User,
		config.PostgresConfig.Password,
		config.PostgresConfig.DBName,
		config.PostgresConfig.Port)

	postgres, err := gorm.Open(postgres.Open(pgDsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	pgClient := pg.New(postgres)

	uc := usecase.NewUsecases(vald, pgClient, rds)

	config.App.Port = 8080

	app := handler.NewRouter(config, rds, pgClient, uc)

	app.Start(ctx)
}
