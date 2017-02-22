package proxy

import (
    "fmt"

    "github.com/codegangsta/cli"
)

var CmdProxyFetch = cli.Command{
    Name: "fetch",
    Usage: "Create and fetch remote proxy data",
    Description: "The tools provide you to create and fetch proxy data",
    Action: proxyFetch,
    Flags: []cli.Flag{
    },
}

func proxyFetch(ctx *cli.Context) error {
    fmt.Println("This is a proxy fetch command")

    return nil
}
