package models

type CronTask struct {
    Performances []CronTaskPerformance  `yaml:"performances"`
    Mail         CronTaskMail           `yaml:"mail"`
    UserAgents   []string               `yaml:"user_agents"`
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
