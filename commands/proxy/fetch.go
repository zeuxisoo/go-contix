package proxy

import (
    "strings"

    "github.com/codegangsta/cli"

    proxySite "github.com/zeuxisoo/go-contix/proxy/site"
)

var CmdProxyFetch = cli.Command{
    Name: "fetch",
    Usage: "Create and fetch remote proxy data",
    Description: "The tools provide you to create and fetch proxy data",
    Action: proxyFetch,
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "site",
            Usage: "What proxy site you want? [Support: gimme]",
            Value: "gimme",
        },
    },
}

func proxyFetch(ctx *cli.Context) error {
    site := ctx.String("site")

    switch strings.ToLower(site) {
        case "gimme":
            proxySite.FetchGimme()
    }

    return nil
}

