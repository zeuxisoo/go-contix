package cron

import (
    "fmt"
    "os"
    "strconv"

    "github.com/codegangsta/cli"
    "github.com/olekukonko/tablewriter"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/utils/log"
)

var CmdCronList = cli.Command{
    Name: "list",
    Usage: "List the schedule task",
    Description: "The tools provide you to list scheduled task",
    Action: cronList,
    Flags: []cli.Flag{
    },
}

func cronList(cli *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    log.Info("Rendering .....\n")

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{ "ID", "Schedule", "Remark", "Enable", "Proxy" })

    for _, performance := range cronTask.Performances {
        table.Append([]string{
            strconv.Itoa(performance.Id),
            performance.Schedule,
            performance.Remark,
            toYesOrNo(performance.Enable),
            fmt.Sprintf(
                "%s,%s,%s",
                toYesOrNo(performance.Proxy.Enable),
                performance.Proxy.Method,
                toNAString(performance.Proxy.Server),
            ),
        })
    }

    table.Render()

    return nil
}

func toYesOrNo(enable bool) string {
    if enable {
        return "✔"
    }else{
        return "✘"
    }
}

func toNAString(text string) string {
    if text == "" {
        return "n/a"
    }

    return text
}
