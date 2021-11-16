package internal

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() (*logrus.Logger, error) {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	logger := &logrus.Logger{
		Out:   f,
		Level: logrus.InfoLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	}
	return logger, nil
}
