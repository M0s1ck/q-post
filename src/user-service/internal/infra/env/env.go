package env

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	PsgConf   *PostgresConfig
	JWTSecret string
	ApiSecret string
	AppPort   string
}

type PostgresConfig struct {
	User     string
	Password string
	DB       string
	Host     string
	Port     string
	Scheme   string
}

func BuildEnvConfig() *Config {
	loadEnv()

	psgConf := &PostgresConfig{}
	psgConf.User = os.Getenv("POSTGRES_USER")
	psgConf.Password = os.Getenv("POSTGRES_PASSWORD")
	psgConf.DB = os.Getenv("POSTGRES_DB")
	psgConf.Host = os.Getenv("POSTGRES_HOST")
	psgConf.Port = os.Getenv("POSTGRES_PORT")
	psgConf.Scheme = os.Getenv("POSTGRES_COMMUNITY_SCHEME")

	if os.Getenv("IN_DOCKER") == "0" {
		psgConf.Host = "localhost"
	}

	envConf := &Config{}
	envConf.JWTSecret = os.Getenv("JWT_SECRET_KEY")
	envConf.ApiSecret = os.Getenv("API_SECRET_KEY")
	envConf.AppPort = os.Getenv("USER_SERVICE_PORT")
	envConf.PsgConf = psgConf

	return envConf
}

func loadEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "../../../..")

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
	}

	// Check
	_, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		log.Fatalln("Port variable not found in env. Check if env is loaded")
	}
}
