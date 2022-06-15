package commands

import (
	"encoding/json"
	"log"

	"github.com/NautiloosGo/telebot/internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Task     string
	Parametr int
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

// основная рукоятка переключения команд
func (c *Commander) HandleUpdate(update tgbotapi.Update) { //переключатель
	defer func() { //обработка паники перед закрытием
		if panicValue := recover(); panicValue != nil { //! можно создать переменную прямо в ифе так.
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	// обработка кнопок c Data типа "get_1" где get можно интерпретировать как команду
	if update.CallbackQuery != nil {
		//альтернативно через текст args := strings.Split(update.CallbackQuery.Data, "_") //парсим текст в Data кнопки
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		switch parsedData.Task {
		case "Pagenum":
			c.Pagenum(update.CallbackQuery.Message, parsedData.Parametr)
			log.Printf("Pagenum")
		default:
			log.Printf("not Pagenum")
		}
		//вывожу ответом что спарсено
		// msg := tgbotapi.NewMessage(
		// 	update.CallbackQuery.Message.Chat.ID,
		// 	fmt.Sprintf("Parsed: %+v\n", parsedData),
		// )

		// а вот вызвать функцию из списка не получилось, т.к. аргумент функций не string, а *tgbotapi.Message
		// можно другую мапу с аргументами типа стринг сделать. И другого свитчера чисто под кнопки.
		// if _, err := c.bot.Send(msg); err != nil {
		// 	log.Panic(err)
		// }
		return
	}

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
