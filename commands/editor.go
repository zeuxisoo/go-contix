package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/editor"
)

var CmdEditor = cli.Command{
    Name: "editor",
    Usage: "A tools for you manage the cron task configure file",
    Description: "The tools provide a user interface to manage the cron task configure file",
    Subcommands: []cli.Command{
        editor.CmdEditorRun,
    },
}
