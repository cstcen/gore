package gocore

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"runtime"
	"strconv"
)

var (
	ErrAppNameEmpty = errors.New("the app name is empty")
)

// GetRotateLogs 获取文件切割Writer
// appName：项目名称
// opts[0]：最大日志保留数量（个），默认：30
// opts[1]：历史日志路径，默认格式：/xk5/logs/{项目名称}/history/{项目名称}-%Y-%m-%d.log
// opts[2]：LinuxOS中，配置软链接，默认格式：/xk5/logs/{项目名称}/{项目名称}.log
func GetRotateLogs(appName string, opts ...string) (*rotatelogs.RotateLogs, error) {
	if "" == appName {
		return nil, ErrAppNameEmpty
	}
	var (
		path     string
		link     string
		maxFiles uint
	)

	for i, opt := range opts {
		if "" != opt {
			if i == 0 {
				files, err := strconv.ParseUint(opt, 10, 64)
				if err != nil {
					return nil, errors.Wrap(err, "GetRotateLogs")
				}
				maxFiles = uint(files)
			} else if i == 1 {
				path = opt
			} else if i == 2 {
				link = opt
			}
		}
	}

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
