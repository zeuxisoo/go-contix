package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/mail"
)

var CmdMail = cli.Command{
    Name: "mail",
    Usage: "A tools for run the mail test",
    Description: "The tools provide you can to send the dummy mail",
    Subcommands: []cli.Command{
        mail.CmdMailSend,
        mail.CmdMailRender,
    },
}
