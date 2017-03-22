package telegram

import (
    "github.com/codegangsta/cli"
    "github.com/fatih/color"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/mail"
    "github.com/zeuxisoo/go-contix/utils/telegram"
)

var CmdTelegramSend = cli.Command{
    Name: "send",
    Usage: "Send the dummy data to telegram",
    Description: "The tools provide you to send the dummy data to telegram",
    Action: telegramSend,
    Flags: []cli.Flag{
    },
}

func telegramSend(ctx *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Infof("Your bot token: %s", cronTask.Telegram.Token)
    log.Infof("Your chat ids : %s", configs.ConvertChatIds(cronTask.Telegram.ChatIds))

    if len(cronTask.Telegram.ChatIds) <= 0 {
        log.Fatalf("The telegram ids is empty, Please ran `%s` first", color.GreenString("telegram run"))
    }

    log.Info("Reading notification mail content ...")

    dummy := models.MailNotificationData{
        Name: "Title",
        Performances: []models.MailNotificationDataPerformance{
            models.MailNotificationDataPerformance{ Name: "A Performance", State: "YES" },
            models.MailNotificationDataPerformance{ Name: "B Performance", State: "YES" },
        },
    }

    content, err := mail.RenderMailNotification(dummy)
    if err != nil {
        return err
    }

    log.Info("Sending to telegram .....")

    telegramBot, err := telegram.NewTelegram(cronTask.Telegram.Token)
    if err != nil {
        return err
    }

    for index := range cronTask.Telegram.ChatIds {
        chatId  := cronTask.Telegram.ChatIds[index]
        message := telegramBot.CreateMessage(chatId, content)
        telegramBot.SendMessage(message)

        log.Infof("chat id: %d sent", chatId)
    }

    return nil
}
