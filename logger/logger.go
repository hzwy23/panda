package logger


var log *Logger

// Error logs a message at error level.
func Error(v ...interface{}) {
	log.Error(v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	log.Warn(v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	log.Info(v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	log.Debug(v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

func init() {
	log = NewLogger(defaultConfig)
}
