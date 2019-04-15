package log

import (
	"path"
	"runtime"
	"strings"
)

// FilelineCaller returns file and line for caller
func FilelineCaller(skip int) (file string, line int) {
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(skip + i)
		if !ok {
			return "???", 0
		}

		// in go module mode, the import path like this
		//     github.com/fangdingjun/go-log@v4.0.1-incomple/
		if strings.Contains(file, "go-log/") || strings.Contains(file, "go-log@v") {
			continue
		}
		return path.Base(file), line
	}

	return "???", 0
}
