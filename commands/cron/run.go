package cron

import (
    "strconv"
    "math/rand"

    "github.com/codegangsta/cli"
    "github.com/robfig/cron"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
    "github.com/zeuxisoo/go-contix/utils/checker"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/file"
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
                if err := checkPerformanceStateTask(task); err != nil {
                    log.Infof("✘ ... %s", err)
                }
            })
        }
    }

    cronTab.Start()
    select{}

    return nil
}

func checkPerformanceStateTask(task models.CronTaskTicket) error {
    log.Infof("Name: %s", task.Remark)
    log.Infof("Checking .....")

    lines, err := file.ReadByLines(configs.ProxyPoolFilePath)
    if err != nil {
        log.Infof("✘ ... Cannot read the proxy pool file: %s", configs.ProxyPoolFilePath)
        return err
    }

    log.Infof("Proxy pool size: %d", len(lines))
    log.Infof("Shuffling .....")

    proxy := ""
    if len(lines) > 0 {
        proxy = lines[rand.Intn(len(lines))]
    }

    if proxy == "" {
        log.Infof("✘ ... No more proxy can pick")
    }else{
        log.Infof("✔ ... Picked proxy: %s", proxy)
    }

    performanceStateChecker := checker.NewPerformanceStateChecker().
        SetPerformanceId(strconv.Itoa(task.Id)).
        SetProxy(proxy)

    performances, err := performanceStateChecker.GetPerformanceList()
    if err != nil {
        log.Infof("✘ ... Cannnot get the performance list")
        return err
    }

    for _, performance := range performances {
        if performance.Status == "AVAILABLE" || performance.Status == "LIMIT" {
            log.Infof("✔ ... Tickets are %s", performance.Status)
        }else{
            log.Infof("✘ ... Tickets are %s", performance.Status)
        }
    }

    return nil
}
