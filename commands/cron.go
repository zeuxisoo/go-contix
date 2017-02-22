package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/cron"
)

var CmdCron = cli.Command{
    Name: "cron",
    Usage: "A tools for run the scheduled task",
    Description: "The tools provide you can to loop the default task",
    Subcommands: []cli.Command{
        cron.CmdCronRun,
        cron.CmdCronList,
    },
}
