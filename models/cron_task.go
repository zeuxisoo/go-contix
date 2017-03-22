package models

type CronTask struct {
    Performances []CronTaskPerformance  `yaml:"performances"`
    Mail         CronTaskMail           `yaml:"mail"`
    UserAgents   []CronTaskUserAgent    `yaml:"user_agents"`
    Telegram     CronTaskTelegram       `yaml:"telegram"`
}

type CronTaskPerformance struct {
    Id          int                         `yaml:"id"`
    Schedule    string                      `yaml:"schedule"`
    Remark      string                      `yaml:"remark"`
    Enable      bool                        `yaml:"enable"`
    Timeout     int                         `yaml:"timeout"`
    Proxy       CronTaskPerformanceProxy    `yaml:"proxy"`
}

type CronTaskPerformanceProxy struct {
    Enable  bool    `yaml:"enable"`
    Method  string  `yaml:"method"`
    Server  string  `yaml:"server"`
}

type CronTaskMail struct {
    Sender      string               `yaml:"sender"`
    Recipient   string               `yaml:"recipient"`
    Subject     string               `yaml:"subject"`
    Mailgun     CronTaskMailMailgun  `yaml:"mailgun"`
}

type CronTaskMailMailgun struct {
    Domain string `yaml:"domain"`
    ApiKey string `yaml:"api_key"`
}

type CronTaskUserAgent struct {
    Name    string  `yaml:"name"`
    Agent   string  `yaml:"agent"`
}

type CronTaskTelegram struct {
    Enable  bool                        `yaml:"enable"`
    Token   string                      `yaml:"token"`
    ChatIds []CronTaskTelegramChatId    `yaml:"chat_ids"`
}

type CronTaskTelegramChatId struct {
    Name string `yaml:"name"`
    Code int64  `yaml:"code"`
}
