package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	// загоняем в json параметры для кнопок
	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup( // добавляем кнопку
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonURL("yandex", "ya.ru"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Next Page", "ya.ru"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("Next Page", "Button info"),
			tgbotapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
		),
	)
	//c.bot.Send(msg) можно просто одной строкой без проверки ошибки
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
func init() {
	registeredCommands["list"] = (*Commander).List
}
