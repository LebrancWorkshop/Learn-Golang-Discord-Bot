package server

import (
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/config"
	"github.com/bwmarrin/discordgo"
)

type IDiscordServer interface {

}

type discordServer struct {
	cfg 				config.IConfig
	dg 					*discordgo.Session
	commands 		[]*discordgo.ApplicationCommand 
}


