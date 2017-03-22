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

func ConvertChatIds(chatIds []int64) string {
    result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(chatIds)), ", "), "[]")

    if len(result) <= 0 {
        result = "n/a"
    }

    return result
}
