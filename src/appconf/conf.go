package appconf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type redisConf struct {
	Host string
	Port string
	Pass string
	Db   int
}

type config struct {
	Port  string
	Redis redisConf
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

	db, err := strconv.Atoi(os.Getenv("APP_REDIS_DB"))
	if err != nil {
		log.Println("Failed to load redis db, falling back to zero")
		db = 0
	}
	Config = config{
		Port: os.Getenv("APP_API_PORT"),
		Redis: redisConf{
			Host: os.Getenv("APP_REDIS_HOST"),
			Port: os.Getenv("APP_REDIS_PORT"),
			Pass: os.Getenv("APP_REDIS_PASS"),
			Db:   db,
		},
	}
}
