package mail

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
)

var CmdMailSend = cli.Command{
    Name: "send",
    Usage: "Send the dummy notification mail",
    Description: "The tools provide you to send dummy mail",
    Action: mailSend,
    Flags: []cli.Flag{
    },
}

func mailSend(cli *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Info("Sending .....\n")

    // TODO: send dummy mail
    log.Info(cronTask)

    return nil
}
