package app

import (
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
	if err != nil {
		log.Println(".env file not found. That means we're in docker")

		setEnvErr = os.Setenv("IN_DOCKER", "1")
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
