package telegram

import (
    "fmt"
    "strings"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
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

    log.Infof("Your bot token: %s", cronTask.Telegram.Token)
    log.Infof("Your chat ids : %s", convertChatIds(cronTask.Telegram.ChatIds))

    return nil
}

func convertChatIds(chatIds []int64) string {
    result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(chatIds)), ", "), "[]")

    if len(result) <= 0 {
        result = "n/a"
    }

    return result
}
