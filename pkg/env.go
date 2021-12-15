package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

func GetEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	log.Println(1)

	var buf bytes.Buffer
	f, err := os.Open(fmt.Sprintf("/usr/local/app/%s", key))
	if err != nil {
		log.Println(3)
		return "", err
	}
	log.Println(2)

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	n, err := buf.ReadFrom(f)
	if err != nil {
		return "", err
	}
	if n == 0 {
		return "", errors.New("file has been empty")
	}

	log.Println("N, buf", n, buf.String())

	return buf.String(), nil
}
