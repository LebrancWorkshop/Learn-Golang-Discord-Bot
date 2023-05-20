package botinfoUsecases

import "fmt"

type IBotinfoUsecase interface {
	Feature(msg string) string
}

type botinfoUsecase struct {

}

func (u *botinfoUsecase) Feature(msg string) string {
	return fmt.Sprintf("```Aniki: %v```", msg)
}

func NewBotinfoUsecase() IBotinfoUsecase {
	return &botinfoUsecase{}
}
