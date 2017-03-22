package configs

import (
    "fmt"
    "strings"
    "io/ioutil"

    "gopkg.in/yaml.v2"

    "github.com/zeuxisoo/go-contix/models"
)

func LoadCronTask() (models.CronTask, error) {
    var cronTask models.CronTask

    cronTaskFileBytes, err := ioutil.ReadFile(CronTaskFilePath)
    if err != nil {
        return cronTask, err
    }

    if err := yaml.Unmarshal(cronTaskFileBytes, &cronTask); err != nil {
        return cronTask, err
    }

    return cronTask, nil
}

func ConvertChatIds(chatIds []models.CronTaskTelegramChatId) string {
    var temp []string
    for _, chatId := range chatIds {
        temp = append(temp, fmt.Sprintf("%s (%d)", chatId.Name, chatId.Code))
    }

    result := strings.Join(temp, ", ")

    if len(result) <= 0 {
        result = "n/a"
    }

    return result
}

func ConvertEnableStatus(status bool) string {
    if status == true {
        return "✔"
    }else{
        return "✘"
    }
}
