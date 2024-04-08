package log

import (
	"context"
	"github.com/cstcen/gore/gonfig"
	"github.com/natefinch/lumberjack"
	"io"
	"log/slog"
	"os"
	"time"
)

var (
	defaultRequestIdKey = "X-Request-ID"
)

func MustRequestID(c context.Context) string {
	id, ok := c.Value(defaultRequestIdKey).(string)
	if !ok {
		return ""
	}
	return id
}

type Config struct {
	*lumberjack.Logger
	Level slog.Level
}

func SetupDefault() error {
	cfg := Config{Logger: &lumberjack.Logger{}}
	if gonfig.Instance().GetBool("log") {
		if err := gonfig.Instance().UnmarshalKey("gore.logger", cfg.Logger); err != nil {
			return err
		}
	}
	if err := cfg.Level.UnmarshalText([]byte(gonfig.Instance().GetString("gore.logger.level"))); err != nil {
		cfg.Level = slog.LevelDebug
	}
	return Setup(&cfg)
}

func Setup(cfg *Config) error {

	var logWriter io.Writer
	if cfg.Logger != nil {
		logWriter = io.MultiWriter(os.Stdout, cfg)
	} else {
		logWriter = os.Stdout
	}
	slog.SetDefault(slog.New(NewSlogHandler(logWriter, &slog.HandlerOptions{Level: cfg.Level, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Value = slog.StringValue(a.Value.Time().Format(time.DateTime))
			return a
		}
		return a
	}})))

	return nil
}

type SlogHandler struct {
	*slog.TextHandler
}

func NewSlogHandler(w io.Writer, opts *slog.HandlerOptions) *SlogHandler {
	return &SlogHandler{TextHandler: slog.NewTextHandler(w, opts)}
}

func (h *SlogHandler) Handle(c context.Context, r slog.Record) error {
	id := MustRequestID(c)
	r.Add(defaultRequestIdKey, id)
	return h.TextHandler.Handle(c, r)
}
