package commands

import (
	"github.com/NautiloosGo/telebot/internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) { //переключатель
	if update.Message == nil { // If we got a message не обрабатывает другие обновления non-message
		return
	}

	command, ok := registeredCommands[update.Message.Command()] // если команда, то вызвать функ из мап по названию
	if ok {
		command(c, update.Message) // в мапе лежат процедуры с аргументами с (коммандер) и мессага
	} else {
		c.Default(update.Message)
	}

}
