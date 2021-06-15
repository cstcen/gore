package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
	ErrAppNameEmpty = errors.New("the app name is empty")

	logLink     string
	logPath     string
	logMaxFiles uint
)

var (
	std = logrus.New()
)

// SetupLog 设置日志
func SetupLog() error {
	err := setLogOutput()
	if err != nil {
		return err
	}

	setLogFormatter()

	Infof("Current log path: %s", logPath)
	Infof("Current log link: %s", logLink)
	Infof("Current log max files: %v", logMaxFiles)

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

func WithError(err error) *logrus.Entry {
	return setEntryFileField(std.WithError(err))
}

func WithField(key string, value interface{}) *logrus.Entry {
	return setEntryFileField(std.WithField(key, value))
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return setEntryFileField(std.WithFields(fields))
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

func setLogOutput() error {
	rotateLogs, err := GetRotateLogs()
	if err != nil {
		return err
	}
	std.SetOutput(rotateLogs)
	return nil
}

func setLogFormatter() {
	std.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  FormatTimestamp,
		DisableSorting:   true,
		CallerPrettyfier: callerPrettyfier,
	},
	)
}

func SetLogLevel(lvl string) {
	if lvl == "" {
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

func GetRotateLogs() (*rotatelogs.RotateLogs, error) {

	appName, err := getCurDirName()
	if err != nil {
		return nil, err
	}

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

func getLogEntry() *logrus.Entry {
	frame := getCaller()
	file := shortFile(frame.File)
	file = fmt.Sprintf("%s:%d", file, frame.Line)
	return std.WithField("file", file)
}

func setEntryFileField(entry *logrus.Entry) *logrus.Entry {
	frame := getCaller()
	file := shortFile(frame.File)
	file = fmt.Sprintf("%s:%d", file, frame.Line)
	return entry.WithField("file", file)
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
// There really ought to be to be a better way...
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
