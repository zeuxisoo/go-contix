package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/mail"
)

var CmdMail = cli.Command{
    Name: "mail",
    Usage: "A tools for you test the mail action",
    Description: "The tools provide you can to send or render the dummy mail",
    Subcommands: []cli.Command{
        mail.CmdMailSend,
        mail.CmdMailRender,
    },
}
