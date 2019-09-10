go-log
================
<!--
[![GoDoc](https://godoc.org/github.com/fangdingjun/go-log?status.svg)](https://godoc.org/github.com/fangdingjun/go-log)
[![Build Status](https://travis-ci.org/fangdingjun/go-log.svg?branch=master)](https://travis-ci.org/fangdingjun/go-log)
[![Coverage Status](https://coveralls.io/repos/github/fangdingjun/go-log/badge.svg?branch=master)](https://coveralls.io/github/fangdingjun/go-log?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fangdingjun/go-log)](https://goreportcard.com/report/github.com/fangdingjun/go-log)
[![License](http://img.shields.io/badge/License-Apache_2-red.svg?style=flat)](http://www.apache.org/licenses/LICENSE-2.0)
-->

Logging package similar to log4j for the Golang.

- Support dynamic log level
- Support customized formatter
  - TextFormatter
  - JSONFormatter
- Support multiple rolling file log
  - FixedSizeFileWriter
  - DailyFileWriter
  - AlwaysNewFileWriter

Installation
---------------

```bash
$ go get github.com/fangdingjun/go-log
```

Usage
---------------

```go
package main

import (
	"os"
	"errors"
	log "github.com/fangdingjun/go-log/v5"
)

func main() {
	log.Debugf("app = %s", os.Args[0])
	log.Errorf("error = %v", errors.New("some error"))

	// dynamic set level
	log.Default.Level = log.WARN

	log.Debug("cannot output debug message")
	log.Errorln("can output error message", errors.New("some error"))
}
```

### Output

Default log to console, you can set `Logger.Out` to set a file writer into log.

```go
import (
	log "github.com/fangdingjun/go-log/v5"
)

log.Default.Out = &log.FixedSizeFileWriter{
	Name:	 "/tmp/test.log",
	MaxSize:  10 * 1024 * 1024, // 10m
	MaxCount: 10,
})
```

Three builtin log for use

```go
// Create log file if file size large than fixed size (10m)
// files: /tmp/test.log.0 .. test.log.10
&log.FixedSizeFileWriter{
	Name:	 "/tmp/test.log",
	MaxSize:  10 * 1024 * 1024, // 10m
	MaxCount: 10,
}

// Create log file every day.
// files: /tmp/test.log.20160102
&log.DailyFileWriter{
	Name: "/tmp/test.log",
	MaxCount: 10,
}

// Create log file every process.
// files: /tmp/test.log.20160102_150405
&log.AlwaysNewFileWriter{
	Name: "/tmp/test.log",
	MaxCount: 10,
}

// Output to multiple writes
io.MultiWriter(
	os.Stdout,
	&log.DailyFileWriter{
		Name: "/tmp/test.log",
		MaxCount: 10,
	}
	//...
)
```

### Formatter

```go
import (
	log "github.com/fangdingjun/go-log/v5"
)

log.Default.Formatter = new(log.TextFormatter)
```


### New Logger instance

```go
import (
	log "github.com/fangdingjun/go-log/v5"
)

func main() {
	logger := &log.Logger{
		Level:     log.INFO,
		Formatter: new(log.JSONFormatter),
		Out:       os.Stdout,
	}

	logger.Infof("i = %d", 99)
}
```

## LICENSE

Apache 2.0
