package botinfoHandlers

import (
	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/modules/botinfo/botinfoUsecases"
	"github.com/bwmarrin/discordgo"
)

type IBotinfoHandler interface {
	Help(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type botinfoHandler struct {
	botinfoUsecase botinfoUsecases.IBotinfoUsecase
}

func (h *botinfoHandler) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()
	messageContent := command.Options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: h.botinfoUsecase.Feature(messageContent),
		},
	})
}

func NewBotinfoHandler(botinfoUsecase botinfoUsecases.IBotinfoUsecase) IBotinfoHandler {
	return &botinfoHandler{
		botinfoUsecase: botinfoUsecase,
	}
}
