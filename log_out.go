package log

import (
	"context"
	"io"
	stdlog "log"
	"log/slog"
	"runtime"
	"time"
)

type logOut struct {
}

func (o *logOut) Write(buf []byte) (int, error) {
	h := Default.logger.Handler()
	if !h.Enabled(context.Background(), slog.LevelInfo) {
		return len(buf), nil
	}
	pc := [1]uintptr{}
	runtime.Callers(4, pc[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, string(buf[:len(buf)-1]), pc[0])
	h.Handle(context.Background(), r)
	return len(buf), nil
}

var _ io.Writer = &logOut{}

func WrapStdLog() {
	slog.SetDefault(Default.logger)
	stdlog.SetFlags(0)
	stdlog.SetOutput(&logOut{})
}
