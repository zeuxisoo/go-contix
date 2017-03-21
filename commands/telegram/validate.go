package telegram

import (
    "github.com/codegangsta/cli"

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
    log.Info("Validating ...")

    return nil
}
