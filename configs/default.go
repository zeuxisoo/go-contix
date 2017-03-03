package configs

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/zeuxisoo/go-contix/utils/log"
)

var (
    ApplicationPath             string

    ProxyFetchFilePath          string
    ProxyPoolFilePath           string

    CronTaskFilePath            string

    MailNotificationFilePath    string
)

func init() {
    var err error
    if ApplicationPath, err = executePath(); err != nil {
        log.Fatalf("Fail to get the application path, error = %v", err)
    }

    ApplicationPath = strings.Replace(ApplicationPath, "\\", "/", -1)

    appDirectory, err := applicationDirectory()
    if err != nil {
        log.Fatalf("Fail to get the application directory, error = %v", err)
    }

    ProxyFetchFilePath = appDirectory + "/data/proxy-fetch.txt"
    ProxyPoolFilePath  = appDirectory + "/data/proxy-pool.txt"

    CronTaskFilePath = appDirectory + "/data/cron-task.yaml"

    MailNotificationFilePath = appDirectory + "/data/mail-notification.txt"
}

func executePath() (string, error) {
    file, err := exec.LookPath(os.Args[0])
    if err != nil {
        return "", err
    }

    return filepath.Abs(file)
}

func applicationDirectory() (string, error) {
    i := strings.LastIndex(ApplicationPath, "/")
    if i == -1 {
        return ApplicationPath, nil
    }

    return ApplicationPath[:i], nil
}
