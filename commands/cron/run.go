package cron

import (
    "fmt"
    "strconv"
    "math/rand"

    "github.com/codegangsta/cli"
    "github.com/robfig/cron"
    "github.com/fatih/color"

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

    yellow  := color.New(color.FgYellow).SprintFunc()
    red     := color.New(color.FgRed).SprintFunc()
    cyan    := color.New(color.FgCyan).SprintFunc()
    green   := color.New(color.FgGreen).SprintFunc()
    cronTab := cron.New()

    for i := 0; i < len(cronTask.Performances); i++ {
        task := cronTask.Performances[i]

        if task.Enable == true {
            cronTab.AddFunc(task.Schedule, func() {
                log.Infof(yellow(fmt.Sprintf("Remark: %s, Checking ...", task.Remark)))

                available, err := checkPerformanceStateTask(i, task);
                if err != nil {
                    log.Infof(red(fmt.Sprintf("%s: %s", task.Remark, err)))
                }

                if available {
                    log.Infof(green(fmt.Sprintf("Remark: %s, State: ✔", task.Remark)))
                }else{
                    log.Infof(cyan(fmt.Sprintf("Remark: %s, State: ✘", task.Remark)))
                }
            })
        }
    }

    cronTab.Start()
    select{}

    return nil
}

func checkPerformanceStateTask(id int, task models.CronTaskPerformance) (bool, error) {
    lines, err := file.ReadByLines(configs.ProxyPoolFilePath)
    if err != nil {
        return false, err
    }

    proxy := ""
    if len(lines) > 0 {
        proxy = lines[rand.Intn(len(lines))]
    }

    performanceStateChecker := checker.NewPerformanceStateChecker().
        SetPerformanceId(strconv.Itoa(task.Id)).
        SetProxy(proxy).
        SetTimeout(task.Timeout)

    performances, err := performanceStateChecker.GetPerformanceList()
    if err != nil {
        return false, err
    }

    isAvailable := false
    for _, performance := range performances {
        if performance.Status == "AVAILABLE" || performance.Status == "LIMIT" {
            isAvailable = true
            break
        }
    }

    return isAvailable, nil
}
