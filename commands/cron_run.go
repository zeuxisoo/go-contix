package commands

import (
    "fmt"

    "github.com/codegangsta/cli"
)

var CmdCronRun = cli.Command{
    Name: "run",
    Usage: "Run the schedule task",
    Description: "The tools provide you to run scheduled task",
    Action: cronRun,
    Flags: []cli.Flag{
    },
}

func cronRun(cli *cli.Context) error {
    fmt.Println("This is a cron run command")

    return nil
}
