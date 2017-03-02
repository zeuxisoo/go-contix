package models

type MailNotificationData struct {
    Name            string
    Performances    []MailNotificationDataPerformance
}

type MailNotificationDataPerformance struct {
    Name    string
    State   string
}
