package logger

import (
	"github.com/BounkBU/doonangfangpleng/config"
	"github.com/sirupsen/logrus"
)

func InitLogger(config *config.Config) {
	if config.Env == "prod" {
		timeFormatLayout := "2006-01-02T15:04:05.000Z"
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "log_level",
			},
			TimestampFormat: timeFormatLayout,
		})
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
