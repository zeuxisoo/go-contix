package main

import (
    "os"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands"
)

const (
    APP_VERSION = "0.1.0"
)

func main() {
    app := cli.NewApp()
    app.Name = "Contix"
    app.Usage = "A application console for contix"
    app.Version = APP_VERSION
    app.Commands = []cli.Command{
        commands.CmdProxy,
        commands.CmdCron,
        commands.CmdMail,
    }
    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
