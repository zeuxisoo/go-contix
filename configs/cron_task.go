package configs

import (
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
