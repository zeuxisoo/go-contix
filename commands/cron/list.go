package cron

import (
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
    table.SetHeader([]string{ "ID", "Schedule", "Remark", "Enable" })

    for _, ticket := range cronTask.Tickets {
        table.Append([]string{
            strconv.Itoa(ticket.Id),
            ticket.Schedule,
            ticket.Remark,
            toYesOrNo(ticket.Enable),
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
