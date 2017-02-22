package commands

import (
    "fmt"

    "github.com/codegangsta/cli"
)

var CmdProxyUpdate = cli.Command{
    Name: "update",
    Usage: "Update exists proxy data",
    Description: "The tools provide you to update exists proxy data",
    Action: proxyUpdate,
    Flags: []cli.Flag{
    },
}

func proxyUpdate(cli *cli.Context) error {
    fmt.Println("This is a proxy update command")

    return nil
}
