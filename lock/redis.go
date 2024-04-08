package lock

import (
	"context"
	goreRedis "github.com/cstcen/gore/db/redis"
	"github.com/cstcen/gore/gonfig"
	"github.com/go-redis/redis/v8"
	"log/slog"
	"strings"
	"time"
)

func NewDefaultRedis(c context.Context, keys ...string) *Redis {
	return NewRedis(c, goreRedis.Instance(), keys...)
}

func NewRedis(c context.Context, cmdable redis.Cmdable, keys ...string) *Redis {
	env := gonfig.Instance().GetString("env")
	name := gonfig.Instance().GetString("name")
	k := []string{env, name, "lock"}
	for _, key := range keys {
		k = append(k, key)
	}
	return &Redis{c: c, cmdable: cmdable, key: strings.Join(k, ":")}
}

type Redis struct {
	c       context.Context
	key     string
	cmdable redis.Cmdable
}

func (r *Redis) Lock() {
	if err := r.cmdable.SetNX(r.c, r.key, 1, 10*time.Minute).Err(); err != nil {
		slog.WarnContext(r.c, "lock failure", "key", r.key, "err", err.Error())
		return
	}
	slog.InfoContext(r.c, "lock successful", "key", r.key)
}

func (r *Redis) Unlock() {
	if err := r.cmdable.Del(r.c, r.key).Err(); err != nil {
		slog.WarnContext(r.c, "unlock failure", "key", r.key, "err", err.Error())
		return
	}
	slog.InfoContext(r.c, "unlock successful", "key", r.key)
}

func (r *Redis) TryLock() bool {
	result, err := r.cmdable.SetNX(r.c, r.key, 1, 10*time.Minute).Result()
	if err != nil {
		slog.WarnContext(r.c, "try lock failure", "key", r.key, "err", err.Error())
		return false
	}
	if !result {
		slog.WarnContext(r.c, "try lock failure", "key", r.key)
		return false
	}
	slog.InfoContext(r.c, "try lock successful", "key", r.key)
	return true
}

func (r *Redis) Key() string {
	return r.key
}
