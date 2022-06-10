package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text+"\n/Help for more info") // куда отправлять и что отправлять
	msg.ReplyToMessageID = inputMessage.MessageID                                                            //отправляет сообщения как реплай (можно убрать)
	c.bot.Send(msg)
}
