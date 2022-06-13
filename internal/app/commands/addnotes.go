package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Addnotes(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "what ever..")
	c.bot.Send(msg)
}
func init() {
	registeredCommands["addnote"] = (*Commander).Addnotes
}

// func NewCommand(n *tgbotapi.Message) {
// 	// создать файл
// 	notesFile, err := os.OpenFile("notes.txt")
// 	if err != nil {
// 		return
// 	}
// 	defer notesFile.Close()

// 	notesFile.WriteString(n)
// 	// загрузить название в мапу

// }
