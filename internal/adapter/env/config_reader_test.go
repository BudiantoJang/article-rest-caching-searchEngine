package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetConfigFromEnvg(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	os.Setenv("REDIS_HOST", "localhost:redis")
	os.Setenv("REDIS_PASSWORD", "admin")
	os.Setenv("REDIS_DB", "3")

	os.Setenv("REDISEARCH_HOST", "localhost:redisearch")
	os.Setenv("REDISEARCH_INDEX", "index")

	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "jang")
	os.Setenv("POSTGRES_PASSWORD", "superuser")
	os.Setenv("POSTGRES_DB_NAME", "article")
	os.Setenv("POSTGRES_PORT", "5432")

	cr := NewEnvBasedReader()
	cfg, err := cr.Read()
	assert.NoError(t, err)

	assert.Equal(t, "localhost:redis", cfg.RedisConfig.Host)
	assert.Equal(t, "admin", cfg.RedisConfig.Password)
	assert.Equal(t, "3", cfg.RedisConfig.DB)

	assert.Equal(t, "localhost:redisearch", cfg.RedisearchConfig.Host)
	assert.Equal(t, "index", cfg.RedisearchConfig.Index)

	assert.Equal(t, "localhost", cfg.PostgresConfig.Host)
	assert.Equal(t, "jang", cfg.PostgresConfig.User)
	assert.Equal(t, "superuser", cfg.PostgresConfig.Password)
	assert.Equal(t, "article", cfg.PostgresConfig.DBName)
	assert.Equal(t, "5432", cfg.PostgresConfig.Port)

	cfg, err = cr.Read()
	assert.NoError(t, err)

}
