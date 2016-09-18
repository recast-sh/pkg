package log

// LogLevel is the set of all log levels.
type LogLevel int8

const (
	// CRITICAL is the lowest log level; only errors which will end the program will be propagated.
	CRITICAL LogLevel = iota - 1
	// ERROR is for errors that are not fatal but lead to troubling behavior.
	ERROR
	// WARNING is for errors which are not fatal and not errors, but are unusual. Often sourced from misconfigurations.
	WARNING
	// NOTICE is for normal but significant conditions.
	NOTICE
	// INFO is a log level for common, everyday log updates.
	INFO
	// DEBUG is the default hidden level for more verbose updates about internal processes.
	DEBUG
)

// Char returns a single-character representation of the log level.
func (l LogLevel) Char() byte {
	switch l {
	case CRITICAL:
		return byte('C')
	case ERROR:
		return byte('E')
	case WARNING:
		return byte('W')
	case NOTICE:
		return byte('N')
	case INFO:
		return byte('I')
	case DEBUG:
		return byte('D')
	default:
		panic("Unhandled loglevel")
	}
}

// String returns a multi-character representation of the log level.
func (l LogLevel) String() string {
	switch l {
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		panic("Unhandled loglevel")
	}
}
