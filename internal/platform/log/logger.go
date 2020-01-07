// Package log configures a new logger for an application.
package log

import (
	"errors"
)

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	// InstanceZapLogger indicates the logger type is zap
	InstanceZapLogger int = iota
	// InstanceLogrusLogger indicates the logger type is logrus
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

//NewLogger returns an instance of logger
func NewLogger(config Config, loggerInstance int) (Logger, error) {
	switch loggerInstance {
	case InstanceZapLogger:
		logger := newZapLogger(config)
		return logger, nil

	case InstanceLogrusLogger:
		logger := newLogrusLogger(config)
		return logger, nil

	default:
		return nil, errInvalidLoggerInstance
	}
}
