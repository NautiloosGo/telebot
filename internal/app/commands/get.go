package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments() // вернет текст то что после команды в сообщении. добавили для получения аргумента из сообщения. Вообще можно распарсить значение перем. Text
	idx, err := strconv.Atoi(args)          // конвертируем в инт (буквы сконвертит в 0)
	if err != nil {
		log.Println("wrong argument number: ", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument number.\nFor example: \n/get 2"))
		return
	}

	product, ok := c.productService.Get(idx)
	if ok == false {
		log.Printf("fail to get product with idx = %v", idx, err)
	} else {

		//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument : %v", arg))
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("id: %v\nTitle: %s\nDescription: %s\nPrice: %f", idx, product.Title, product.Description, product.Price))
		if _, err := c.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
func init() {
	registeredCommands["get"] = (*Commander).Get
}
