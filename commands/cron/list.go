package cron

import (
    "os"
    "io/ioutil"
    "strconv"

    "github.com/codegangsta/cli"
    "gopkg.in/yaml.v2"
    "github.com/olekukonko/tablewriter"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
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
    cronTaskFileBytes, err := ioutil.ReadFile(configs.CronTaskFilePath)
    if err != nil {
        return err
    }

    var cronTask models.CronTask
    if err := yaml.Unmarshal(cronTaskFileBytes, &cronTask); err != nil {
        return err
    }

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
