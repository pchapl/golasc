package appconf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	Port string
}

var Config config

func load(file string) {
	err := godotenv.Load()
	if err == nil {
		log.Println(file, "loaded")
	} else {
		log.Println(file, "not loaded:", err)
	}
}

func init() {
	load(".env.local")
	load(".env")

	Config = config{
		Port: os.Getenv("APP_API_PORT"),
	}
}
