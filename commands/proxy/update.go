package proxy

import (
    "fmt"
    "os"
    "bytes"
    "io"
    "bufio"

    "github.com/codegangsta/cli"

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

    fmt.Println(proxyList)

    return nil
}

func readProxyFetchFile() ([]string, error) {
    file, err := os.Open(configs.ProxyFetchFilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte,1024))

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
