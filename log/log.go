package log

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/util"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (

	// qualified package name, cached at first use
	logPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
	FormatTimestamp        = "2006-01-02 15:04:05"
)

func init() {
	// start at the bottom of the stack before the package-name cache is primed
	minimumCallerDepth = 1
}

var (
	std = logrus.New()
)

type Config struct {
	Level string
}

func Setup() error {

	lumberjackLogger := lumberjack.Logger{}
	if err := gonfig.Instance().UnmarshalKey("gore.logger", &lumberjackLogger); err != nil {
		return err
	}
	multiWriter := io.MultiWriter(os.Stdout, &lumberjackLogger)
	std.SetOutput(multiWriter)

	setLogFormatter()

	level := gonfig.Instance().GetString("gore.logger.level")
	if len(level) > 0 {
		SetLogLevel(level)
	} else {
		SetLogLevel(logrus.TraceLevel.String())
	}

	Infof("Current log filename: %s", lumberjackLogger.Filename)
	Infof("Current log maxage: %v", lumberjackLogger.MaxAge)

	return nil
}

// GetLogWriter 获取日志输出Writer
func GetLogWriter() io.Writer {
	return std.Out
}

func StandardLogger() *logrus.Logger {
	return std
}

func GetLevel() logrus.Level {
	return std.GetLevel()
}

func WithContext(c context.Context) *logrus.Entry {
	return WithField(util.RequestIDContextKey, util.GetRequestID(c))
}

func WithError(err error) *logrus.Entry {
	return getLogEntry().WithError(err)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return getLogEntry().WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return getLogEntry().WithFields(fields)
}

func ErrorE(err error) {
	getLogEntry().WithError(err).Error()
}

func WarnE(err error) {
	getLogEntry().WithError(err).Warn()
}

func Fatal(args ...interface{}) {
	getLogEntry().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	getLogEntry().Fatalf(format, args...)
}

func Error(args ...interface{}) {
	getLogEntry().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	getLogEntry().Errorf(format, args...)
}

func Warn(args ...interface{}) {
	getLogEntry().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	getLogEntry().Warnf(format, args...)
}

func Info(args ...interface{}) {
	getLogEntry().Info(args...)
}

func Infof(format string, args ...interface{}) {
	getLogEntry().Infof(format, args...)
}

func Debug(args ...interface{}) {
	getLogEntry().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	getLogEntry().Debugf(format, args...)
}

func Trace(args ...interface{}) {
	getLogEntry().Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	getLogEntry().Tracef(format, args...)
}

func Print(args ...interface{}) {
	getLogEntry().Print(args...)
}
func Printf(format string, args ...interface{}) {
	getLogEntry().Printf(format, args...)
}
func Println(args ...interface{}) {
	getLogEntry().Println(args...)
}

func ErrorCE(c context.Context, err error) {
	WithContext(c).WithError(err).Error()
}

func WarnCE(c context.Context, err error) {
	WithContext(c).WithError(err).Warn()
}

func FatalC(c context.Context, args ...interface{}) {
	WithContext(c).Fatal(args...)
}

func FatalfC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Fatalf(format, args...)
}

func ErrorC(c context.Context, args ...interface{}) {
	WithContext(c).Error(args...)
}

func ErrorfC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Errorf(format, args...)
}

func WarnC(c context.Context, args ...interface{}) {
	WithContext(c).Warn(args...)
}

func WarnfC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Warnf(format, args...)
}

func InfoC(c context.Context, args ...interface{}) {
	WithContext(c).Info(args...)
}

func InfofC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Infof(format, args...)
}

func DebugC(c context.Context, args ...interface{}) {
	WithContext(c).Debug(args...)
}

func DebugfC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Debugf(format, args...)
}

func TraceC(c context.Context, args ...interface{}) {
	WithContext(c).Trace(args...)
}

func TracefC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Tracef(format, args...)
}

func PrintC(c context.Context, args ...interface{}) {
	WithContext(c).Print(args...)
}
func PrintfC(c context.Context, format string, args ...interface{}) {
	WithContext(c).Printf(format, args...)
}
func PrintlnC(c context.Context, args ...interface{}) {
	WithContext(c).Println(args...)
}

func setLogFormatter() {
	std.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  FormatTimestamp,
		CallerPrettyfier: callerPrettyfier,
		DisableQuote:     true,
	},
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

func getCurDirName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	dirNames := strings.Split(dir, string(os.PathSeparator))
	if len(dirNames) == 0 {
		return "", errors.New("current dir name is empty")
	}
	return dirNames[len(dirNames)-1], nil
}

func getLogEntry() *logrus.Entry {
	frame := getCaller()
	file := shortFile(frame.File)
	file = fmt.Sprintf("%s:%d", file, frame.Line)
	return std.WithField("file", file)
}

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != logPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
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
