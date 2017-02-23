package proxy

import (
    "os"
    "fmt"
    "strings"

    "github.com/codegangsta/cli"

    proxySite "github.com/zeuxisoo/go-contix/proxy/site"
)

const (
    FetchFilePath = "data/proxy-fetch.txt"

    FreeProxyList = "free-proxy-list"
    NyLoner       = "nyloner"
    XiCiDaiLi     = "xicidaili"
    Gimme         = "gimme"
)

var CmdProxyFetch = cli.Command{
    Name: "fetch",
    Usage: "Create and fetch remote proxy data",
    Description: "The tools provide you to create and fetch proxy data",
    Action: proxyFetch,
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "site",
            Usage: fmt.Sprintf("What proxy site you want? [Support: %s]", strings.Join([]string{
                FreeProxyList,
                NyLoner,
                XiCiDaiLi,
                Gimme,
            }, ",")),
            Value: FreeProxyList,
        },
    },
}

func proxyFetch(ctx *cli.Context) error {
    site := ctx.String("site")

    var theProxySite proxySite.Contract
    switch strings.ToLower(site) {
        case FreeProxyList:
            theProxySite = new(proxySite.FreeProxyListProxySite)
        case NyLoner:
            theProxySite = new(proxySite.NyLonerProxySite)
        case XiCiDaiLi:
            theProxySite = new(proxySite.XiCiDaiLiProxySite)
        case Gimme:
            theProxySite = new(proxySite.GimmeProxySite)
    }

    proxyList, err := theProxySite.Fetch();
    if err != nil {
        return err
    }

    if len(proxyList) > 0 {
        if err := os.Remove(FetchFilePath); err != nil {
            return err
        }

        file, err := os.OpenFile(FetchFilePath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
        if err != nil {
            return err
        }
        defer file.Close()

        for _, proxy := range proxyList {
            ipAndPort := fmt.Sprintf("%s://%s:%s\n", proxy.Protocol, proxy.IP, proxy.Port)

            if _, err = file.WriteString(ipAndPort); err != nil {
                continue
            }
        }
    }

    return nil
}

