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
        cli.BoolFlag{
            Name:  "disable",
            Usage: "Is it show diabled task only?",
        },
    },
}

func cronList(ctx *cli.Context) error {
    log.Info("Reading cron task file ...")

    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    var rows [][]string
    for _, performance := range cronTask.Performances {
        // Show disabled task only when --disable is assigned
        if ctx.Bool("disable") == true && performance.Enable == false {
            rows = append(rows, []string{
                strconv.Itoa(performance.Id),
                performance.Schedule,
                performance.Remark,
                configs.ConvertEnableStatus(performance.Enable),
                fmt.Sprintf(
                    "%s,%s,%s",
                    configs.ConvertEnableStatus(performance.Proxy.Enable),
                    performance.Proxy.Method,
                    toNAString(performance.Proxy.Server),
                ),
            })
        }

        // Show all task by default when --disable is not assigned
        if ctx.Bool("disable") == false {
            rows = append(rows, []string{
                strconv.Itoa(performance.Id),
                performance.Schedule,
                performance.Remark,
                configs.ConvertEnableStatus(performance.Enable),
                fmt.Sprintf(
                    "%s,%s,%s",
                    configs.ConvertEnableStatus(performance.Proxy.Enable),
                    performance.Proxy.Method,
                    toNAString(performance.Proxy.Server),
                ),
            })
        }
    }

    log.Info("Rendering .....\n")

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{ "ID", "Schedule", "Remark", "Enable", "Proxy" })

    if len(rows) > 0 {
        table.AppendBulk(rows)
        table.SetAlignment(tablewriter.ALIGN_LEFT)
    }else{
        table.Append([]string{ "--", "--", "No any related cron tasks", "--", "--" })
        table.SetAlignment(tablewriter.ALIGN_CENTER)
    }

    table.SetRowLine(true)
    table.Render()

    return nil
}

func toNAString(text string) string {
    if text == "" {
        return "n/a"
    }

    return text
}
