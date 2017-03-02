package mail

import (
    "fmt"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/mail"
)

var CmdMailRender = cli.Command{
    Name: "render",
    Usage: "Test to render the notification content",
    Description: "The tools provide you to show the rendered notification content by dummy data",
    Action: mailRender,
    Flags: []cli.Flag{
    },
}

type Dummy struct {
    Name            string
    Performances    []DummyPerformance
}

type DummyPerformance struct {
    Name    string
    State   string
}

func mailRender(cli *cli.Context) error {
    log.Info("Reading notification mail content ...")
    log.Info("Rendering .....\n")

    dummy := Dummy{
        Name: "Title",
        Performances: []DummyPerformance{
            DummyPerformance{ Name: "A Performance", State: "YES" },
            DummyPerformance{ Name: "B Performance", State: "YES" },
        },
    }

    content, err := mail.RenderMailNotification(dummy)
    if err != nil {
        return err
    }

    fmt.Println(content)

    return nil
}
