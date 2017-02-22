package commands

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/commands/proxy"
)

var CmdProxy = cli.Command{
    Name: "proxy",
    Usage: "A tools for manage the proxy data",
    Description: "The tools provide you to fetch or update proxy data etc",
    Subcommands: []cli.Command{
        proxy.CmdProxyFetch,
        proxy.CmdProxyUpdate,
    },
}
