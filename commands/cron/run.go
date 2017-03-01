package cron

import (
    "fmt"
    "strconv"

    "github.com/codegangsta/cli"
    "github.com/robfig/cron"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
    "github.com/zeuxisoo/go-contix/utils/checker"
)

var CmdCronRun = cli.Command{
    Name: "run",
    Usage: "Run the schedule task",
    Description: "The tools provide you to run scheduled task",
    Action: cronRun,
    Flags: []cli.Flag{
    },
}

func cronRun(cli *cli.Context) error {
    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    cronTab := cron.New()

    for i := 0; i < len(cronTask.Tickets); i++ {
        task := cronTask.Tickets[i]

        if task.Enable == true {
            cronTab.AddFunc(task.Schedule, func() {
                checkTicketStateTask(task)
            })
        }
    }

    cronTab.Start()
    select{}

    return nil
}

func checkTicketStateTask(task models.CronTaskTicket) {
    fmt.Println(task.Remark)

    ticketStateChecker := checker.NewTicketStateChecker().
        SetTicketId(strconv.Itoa(task.Id)).
        SetProxy("")

    performances, err := ticketStateChecker.GetPerformanceList()
    if err != nil {

    }

    fmt.Println(performances)
}
