package log

import (
	"log/slog"
	"os"
	"testing"
)

func TestSlogHandler(t *testing.T) {
	h := NewSlogHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true})
	logger := slog.New(h)
	logger.Debug("this is debug")
	logger.Error("error log")
	logger.Info("hello", "a", 1)
	log2 := logger.With("a", "b")
	log2.Info("bbbb")

	logger.WithGroup("cccc").Info("ccccc")

	log3 := slog.NewLogLogger(h, slog.LevelInfo)
	log3.Printf("aaaa from log")
	log3.Println("bbb from log")

	slog.SetDefault(logger)

	slog.Info("log from default logger")

	slog.Info("test log", "a", "1", "b", "2")
}
