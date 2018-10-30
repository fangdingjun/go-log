package log

// Won't compile if LogInterface can't be realized by a log.Logger
var (
	_ StdLog    = Default
	_ Interface = Default
)

// StdLog is interface for builtin log
type StdLog interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// Interface is interface for this logger
type Interface interface {
	Debug(...interface{})
	Info(...interface{})
	Print(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})

	Debugln(...interface{})
	Infoln(...interface{})
	Println(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Panicln(...interface{})
	Fatalln(...interface{})

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Printf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}
