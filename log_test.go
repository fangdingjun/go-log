package log

import (
	stdlog "log"
	"log/slog"
	"testing"
)

func TestLogOut(t *testing.T) {
	Debugf("this is debug")
	Infof("this is info")
	Warnf("this is warn")
	Errorf("this is error")
}

func TestSyslog(t *testing.T) {
	Debugf("this is debug")
	Infof("this is info")
	Warnf("this is warn")
	Errorf("this is error")
}

func TestStdlog(t *testing.T) {
	WrapStdLog()
	stdlog.Printf("aaa stdlog")
	a := stdlog.Default()
	a.Printf("bb")

	slog.Info("from slog")
	slog.Debug("slog debug")
	slog.Info("slog info")
	slog.Error("slog error")

	Infof("infof wrapper")
	Errorf("errorf wrapper")

	SetLevel(slog.LevelError)
	slog.Info("info")
	slog.Error("error")
}

func TestLogger(t *testing.T) {
	log1 := New()
	log1.Infof("aaa")
}
