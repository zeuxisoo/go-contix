package cron

import (
    "os"
    "os/signal"
    "strconv"
    "strings"
    "syscall"
    "math/rand"

    "github.com/codegangsta/cli"
    "github.com/robfig/cron"
    "github.com/fatih/color"

    "github.com/zeuxisoo/go-contix/configs"
    "github.com/zeuxisoo/go-contix/models"
    "github.com/zeuxisoo/go-contix/utils/checker"
    "github.com/zeuxisoo/go-contix/utils/log"
    "github.com/zeuxisoo/go-contix/utils/file"
    "github.com/zeuxisoo/go-contix/utils/mail"
)

var CmdCronRun = cli.Command{
    Name: "run",
    Usage: "Run the schedule task",
    Description: "The tools provide you to run scheduled task",
    Action: cronRun,
    Flags: []cli.Flag{
    },
}

func cronRun(ctx *cli.Context) error {
    cronTask, err := configs.LoadCronTask()
    if err != nil {
        return err
    }

    yellow  := color.New(color.FgYellow).SprintfFunc()
    red     := color.New(color.FgRed).SprintfFunc()
    cyan    := color.New(color.FgCyan).SprintfFunc()
    green   := color.New(color.FgGreen).SprintfFunc()
    cronTab := cron.New()

    for i := 0; i < len(cronTask.Performances); i++ {
        task := cronTask.Performances[i]

        if task.Enable == true {
            cronTab.AddFunc(task.Schedule, func() {
                log.Infof(yellow("Remark: %s, Checking ...", task.Remark))

                available, performances, err := checkPerformanceStateTask(i, cronTask, task);
                if err != nil {
                    log.Infof(red("Remark: %s, Check: ✘, error: %v", task.Remark, err))
                }

                if available {
                    log.Infof(green("Remark: %s, Status: ✔", task.Remark))

                    if err := sendMailNotification(cronTask, task, performances); err != nil {
                        log.Infof(red("Remark: %s, Mail: ✘, error: %v", task.Remark, err))
                    }else{
                        log.Infof(green("Remark: %s, Mail: ✔", task.Remark))
                    }
                }else{
                    log.Infof(cyan("Remark: %s, Status: ✘", task.Remark))
                }
            })
        }
    }

    cronTab.Start()

    signalChannel := make(chan os.Signal)
    signal.Notify(signalChannel, os.Kill, os.Interrupt, syscall.SIGTERM)
    <-signalChannel

    return nil
}

func checkPerformanceStateTask(id int, cronTask models.CronTask, task models.CronTaskPerformance) (bool, []models.PerformanceList, error) {
    lines, err := file.ReadByLines(configs.ProxyPoolFilePath)
    if err != nil {
        return false, []models.PerformanceList{}, err
    }

    proxy := ""
    if task.Proxy.Enable && len(lines) > 0 {
        switch strings.ToLower(task.Proxy.Method) {
            case "pool":
                proxy = lines[rand.Intn(len(lines))]
            case "custom":
                proxy = task.Proxy.Server
        }
    }

    performanceStateChecker := checker.NewPerformanceStateChecker().
        SetPerformanceId(strconv.Itoa(task.Id)).
        SetProxy(proxy).
        SetTimeout(task.Timeout).
        SetUserAgents(cronTask.UserAgents)

    performanceList, err := performanceStateChecker.GetPerformanceList()
    if err != nil {
        return false, performanceList, err
    }

    isAvailable := false
    for _, performance := range performanceList {
        if performance.Status == "AVAILABLE" || performance.Status == "LIMIT" {
            isAvailable = true
            break
        }
    }

    return isAvailable, performanceList, nil
}

func sendMailNotification(cronTask models.CronTask, task models.CronTaskPerformance, performanceList []models.PerformanceList) error {
    var templatePerformances []models.MailNotificationDataPerformance
    for _, performance := range performanceList {
        templatePerformances = append(
            templatePerformances,
            models.MailNotificationDataPerformance{
                Name: performance.Name,
                State: performance.Status,
            },
        )
    }

    templateData := models.MailNotificationData{
        Name: task.Remark,
        Performances: templatePerformances,
    }

    mailNotificationContent, err := mail.RenderMailNotification(templateData)
    if err != nil {
        return err
    }

    _, _, err = mail.NewMailgun(cronTask.Mail.Mailgun.Domain, cronTask.Mail.Mailgun.ApiKey, "",).
        SetSender(cronTask.Mail.Sender).
        SetRecipient(cronTask.Mail.Recipient).
        SetSubject(cronTask.Mail.Subject).
        SetContent(mailNotificationContent).
        Send()
    if err != nil {
        return err
    }

    return nil
}
