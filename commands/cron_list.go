package commands

import (
    "fmt"

    "github.com/codegangsta/cli"
)

var CmdCronList = cli.Command{
    Name: "list",
    Usage: "List the schedule task",
    Description: "The tools provide you to list scheduled task",
    Action: cronList,
    Flags: []cli.Flag{
    },
}

func cronList(cli *cli.Context) error {
    fmt.Println("This is a cron list command")

    return nil
}
