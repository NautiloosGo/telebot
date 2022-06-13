package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	mess := "all commads available: \n\n"
	for names, _ := range registeredCommands {
		mess = mess + "\n/" + names
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)

	//"/help - описание\n/nope - еще\n/list - все команды\n")
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
func init() {
	registeredCommands["help"] = (*Commander).Help
}
