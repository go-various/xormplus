package xormplus

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"xorm.io/core"
)

type Logger struct {
	hclog.Logger
}

func DefaultLogger() *Logger {
	return &Logger{Logger: hclog.Default().Named("xormplus")}
}

func NewLogger(logger hclog.Logger) *Logger {
	return &Logger{Logger: logger.Named("xormplus")}
}

func (l *Logger) Debug(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if len(v) == 1 {
		l.Logger.Debug(v[0].(string))
	}
	l.Logger.Debug(v[0].(string), v[1:]...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if len(v) == 1 {
		l.Logger.Error(v[0].(string))
	}
	l.Logger.Error(v[0].(string), v[1:]...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if len(v) == 1 {
		l.Logger.Info(v[0].(string))
	}
	l.Logger.Info(v[0].(string), v[1:]...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	panic("implement me")
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, v...))
}

func (l *Logger) Level() core.LogLevel {
	return core.LOG_INFO
}

func (l *Logger) SetLevel(lv core.LogLevel) {
}

func (l *Logger) ShowSQL(show ...bool) {
}

func (l *Logger) IsShowSQL() bool {
	return true
}
