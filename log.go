package gocore

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
)

var (
	ErrAppNameEmpty = errors.New("the app name is empty")
)

var (
	std = logrus.New()
)

// SetupLog 设置日志
func SetupLog(appName string) error {
	if "" == appName {
		return ErrAppNameEmpty
	}
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
	std.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func setLogOutput(appName string) error {
	rotateLogs, err := getRotateLogs(appName)
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
	lvl, ok := EnvLevelMap[EnvCurrent]
	if ok {
		std.SetLevel(lvl)
	}
}

func getRotateLogs(appName string) (*rotatelogs.RotateLogs, error) {
	var (
		link          = fmt.Sprintf("/xk5/logs/%s/%s.log", appName, appName)
		path          = fmt.Sprintf("/xk5/logs/%s/%s-%s", appName, appName, "%Y-%m-%d.log")
		maxFiles uint = 30
	)

	if sysType := runtime.GOOS; sysType != "linux" {
		link = ""
	}
	return rotatelogs.New(
		path,
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(link),

		// WithRotationTime设置日志分割的时间,默认：86400s分割一次
		// rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// rotatelogs.WithMaxAge(time.Hour*24),
		// WithRotationCount设置文件清理前最多保存的个数.
		rotatelogs.WithRotationCount(maxFiles),
	)
}
