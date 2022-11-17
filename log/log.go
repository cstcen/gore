package log

import (
	"context"
	"fmt"
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
	defaultLevel        = LevelDebug
	defaultLevelName    = LevelDebug.Name()
	defaultRequestIdKey = "X-Request-ID"
)

func MustRequestID(c context.Context) string {
	id, ok := c.Value(defaultRequestIdKey).(string)
	if !ok {
		return ""
	}
	return id
}

func SetRequestIdKey(key string) {
	defaultRequestIdKey = key
}

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
	Logger *lumberjack.Logger
	Level  string
}

func Setup(cfg *Config) error {
	if cfg == nil {
		return errors.New("log config not found")
	}

	log.SetPrefix("[GORE] ")
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)

	if cfg.Logger != nil {
		multiWriter := io.MultiWriter(os.Stdout, cfg.Logger)
		log.SetOutput(multiWriter)
	} else {
		log.SetOutput(os.Stdout)
	}

	level, err := ParseLevel(cfg.Level)
	if err != nil {
		log.Println(err.Error())
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
		return LevelDebug, errors.New("invalid log level")
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
	log.Printf("[%s] [%s] %s", MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
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
	log.Printf("[%s] [%s] %s", MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
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
	log.Printf("[%s] [%s] %s", MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
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
	log.Printf("[%s] [%s] %s", MustRequestID(ctx), defaultLevelName, fmt.Sprintf(format, v))
}

func Panicf(format string, v ...any) {
	log.Panicf(format, v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}
