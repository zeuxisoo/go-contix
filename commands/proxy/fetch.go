package proxy

import (
    "os"
    "fmt"
    "strings"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"

    proxySite "github.com/zeuxisoo/go-contix/proxy/site"
)

const (
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
        cli.StringFlag{
            Name:  "log-file-path",
            Usage: "Where do you want to save the log file?",
            Value: "",
        },
    },
}

func proxyFetch(ctx *cli.Context) error {
    site := ctx.String("site")

    logFilePath := ctx.String("log-file-path")
    if logFilePath != "" {
        log.SetLogFilePath(logFilePath)
    }

    log.Infof("Proxy site: %s", site)
    log.Infof("Starting .....")

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

    log.Infof("Totoal proxy site: %d", len(proxyList))

    if len(proxyList) > 0 {
        if err := os.Remove(configs.ProxyFetchFilePath); err != nil {
            return err
        }

        file, err := os.OpenFile(configs.ProxyFetchFilePath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
        if err != nil {
            return err
        }
        defer file.Close()

        for _, proxy := range proxyList {
            ipAndPort := fmt.Sprintf("%s://%s:%s", proxy.Protocol, proxy.IP, proxy.Port)
            ipAndPortWithNewLine := fmt.Sprintf("%s\n", ipAndPort)

            if _, err = file.WriteString(ipAndPortWithNewLine); err != nil {
                log.Infof("%s ..... %s", "✘", ipAndPort)
                continue
            }

            log.Infof("%s ..... %s", "✔", ipAndPort)
        }
    }

    return nil
}

