/*
Package log is a log/slog wrapper
*/
package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"
)

// Default is a default Logger instance
var Default *Logger

type Logger struct {
	logger *slog.Logger
}

func New() *Logger {
	h := NewSlogHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug})
	l := slog.New(h)
	return &Logger{
		logger: l,
	}
}

func (l *Logger) log(lvl slog.Level, fmts string, args ...interface{}) {
	h := l.logger.Handler()
	if !h.Enabled(context.Background(), lvl) {
		return
	}
	a := fmt.Sprintf(fmts, args...)
	pc := [1]uintptr{}
	runtime.Callers(3, pc[:])
	r := slog.NewRecord(time.Now(), lvl, a, pc[0])
	h.Handle(context.Background(), r)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.log(slog.LevelDebug, msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.log(slog.LevelInfo, msg, args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.log(slog.LevelWarn, msg, args...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.log(slog.LevelError, msg, args...)
}

func (l *Logger) SetOutput(w io.Writer) {
	h := l.logger.Handler()
	h1 := h.(*slogHandler)
	h1.mu.Lock()
	defer h1.mu.Unlock()
	h1.w = w
}

func (l *Logger) SetLevel(level slog.Level) {
	h := l.logger.Handler()
	h1 := h.(*slogHandler)
	h1.mu.Lock()
	defer h1.mu.Unlock()
	h1.opts.Level = level

}

func (l *Logger) SetOutputFile(fn string) {
	out := &FixedSizeFileWriter{
		Name:     fn,
		MaxSize:  100 * 1024 * 1024,
		MaxCount: 10,
	}
	SetOutput(out)
}

// Debugf outputs message, Arguments are handled by fmt.Sprintf
func Debugf(msg string, args ...interface{}) {
	Default.log(slog.LevelDebug, msg, args...)
}

// Infof outputs message, Arguments are handled by fmt.Sprintf
func Infof(msg string, args ...interface{}) {
	Default.log(slog.LevelInfo, msg, args...)
}

// Warnf outputs message, Arguments are handled by fmt.Sprintf
func Warnf(msg string, args ...interface{}) {
	Default.log(slog.LevelWarn, msg, args...)
}

// Errorf outputs message, Arguments are handled by fmt.Sprintf
func Errorf(msg string, args ...interface{}) {
	Default.log(slog.LevelError, msg, args...)
}

func Errorln(msg ...interface{}) {
	fmts := []string{}
	for i := 0; i < len(msg); i++ {
		fmts = append(fmts, "%+v")
	}
	Default.log(slog.LevelError, strings.Join(fmts, " "), msg)
}

func Fatal(msg ...interface{}) {
	fmts := []string{}
	for i := 0; i < len(msg); i++ {
		fmts = append(fmts, "%+v")
	}
	Default.log(slog.LevelError, strings.Join(fmts, " "), msg)
	os.Exit(1)
}

func SetOutput(w io.Writer) {
	Default.SetOutput(w)
}

func SetLevel(level slog.Level) {
	Default.SetLevel(level)
}

func SetOutputFile(fn string) {
	Default.SetOutputFile(fn)
}

func init() {
	Default = New()
	//setLogDefault("")
}
