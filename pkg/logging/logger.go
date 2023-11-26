package logging

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		logger: l,
	}
}

func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(message)
}

func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(message)
}

func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(message)
}

func (l *Logger) Fatal(message string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(message)
}
