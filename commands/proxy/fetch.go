package proxy

import (
    "os"
    "fmt"
    "strings"

    "github.com/codegangsta/cli"

    proxySite "github.com/zeuxisoo/go-contix/proxy/site"
)

const (
    Gimme = "gimme"
)

var CmdProxyFetch = cli.Command{
    Name: "fetch",
    Usage: "Create and fetch remote proxy data",
    Description: "The tools provide you to create and fetch proxy data",
    Action: proxyFetch,
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "site",
            Usage: fmt.Sprintf("What proxy site you want? [Support: %s]", Gimme),
            Value: Gimme,
        },
    },
}

func proxyFetch(ctx *cli.Context) error {
    site := ctx.String("site")

    var theProxySite proxySite.Contract
    switch strings.ToLower(site) {
        case Gimme:
            theProxySite = new(proxySite.GimmeProxySite)
    }

    proxyList, err := theProxySite.Fetch();
    if err != nil {
        return err
    }

    if len(proxyList) > 0 {
        file, err := os.OpenFile("data/proxy-fetch.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
        if err != nil {
            return err
        }
        defer file.Close()

        for _, proxy := range proxyList {
            if _, err = file.WriteString(proxy); err != nil {
                continue
            }
        }
    }

    return nil
}

