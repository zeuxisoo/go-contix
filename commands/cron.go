package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/cron"
)

var CmdCron = cli.Command{
    Name: "cron",
    Usage: "A tools for you control the task list",
    Description: "The tools provide you can to lookup / start the task list",
    Subcommands: []cli.Command{
        cron.CmdCronRun,
        cron.CmdCronList,
    },
}
