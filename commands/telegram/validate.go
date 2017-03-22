package telegram

import (
    "github.com/codegangsta/cli"
    "github.com/fatih/color"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/telegram"
)

var CmdTelegramValidate = cli.Command{
    Name: "validate",
    Usage: "Validate the telegram bot token",
    Description: "The tools provide you to validate the telegram bot token is or not valided",
    Action: telegramValidate,
    Flags: []cli.Flag{
    },
}

func telegramValidate(ctx *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Info("Validating ...")

    log.Infof("Bot is enable : %s", configs.ConvertEnableStatus(cronTask.Telegram.Enable))
    log.Infof("Your bot token: %s", cronTask.Telegram.Token)
    log.Infof("Your chat ids : %s", configs.ConvertChatIds(cronTask.Telegram.ChatIds))

    telegramBot, err := telegram.NewTelegram(cronTask.Telegram.Token)

    var validateStatus string
    if telegramBot.ValidateToken() == true {
        validateStatus = color.GreenString("passed")
    }else{
        validateStatus = color.RedString("failed")
    }

    log.Infof("Your token is : %s", validateStatus)

    return nil
}
