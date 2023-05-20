package main

import (
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/config"
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/modules/server"
)

func main() {
	cfg := config.NewConfig("./.env")

	server.NewDiscordServer(cfg).Start()
}
