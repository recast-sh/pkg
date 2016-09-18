package log // import "recast.sh/v0/core/log"

import (
	"fmt"
	"sync"
)

func init() {
	logger.formatter = NewStdoutFormatter()
	logger.level = INFO
}

var logger = new(loggerStruct)

type loggerStruct struct {
	sync.Mutex
	level     LogLevel
	formatter Formatter
}

func (p *loggerStruct) internalLog(depth int, inLevel LogLevel, entries ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if inLevel != CRITICAL && p.level < inLevel {
		return
	}
	if logger.formatter != nil {
		logger.formatter.Format(inLevel, depth+1, entries...)
	}
}

func SetLevel(l LogLevel) {
	logger.level = l
}

const calldepth = 2 // TODO is this needed?

// Panic and fatal

func Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	logger.internalLog(calldepth, CRITICAL, s)
	panic(s)
}

func Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	logger.internalLog(calldepth, CRITICAL, s)
	panic(s)
}

// Error Functions

func Errorf(format string, args ...interface{}) {
	logger.internalLog(calldepth, ERROR, fmt.Sprintf(format, args...))
}

func Error(entries ...interface{}) {
	logger.internalLog(calldepth, ERROR, entries...)
}

// Warning Functions

func Warningf(format string, args ...interface{}) {
	logger.internalLog(calldepth, WARNING, fmt.Sprintf(format, args...))
}

func Warning(entries ...interface{}) {
	logger.internalLog(calldepth, WARNING, entries...)
}

// Notice Functions

func Noticef(format string, args ...interface{}) {
	logger.internalLog(calldepth, NOTICE, fmt.Sprintf(format, args...))
}

func Notice(entries ...interface{}) {
	logger.internalLog(calldepth, NOTICE, entries...)
}

func Infof(format string, args ...interface{}) {
	if logger.level < INFO {
		return
	}
	logger.internalLog(calldepth, INFO, fmt.Sprintf(format, args...))
}

func Info(entries ...interface{}) {
	if logger.level < INFO {
		return
	}
	logger.internalLog(calldepth, INFO, entries...)
}

func Debugf(format string, args ...interface{}) {
	if logger.level < DEBUG {
		return
	}
	logger.internalLog(calldepth, DEBUG, fmt.Sprintf(format, args...))
}

func Debug(entries ...interface{}) {
	if logger.level < DEBUG {
		return
	}
	logger.internalLog(calldepth, DEBUG, entries...)
}

func Flush() {
	logger.Lock()
	defer logger.Unlock()
	logger.formatter.Flush()
}
