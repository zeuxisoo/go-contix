package mail

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/mail"
)

var CmdMailSend = cli.Command{
    Name: "send",
    Usage: "Send the dummy notification mail",
    Description: "The tools provide you to send dummy mail",
    Action: mailSend,
    Flags: []cli.Flag{
    },
}

func mailSend(ctx *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Info("Sending .....")

    response, id, err := mail.NewMailgun(cronTask.Mail.Mailgun.Domain, cronTask.Mail.Mailgun.ApiKey, "",).
        SetSender(cronTask.Mail.Sender).
        SetRecipient(cronTask.Mail.Recipient).
        SetSubject(cronTask.Mail.Subject).
        SetContent("This is a test mail sent from contix").
        Send()
    if err != nil {
        return err
    }

    log.Infof("ID: %s, Response: %s", id, response)

    return nil
}
