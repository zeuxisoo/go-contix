package mail

import (
    "bytes"
    "text/template"

    "github.com/zeuxisoo/go-contix/configs"
)

func RenderMailNotification(data interface{}) (string, error) {
    templateFile, err := template.ParseFiles(configs.MailNotificationFilePath)
    if err != nil {
        return "", err
    }

    var content bytes.Buffer
    err = templateFile.Execute(&content, data)
    if err != nil {
        return "", err
    }

    return content.String(), err
}
