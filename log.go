package gocore

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
)

var (
	std = logrus.New()
)

// LogSetup 设置日志
func LogSetup(appName string) error {
	err := setLogOutput(appName)
	if err != nil {
		return err
	}

	std.SetReportCaller(true)

	setLogLevel()

	setLogFormatter()

	return nil
}

// GetLogWriter 获取日志输出Writer
func GetLogWriter() io.Writer {
	return std.Out
}

func WithField(key string, value interface{}) *logrus.Entry {
	return std.WithField(key, value)
}

func ErrorE(err error) {
	std.WithError(err).Error()
}

func WarnE(err error) {
	std.WithError(err).Warn()
}

func Error(args ...interface{}) {
	std.Error(args)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args)
}

func Warn(args ...interface{}) {
	std.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args)
}

func Info(args ...interface{}) {
	std.Info(args)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args)
}

func Debug(args ...interface{}) {
	std.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args)
}

func setLogOutput(appName string) error {
	rotateLogs, err := GetRotateLogs(appName)
	if err != nil {
		return err
	}
	std.SetOutput(rotateLogs)
	return nil
}

func setLogFormatter() {
	std.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: FormatTimestamp,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			file = frame.File
			short := file
			var count int
			maxCount := 1
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					if count >= maxCount {
						short = file[i+1:]
						break
					}
					count++
				}
			}
			file = short
			file = fmt.Sprintf("%s:%d", file, frame.Line)
			return
		}})
}

func setLogLevel() {
	lvl, ok := EnvLevelMap[CurrentEnv]
	if ok {
		std.SetLevel(lvl)
	}
}
