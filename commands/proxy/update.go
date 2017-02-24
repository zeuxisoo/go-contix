package proxy

import (
    "os"
    "fmt"
    "bytes"
    "io"
    "bufio"
    "time"
    "net/http"

    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"

    "github.com/zeuxisoo/go-contix/configs"
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
    proxyList, err := readProxyFetchFile()
    if err != nil {
        return err
    }

    request := gorequest.New()
    passedProxyChannel := make(chan string)

    go func() {
        for _, proxy := range proxyList {
            if proxy == "" {
                continue
            }

            response, _, errs := request.
                Proxy(proxy).
                Get("http://httpbin.org/ip").
                Retry(1, 2 * time.Second, http.StatusBadRequest, http.StatusInternalServerError). // 1 times each 2 second
                End()
            if errs != nil {
                continue
            }

            if response.StatusCode == 200 {
                passedProxyChannel <- proxy
            }
        }

        close(passedProxyChannel)
    }()

    var passedProxyList []string
    for proxy := range passedProxyChannel {
        fmt.Printf("Validated: %s\n", proxy)

        passedProxyList = append(passedProxyList, proxy)
    }

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
                    continue
                }
            }
        }
    }

    return nil
}

func readProxyFetchFile() ([]string, error) {
    file, err := os.Open(configs.ProxyFetchFilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    var buffer bytes.Buffer
    var lines []string
    for {
        line, prefix, err := reader.ReadLine()
        if err != nil {
            break
        }

        buffer.Write(line)

        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }

    if err == io.EOF {
        return nil, err
    }

    return lines, nil
}
