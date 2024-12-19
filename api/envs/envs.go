package envs

import (
	"os"
)

type Envs struct {
	API_PORT string

	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

var ServerEnvs Envs

func LoadEnvs() error {

	// if err := godotenv.Load("../config.env"); err != nil {
	// 	return err
	// }

	ServerEnvs.API_PORT = os.Getenv("API_PORT")

	ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ServerEnvs.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	ServerEnvs.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")

	return nil
}
