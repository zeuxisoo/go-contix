package telegram

import (
    "fmt"
    "strings"

    "github.com/codegangsta/cli"
    "github.com/fatih/color"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/telegram"
)

var CmdTelegramRun = cli.Command{
    Name: "run",
    Usage: "Run the telegram bot to get the chat id",
    Description: "The tools provide you to run the telegram bot ask the chat id",
    Action: telegramRun,
    Flags: []cli.Flag{
    },
}

func telegramRun(ctx *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Info("Starting ...")

    telegramBot, err := telegram.NewTelegram(cronTask.Telegram.Token)
    if err != nil {
        return err
    }

    log.Infof("Your bot token : %s", cronTask.Telegram.Token)
    log.Infof("Your bot name  : %s", telegramBot.GetBotUsername())

    updates, err := telegramBot.GetUpdatesChannel()
    if err != nil {
        return err
    }

    log.Infof("--------------------")
    log.Infof("Bot is stared, Please update the telegram.chat_ids in cron_task.yaml file")
    log.Infof("--------------------")
    log.Infof("1. Open this link http://t.me/contixbot")
    log.Infof("2. Chat with @contixbot bot")
    log.Infof("3. Send message: %s to get your chat id", color.GreenString("/chatid"))
    log.Infof("--------------------")

    for update := range updates {
        if update.Message == nil {
            continue
        }

        username := update.Message.From.UserName
        text     := update.Message.Text
        chatId   := update.Message.Chat.ID

        log.Infof("[%s] >> %s", username, text)

        if strings.ToLower(text) == "/chatid" {
            reply   := fmt.Sprintf("@%s, Your chat id is %d", username, chatId)
            message := telegramBot.CreateMessage(chatId, reply)

            telegramBot.SendMessage(message)

            log.Infof("[[Bot]] >> %s", reply)
        }
    }

    return nil
}
