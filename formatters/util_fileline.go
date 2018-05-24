package formatters

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

		//fmt.Printf("%s:%d\n", file, line)
		if strings.Contains(file, "go-log/") {
			continue
		}
		/*
			// file = pkg/file.go
			n := 0
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					n++
					if n >= 2 {
						file = file[i+1:]
						break
					}
				}
			}
		*/
		return path.Base(file), line
	}

	return "???", 0
}
