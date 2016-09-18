package log

// SetFormatter sets the formatting function for all logs.
func SetFormatter(f Formatter) {
	logger.Lock()
	defer logger.Unlock()
	logger.formatter = f
}

type Formatter interface {
	Format(level LogLevel, depth int, entries ...interface{})
	Flush()
	// TODO add Sync?
}
