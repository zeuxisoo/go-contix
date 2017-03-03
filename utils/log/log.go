package log

import (
    "github.com/Sirupsen/logrus"
    "github.com/mattn/go-colorable"
)

var (
    log = logrus.New()
)

func init() {
    log.Formatter = &logrus.TextFormatter{
        ForceColors: true,
        TimestampFormat: "2006-01-02 15:04:05",
        FullTimestamp: true,
    }

    log.Level = logrus.DebugLevel

    log.Out = colorable.NewColorableStdout()
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}
