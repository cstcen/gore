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

	logLink     string
	logPath     string
	logMaxFiles uint

	// 各环境对应的日志打印等级
	LogLevelMap = map[string]logrus.Level{
		"sdev0":  logrus.DebugLevel,
		"sdev":   logrus.DebugLevel,
		"dev":    logrus.DebugLevel,
		"dev2":   logrus.DebugLevel,
		"dev3":   logrus.DebugLevel,
		"ios":    logrus.DebugLevel,
		"mod":    logrus.DebugLevel,
		"stg":    logrus.DebugLevel,
		"xingk5": logrus.DebugLevel,
		"xk5":    logrus.InfoLevel,
	}
)

var (
	std = logrus.New()
)

// SetupLog 设置日志
func SetupLog(env, appName string) error {
	if "" == appName {
		return ErrAppNameEmpty
	}
	err := setLogOutput(appName)
	if err != nil {
		return err
	}

	std.SetReportCaller(true)

	setLogLevel(env)

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

func SetLogLink(link string) error {
	if "" == link {
		return errors.New("invalid log link")
	}
	logLink = link
	return nil
}

func SetLogPath(path string) error {
	if "" == path {
		return errors.New("invalid log path")
	}
	logPath = path
	return nil
}

func SetMaxFiles(maxFiles uint) error {
	if 0 == maxFiles {
		return errors.New("invalid log max files")
	}
	logMaxFiles = maxFiles
	return nil
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

func setLogLevel(env string) {
	lvl, ok := LogLevelMap[env]
	if ok {
		std.SetLevel(lvl)
	}
}

func getRotateLogs(appName string) (*rotatelogs.RotateLogs, error) {

	if "" == logLink {
		logLink = fmt.Sprintf("/xk5/logs/%s/%s.log", appName, appName)
	}
	if "" == logPath {
		logPath = fmt.Sprintf("/xk5/logs/%s/%s-%s", appName, appName, "%Y-%m-%d.log")
	}
	if 0 == logMaxFiles {
		logMaxFiles = 30
	}

	if sysType := runtime.GOOS; sysType != "linux" {
		logLink = ""
	}
	return rotatelogs.New(
		logPath,
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logLink),

		// WithRotationTime设置日志分割的时间,默认：86400s分割一次
		// rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// rotatelogs.WithMaxAge(time.Hour*24),
		// WithRotationCount设置文件清理前最多保存的个数.
		rotatelogs.WithRotationCount(logMaxFiles),
	)
}
