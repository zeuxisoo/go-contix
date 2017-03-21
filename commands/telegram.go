package commands

import (
    "github.com/codegangsta/cli"
)

var CmdTelegram = cli.Command{
    Name: "telegram",
    Usage: "A tools for you test the telegram action",
    Description: "The tools provide you can to send or get info base on telegram",
    Subcommands: []cli.Command{
    },
}
