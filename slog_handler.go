package log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type slogHandler struct {
	opts  slog.HandlerOptions
	w     io.Writer
	attrs []slog.Attr
	group []string
	mu    sync.Mutex
}

func NewSlogHandler(w io.Writer, opts *slog.HandlerOptions) *slogHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				return a
			},
		}
	}
	return &slogHandler{w: w, opts: *opts}
}

// Enabled implements slog.Handler.
func (s *slogHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= s.opts.Level.Level()
}

// Handle implements slog.Handler.
func (s *slogHandler) Handle(ctx context.Context, record slog.Record) error {
	buf := new(bytes.Buffer)
	if !record.Time.IsZero() {
		fmt.Fprintf(buf, "%s", record.Time.Format("2006-01-02 15:04:05.000"))
	} else {
		t := time.Now()
		fmt.Fprintf(buf, "%s", t.Format("2006-01-02 15:04:05.000"))
	}
	if s.opts.AddSource {
		fs := runtime.CallersFrames([]uintptr{record.PC})
		f, _ := fs.Next()
		fmt.Fprintf(buf, " [%s:%d]", filepath.Base(f.File), f.Line)
	}
	fmt.Fprintf(buf, " %s:", record.Level.String())
	if len(s.group) > 0 {
		fmt.Fprintf(buf, " %s", strings.Join(s.group, " "))
	}
	fmt.Fprintf(buf, " %s", record.Message)
	for _, a := range s.attrs {
		fmt.Fprintf(buf, " %s=%s", a.Key, a.Value.String())
	}
	if record.NumAttrs() > 0 {
		record.Attrs(func(a slog.Attr) bool {
			fmt.Fprintf(buf, " %s=%s", a.Key, a.Value.String())
			return true
		})
	}
	fmt.Fprintf(buf, "\n")

	s.mu.Lock()
	defer s.mu.Unlock()

	s.w.Write(buf.Bytes())

	return nil
}

// WithAttrs implements slog.Handler.
func (s *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	a := s.attrs
	a = append(a, attrs...)
	return &slogHandler{
		opts:  s.opts,
		w:     s.w,
		attrs: a,
		group: s.group,
	}
}

// WithGroup implements slog.Handler.
func (s *slogHandler) WithGroup(name string) slog.Handler {
	a := s.group
	a = append(a, name)
	return &slogHandler{
		opts:  s.opts,
		w:     s.w,
		attrs: s.attrs,
		group: a,
	}
}

var _ slog.Handler = &slogHandler{}
