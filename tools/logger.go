/*
 -- @author: shanweidi
 -- @date: 2023-04-18 10:03 上午
 -- @Desc:
*/
package tools

import (
	"errors"
	"fmt"
	stdLog "log"
	"os"

	"github.com/go-kit/log"
)

const (
	TraceLog = "TRACE"
	DebugLog = "DEBUG"
	InfoLog  = "INFO"
	WarnLog  = "WARN"
	ErrorLog = "ERROR"
)

// LoggingClient defines the interface for logging operations.
type LoggingClient interface {
	// SetLogLevel sets minimum severity log level. If a logging method is called with a lower level of severity than
	// what is set, it will result in no output.
	SetLogLevel(logLevel string) Error
	// LogLevel returns the current log level setting
	LogLevel() string
	// Debug logs a message at the DEBUG severity level
	Debug(msg string, args ...interface{})
	// Error logs a message at the ERROR severity level
	Error(msg string, args ...interface{})
	// Info logs a message at the INFO severity level
	Info(msg string, args ...interface{})
	// Trace logs a message at the TRACE severity level
	Trace(msg string, args ...interface{})
	// Warn logs a message at the WARN severity level
	Warn(msg string, args ...interface{})
	// Debugf logs a formatted message at the DEBUG severity level
	Debugf(msg string, args ...interface{})
	// Errorf logs a formatted message at the ERROR severity level
	Errorf(msg string, args ...interface{})
	// Infof logs a formatted message at the INFO severity level
	Infof(msg string, args ...interface{})
	// Tracef logs a formatted message at the TRACE severity level
	Tracef(msg string, args ...interface{})
	// Warnf logs a formatted message at the WARN severity level
	Warnf(msg string, args ...interface{})
}

type mLogger struct {
	serviceName  string
	logLevel     *string
	rootLogger   log.Logger
	levelLoggers map[string]log.Logger
}

// NewClient creates an instance of LoggingClient
func NewClient(owningServiceName string, logLevel string) LoggingClient {
	if !isValidLogLevel(logLevel) {
		logLevel = InfoLog
	}

	// Set up logging client
	lc := mLogger{
		serviceName: owningServiceName,
		logLevel:    &logLevel,
	}

	lc.rootLogger = log.NewLogfmtLogger(os.Stdout)
	lc.rootLogger = log.WithPrefix(
		lc.rootLogger,
		"ts",
		log.DefaultTimestampUTC,
		"app",
		owningServiceName,
		"source",
		log.Caller(5))

	// Set up the loggers
	lc.levelLoggers = map[string]log.Logger{}

	for _, logLevel := range logLevels() {
		lc.levelLoggers[logLevel] = log.WithPrefix(lc.rootLogger, "level", logLevel)
	}

	return lc
}

// LogLevels returns an array of the possible log levels in order from most to least verbose.
func logLevels() []string {
	return []string{
		TraceLog,
		DebugLog,
		InfoLog,
		WarnLog,
		ErrorLog}
}

func isValidLogLevel(l string) bool {
	for _, name := range logLevels() {
		if name == l {
			return true
		}
	}
	return false
}

func (lc mLogger) log(logLevel string, formatted bool, msg string, args ...interface{}) {
	// Check minimum log level
	for _, name := range logLevels() {
		if name == *lc.logLevel {
			break
		}
		if name == logLevel {
			return
		}
	}

	if args == nil {
		args = []interface{}{"msg", msg}
	} else if formatted {
		args = []interface{}{"msg", fmt.Sprintf(msg, args...)}
	} else {
		if len(args)%2 == 1 {
			// add an empty string to keep k/v pairs correct
			args = append(args, "")
		}
		if len(msg) > 0 {
			args = append(args, "msg", msg)
		}
	}

	err := lc.levelLoggers[logLevel].Log(args...)
	if err != nil {
		stdLog.Fatal(err.Error())
		return
	}

}

func (lc mLogger) SetLogLevel(logLevel string) Error {
	if isValidLogLevel(logLevel) {
		*lc.logLevel = logLevel

		return nil
	}

	return NewSdkError(ContractInvalidErrorCode, ContractInvalidErrorMessage, errors.New("invalid log level:"+logLevel))
}

func (lc mLogger) LogLevel() string {
	if lc.logLevel == nil {
		return ""
	}
	return *lc.logLevel
}

func (lc mLogger) Info(msg string, args ...interface{}) {
	lc.log(InfoLog, false, msg, args...)
}

func (lc mLogger) Trace(msg string, args ...interface{}) {
	lc.log(TraceLog, false, msg, args...)
}

func (lc mLogger) Debug(msg string, args ...interface{}) {
	lc.log(DebugLog, false, msg, args...)
}

func (lc mLogger) Warn(msg string, args ...interface{}) {
	lc.log(WarnLog, false, msg, args...)
}

func (lc mLogger) Error(msg string, args ...interface{}) {
	lc.log(ErrorLog, false, msg, args...)
}

func (lc mLogger) Infof(msg string, args ...interface{}) {
	lc.log(InfoLog, true, msg, args...)
}

func (lc mLogger) Tracef(msg string, args ...interface{}) {
	lc.log(TraceLog, true, msg, args...)
}

func (lc mLogger) Debugf(msg string, args ...interface{}) {
	lc.log(DebugLog, true, msg, args...)
}

func (lc mLogger) Warnf(msg string, args ...interface{}) {
	lc.log(WarnLog, true, msg, args...)
}

func (lc mLogger) Errorf(msg string, args ...interface{}) {
	lc.log(ErrorLog, true, msg, args...)
}
