package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/shamhub/pdfprovider/types"
)

type UniDocConfig struct {
	Key string
}

func NewUniDocCred() (*UniDocConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	key := os.Getenv(types.UNIDOC_LICENSE_API_KEY)
	if len(key) == 0 {
		return nil, errors.New("uni doc key not found")
	}

	return &UniDocConfig{Key: key}, nil
}

func (u *UniDocConfig) Get(key string) string {
	return u.Key
}
