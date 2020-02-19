package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
    "net/http"
    "os"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")
    botHost := os.Getenv("BOT_HOST")
    botPort := os.Getenv("BOT_PORT")

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Fatal(err)
    }
    bot.Debug = true
    _, err = bot.RemoveWebhook()
    history, err := bot.GetUpdates(tgbotapi.UpdateConfig{})
    for h := range history {
        log.Printf("%+v\n", h)
    }
    webHook := botHost + ":" + botPort + "/" + bot.Token
    log.Printf("WebHook为： %s\n", webHook)
    _, err = bot.SetWebhook(tgbotapi.NewWebhook(webHook))
    if err != nil {
        log.Fatal(err)
    }
    _, err = bot.GetWebhookInfo()
    if err != nil {
        log.Fatal(err)
    }
    //if info.LastErrorDate != 0 {
    //    log.Printf("Telegram回调错误: %s", info.LastErrorMessage)
    //}
    updates := bot.ListenForWebhook("/" + bot.Token)
    go http.ListenAndServe("0.0.0.0:"+botPort, nil)

    for update := range updates {
        log.Printf("%+v\n", update)
    }
}
