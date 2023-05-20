package server

import (
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/modules/botinfo/botinfoHandlers"
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/modules/botinfo/botinfoUsecases"
	"github.com/bwmarrin/discordgo"
)

type IBotinfoModule interface {
	Init()
	Handler() botinfoHandlers.IBotinfoHandler
	Usecase() botinfoUsecases.IBotinfoUsecase
}

type botinfoModule struct {
	*module
	usecase botinfoUsecases.IBotinfoUsecase
	handler botinfoHandlers.IBotinfoHandler
}

func (b *botinfoModule) Init() {
	b.module.commands = append(b.module.commands, &discordgo.ApplicationCommand{
		Name: "help",
		Description: "Just a help menu",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "feature",
				Description: "Feature to get help on",
				Type: discordgo.ApplicationCommandOptionString,
			},
		},
	})

	b.module.commandHandlers["help"] = b.handler.Help
}

func (b *botinfoModule) Handler() botinfoHandlers.IBotinfoHandler {
	return b.handler
}

func (b *botinfoModule) Usecase() botinfoUsecases.IBotinfoUsecase {
	return b.usecase
}

func (m *module) BotinfoModule() IBotinfoModule {
	botinfoUsecase := botinfoUsecases.NewBotinfoUsecase()
	botinfoHandler := botinfoHandlers.NewBotinfoHandler(botinfoUsecase)

	return &botinfoModule{
		module: m,
		handler: botinfoHandler,
		usecase: botinfoUsecase,
	}
}


