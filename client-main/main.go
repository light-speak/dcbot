package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
    "os"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Fatal(err)
    }
    bot.Debug = true

    _, err = bot.RemoveWebhook()
    config := tgbotapi.UpdateConfig{}
    config.Timeout = 60
    updates, err := bot.GetUpdatesChan(config)

    for update := range updates {
        if update.Message == nil {
            continue
        }
        log.Printf("%+v\n", update)
    }
}
