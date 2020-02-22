package main

import (
    "context"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    commandProto "github.com/lty5240/dcbot/service-command/proto"
    "github.com/micro/go-micro/service"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/service/grpc"
    "github.com/micro/go-plugins/registry/consul/v2"
    "log"
    "os"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")
    registryServer := os.Getenv("REGISTRY_SERVER")

    reg := consul.NewRegistry(func(opt *registry.Options) {
        opt.Addrs = []string{
            registryServer,
        }
    })
    srv := grpc.NewService(
        service.Name("dcbot.client.main"),
        service.Registry(reg),
    )
    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Fatal(err)
    }
    bot.Debug = false

    _, err = bot.RemoveWebhook()
    config := tgbotapi.UpdateConfig{}
    config.Timeout = 60
    updates, err := bot.GetUpdatesChan(config)

    commandClient := commandProto.NewCommandServiceClient("dcbot.service.command", srv.Client())

    for update := range updates {
        if update.Message == nil {
            continue
        }
        //log.Printf("%+v\n", update)
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

        if update.Message.IsCommand() {
            msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Command: "+
                update.Message.Command()+
                "\nText: "+
                update.Message.CommandArguments())
            res, err := commandClient.Find(context.TODO(), &commandProto.CommandRequest{
                Command: update.Message.Command(),
            })
            if err != nil {
                log.Printf(err.Error())
            }
            log.Println(res.Service)
        }
        msg.ReplyToMessageID = update.Message.MessageID
        _, _ = bot.Send(msg)
    }
}
