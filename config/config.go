package config

import (
	"log"

	"github.com/joho/godotenv"
)

type config struct {
	app *app
}

type IConfig interface {
	App() IAppConfig
}

type app struct {
	token string
}

type IAppConfig interface {
	GetToken() string
}

func NewConfig() IConfig {
	envMap, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &config{
		app: &app{
			token: envMap["DISCORD_BOT_TOKEN"],
		},
	}
}

func (c *config) App() IAppConfig {
	return c.app
}


func (a *app) GetToken() string {
	return a.token
}
