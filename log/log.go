package log

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/util"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
)

const (
	FormatTimestamp = "2006-01-02 15:04:05"
)

var (
	std = logrus.StandardLogger()
)

type Config struct {
	Level string
}

func Setup() error {

	if gonfig.Instance().GetBool("log") {
		lumberjackLogger := lumberjack.Logger{}
		if err := gonfig.Instance().UnmarshalKey("gore.logger", &lumberjackLogger); err != nil {
			return err
		}
		std.SetOutput(io.MultiWriter(os.Stdout, &lumberjackLogger))
		std.Infof("Current log filename: %s", lumberjackLogger.Filename)
		std.Infof("Current log maxage: %v", lumberjackLogger.MaxAge)
	} else {
		std.SetOutput(io.MultiWriter(os.Stdout))
	}

	setLogFormatter()

	level := gonfig.Instance().GetString("gore.logger.level")
	if len(level) > 0 {
		SetLogLevel(level)
	} else {
		SetLogLevel(logrus.TraceLevel.String())
	}

	return nil
}

func StandardLogger() *logrus.Logger {
	return std
}

func GetLevel() logrus.Level {
	return std.GetLevel()
}

func WithContext(c context.Context) *logrus.Entry {
	return std.WithField(util.RequestIDContextKey, util.GetRequestID(c))
}

func setLogFormatter() {
	std.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  FormatTimestamp,
		CallerPrettyfier: callerPrettyfier,
		DisableQuote:     true},
	)
}

func SetLogLevel(lvl string) {
	if len(lvl) == 0 {
		return
	}
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return
	}
	std.SetLevel(level)
}

func shortFile(file string) string {
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
	return file
}

func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	file = shortFile(frame.File)
	file = fmt.Sprintf("%s:%d", file, frame.Line)
	return
}
