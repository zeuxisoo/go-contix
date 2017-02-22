package commands

import (
    "github.com/codegangsta/cli"
)

var CmdCron = cli.Command{
    Name: "cron",
    Usage: "A tools for run the scheduled task",
    Description: "The tools provide you can to loop the default task",
    Subcommands: []cli.Command{
        CmdCronRun,
        CmdCronList,
    },
}
