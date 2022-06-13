package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	mess := "all commads available: \n\n"
	for names, _ := range registeredCommands {
		mess = mess + "\n/" + names
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)

	//"/help - описание\n/nope - еще\n/list - все команды\n")
	c.bot.Send(msg)
}
func init() {
	registeredCommands["help"] = (*Commander).Help
}
