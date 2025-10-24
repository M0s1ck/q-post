package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "../../..")

	setEnvErr := os.Setenv("IN_DOCKER", "0")
	if setEnvErr != nil {
		panic(setEnvErr)
	}

	err := godotenv.Load(filepath.Join(dir, ".env"))

	if errors.Is(err, os.ErrNotExist) {
		log.Println(".env file not found. That means we're in docker")

		setEnvErr = os.Setenv("IN_DOCKER", "1")

		if setEnvErr != nil {
			panic(setEnvErr)
		}
	} else {
		setEnvErr = os.Setenv("POSTGRES_HOST", "localhost")
		setEnvErr = os.Setenv("USER_SERVICE_HOST", "localhost")

		if setEnvErr != nil {
			panic(setEnvErr)
		}
	}

	// Check
	_, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		log.Fatalln("Port variable not found in env. Check if env is loaded")
	}
}
