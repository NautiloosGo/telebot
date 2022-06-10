package commands

import (
	"github.com/NautiloosGo/telebot/internal/services/product" // то работает, то удаляется, хз

	//".internal/services/product" //странным образом работает
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
