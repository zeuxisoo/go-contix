package main

import (
    "os"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands"
)

const (
    APP_VERSION = "0.3.0"
)

func main() {
    app := cli.NewApp()
    app.Name = "Contix"
    app.Usage = "A tools for checking the target performance is or not available for sales"
    app.Version = APP_VERSION
    app.Commands = []cli.Command{
        commands.CmdProxy,
        commands.CmdCron,
        commands.CmdMail,
        commands.CmdEditor,
        commands.CmdTelegram,
    }
    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
