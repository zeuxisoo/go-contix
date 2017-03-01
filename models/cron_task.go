package models

type CronTask struct {
    Performances []CronTaskPerformance `yaml:"performances"`
}

type CronTaskPerformance struct {
    Id          int     `yaml:"id"`
    Schedule    string  `yaml:"schedule"`
    Remark      string  `yaml:"remark"`
    Enable      bool    `yaml:"enable"`
    Timeout     int     `yaml:"timeout"`
}
