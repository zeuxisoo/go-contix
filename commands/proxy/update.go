package proxy

import (
    "os"
    "fmt"
    "time"
    "net/http"

    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/file"
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
    type proxyState models.ProxyState

    log.Infof("Loading fetched proxy file ...")

    proxyList, err := readProxyFetchFile()
    if err != nil {
        return err
    }

    log.Infof("Validating fetched proxy list ...")

    request := gorequest.New()

    validateProxyStateChannel  := make(chan string, 100)
    validateProxyResultChannel := make(chan proxyState, 100)

    for workerCount := 0; workerCount <= 3; workerCount++ {
        go func() {
            for proxy := range validateProxyStateChannel {
                if proxy == "" {
                    validateProxyResultChannel <- proxyState{
                        Usable: false,
                        Proxy : "",
                    }

                    continue
                }

                response, _, errs := request.
                    Proxy(proxy).
                    Get("http://httpbin.org/ip").
                    // Connection expire on 3 second
                    Timeout(3000 * time.Millisecond).
                    // Retry 1 times with 3 second when got bad request or inter server error
                    Retry(1, 3 * time.Second, http.StatusBadRequest, http.StatusInternalServerError).
                    End()

                if errs != nil {
                    log.Infof("%s ... %s", "✘", proxy)

                    validateProxyResultChannel <- proxyState{
                        Usable: false,
                        Proxy : proxy,
                    }

                    continue
                }

                if response.StatusCode == 200 {
                    log.Infof("%s ... %s", "✔", proxy)

                    validateProxyResultChannel <- proxyState{
                        Usable: true,
                        Proxy : proxy,
                    }
                }
            }
        }()
    }

    for _, proxy := range proxyList {
        validateProxyStateChannel <- proxy
    }
    close(validateProxyStateChannel)

    var passedProxyList []string
    for i := 0; i < len(proxyList); i++ {
        result := <-validateProxyResultChannel

        if result.Usable == true {
            passedProxyList = append(passedProxyList, result.Proxy)
        }
    }

    log.Infof("Total usable proxy site: %d", len(passedProxyList))
    log.Infof("Updating .....")

    if len(passedProxyList) > 0 {
        if err := os.Remove(configs.ProxyPoolFilePath); err != nil {
            return err
        }

        file, err := os.OpenFile(configs.ProxyPoolFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
        if err != nil {
            return err
        }
        defer file.Close()

        for _, proxy := range passedProxyList {
            if (proxy != "") {
                validatedProxy := fmt.Sprintf("%s\n", proxy)

                if _, err = file.WriteString(validatedProxy); err != nil {
                    log.Infof("%s ... %s", "✘", proxy)
                    continue
                }

                log.Infof("%s ... %s", "✔", proxy)
            }
        }
    }

    return nil
}

func readProxyFetchFile() ([]string, error) {
    lines, err := file.ReadByLines(configs.ProxyFetchFilePath)

    return lines, err
}
