package log

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/util"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
)

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
)

var (
	defaultLevel     = LevelDebug
	defaultLevelName = LevelDebug.Name()
)

func GetLevel() Level {
	return defaultLevel
}

func Default() *log.Logger {
	return log.Default()
}

type Level uint8

func (l Level) Name() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	default:
		return "unknown"
	}
}

type Config struct {
	Level string
}

func Setup() error {

	log.SetPrefix("[GORE] ")
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)

	if gonfig.Instance().GetBool("log") {
		lumberjackLogger := lumberjack.Logger{}
		if err := gonfig.Instance().UnmarshalKey("gore.logger", &lumberjackLogger); err != nil {
			return err
		}
		multiWriter := io.MultiWriter(os.Stdout, &lumberjackLogger)

		log.SetOutput(multiWriter)
	} else {
		log.SetOutput(os.Stdout)
	}

	level, err := ParseLevel(gonfig.Instance().GetString("gore.logger.level"))
	if err != nil {
		return err
	}
	defaultLevel = level

	return nil
}

func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warning":
		return LevelWarning, nil
	case "error":
		return LevelError, nil
	default:
		return LevelDebug, errors.New("invalid level")
	}
}

func Debugf(format string, v ...any) {
	if defaultLevel > LevelDebug {
		return
	}
	log.Printf("[%s] %s", defaultLevelName, fmt.Sprintf(format, v))
}

func DebugCf(ctx context.Context, format string, v ...any) {
	if defaultLevel > LevelDebug {
		return
	}
	log.Printf("[%s] [%s] %s", util.MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
}

func Infof(format string, v ...any) {
	if defaultLevel > LevelInfo {
		return
	}
	log.Printf("[%s] %s", defaultLevelName, fmt.Sprintf(format, v))
}

func InfoCf(ctx context.Context, format string, v ...any) {
	if defaultLevel > LevelInfo {
		return
	}
	log.Printf("[%s] [%s] %s", util.MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
}

func Warningf(format string, v ...any) {
	if defaultLevel > LevelWarning {
		return
	}
	log.Printf("[%s] %s", defaultLevelName, fmt.Sprintf(format, v))
}

func WarningCf(ctx context.Context, format string, v ...any) {
	if defaultLevel > LevelWarning {
		return
	}
	log.Printf("[%s] [%s] %s", util.MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
}

func Errorf(format string, v ...any) {
	if defaultLevel > LevelError {
		return
	}
	log.Printf("[%s] %s", defaultLevelName, fmt.Sprintf(format, v))
}

func ErrorCf(ctx context.Context, format string, v ...any) {
	if defaultLevel > LevelError {
		return
	}
	log.Printf("[%s] [%s] %s", util.MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
}

func Panicf(format string, v ...any) {
	log.Panicf(format, v...)
}

func Fatal(format string, v ...any) {
	log.Fatalf(format, v...)
}
