package botinfoHandlers

import "github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/modules/botinfo/botinfoUsecases"

type IBotinfoHandler interface {

}

type botinfoHandler struct {
	botinfoUsecase botinfoUsecases.IBotinfoUsecase
}

func NewBotinfoHandler(botinfoUsecase botinfoUsecases.IBotinfoUsecase) IBotinfoHandler {
	return &botinfoHandler{
		
	}
}
