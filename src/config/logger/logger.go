package config_logger

import (
	config_general "go-template/src/config/general"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(appConfig config_general.AppConfig) *logrus.Logger {
	logger := logrus.New()
	today := time.Now().Format("020106") // DDMMYY

	cwd, _ := os.Getwd()
	logPath := filepath.Join(cwd, "log", appConfig.AppName+"_"+today+".log")

	logger.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	})

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
