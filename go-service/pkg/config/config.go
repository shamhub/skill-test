package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/matryer/resync"
)

type EnvConfig struct {
}

var Once resync.Once

func NewEnvConfig() (envConfig *EnvConfig, err error) {

	Once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			return
		}
		envConfig = &EnvConfig{}
	})
	return
}

func (u *EnvConfig) Get(key string) string {
	return os.Getenv(key)
}
