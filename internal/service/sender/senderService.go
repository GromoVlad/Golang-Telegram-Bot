package sender

import (
	"fmt"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang_telegram_bot/internal/repository/userRepository"
	"log"
	"os"
)

func New(apiKey string) *Client {
	bot, err := telegramBotAPI.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	return &Client{bot: bot}
}

func Dispatch(template string) {
	client := New(os.Getenv("TELEGRAM_TOKEN"))
	contacts := userRepository.FindSubscribers()

	for _, contact := range contacts {
		client.sendMessage(template, int64(contact.TelegramId))
	}
}

func (c *Client) sendMessage(text string, chatId int64) {
	message := telegramBotAPI.NewMessage(chatId, text)
	message.ParseMode = "Markdown"
	_, err := c.bot.Send(message)
	if err != nil {
		fmt.Println(err)
	}
}

type Client struct {
	bot *telegramBotAPI.BotAPI
}
