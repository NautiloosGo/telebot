package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NautiloosGo/telebot/internal/app/commands"     // то работает, то удаляется, хз
	"github.com/NautiloosGo/telebot/internal/services/product" // то работает, то удаляется, хз

	//".internal/services/product" //странным образом работает
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() //подгружаю инфу из файла .env
	myToken := os.Getenv("TOKEN")
	if myToken == "" {
		log.Printf("input telegram bot unique token:") //если нет файла .env то спросить ввести токен в панель
		fmt.Fscan(os.Stdin, &myToken)
		//scanner := bufio.NewScanner(os.Stdin)
		//myToken = os.Getenv(scanner.Text())
	} // сохраняем токен на компе. Еще токен можно ввести перед вызовом TOKEN="dfgdfg" go run main.go
	bot, err := tgbotapi.NewBotAPI(myToken) // секретный токен чата
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true //куча всякой инфы по статусам

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//телега хранит 24 часа все что пришло на бота.
	u := tgbotapi.UpdateConfig{ //почему-то он поменял не как в шаблоне а на config
		Timeout: 60,
	}
	updates := bot.GetUpdatesChan(u)

	product.GetJsonCatalog() //подгружаю каталог из файла

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.CallbackQuery == nil { //проверка если не кнопка, то вывести текст
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text) //то что получено показываю в панель
		}
		commander.HandleUpdate(update)
	}
}
