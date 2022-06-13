package commands

import (
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

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx= %v", idx, err)
	}

	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument : %v", arg))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
func init() {
	registeredCommands["get"] = (*Commander).Get
}
