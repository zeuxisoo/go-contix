package commands

import (
    "github.com/codegangsta/cli"
)

var CmdProxy = cli.Command{
    Name: "proxy",
    Usage: "A tools for manage the proxy data",
    Description: "The tools provide you to fetch or update proxy data etc",
    Subcommands: []cli.Command{
        CmdProxyFetch,
        CmdProxyUpdate,
    },
}
