package pkg

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	value, err := os.ReadFile(fmt.Sprintf("./%s", key))
	if err != nil {
		return "", err
	}

	return string(value), nil
}
