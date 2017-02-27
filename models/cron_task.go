package models

type CronTask struct {
    Tickets []CronTaskTicket `yaml:"tickets"`
}

type CronTaskTicket struct {
    Id          int     `yaml:"id"`
    Schedule    string  `yaml:"schedule"`
    Remark      string  `yaml:"remark"`
    Enable      bool    `yaml:"enable"`
}
