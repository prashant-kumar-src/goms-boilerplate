package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type zLogger struct {
	lClient zerolog.Logger
}

// Debug implements ILogger.
func (l *zLogger) Debug(msg string) {
	l.lClient.Debug().Msg(msg)
}

// Debugf implements ILogger.
func (l *zLogger) Debugf(msg string, args ...interface{}) {
	l.lClient.Debug().Msgf(msg, args...)
}

// Error implements ILogger.
func (l *zLogger) Error(msg string) {
	l.lClient.Error().Msg(msg)

}

// Errorf implements ILogger.
func (l *zLogger) Errorf(msg string, args ...interface{}) {
	l.lClient.Error().Msgf(msg, args...)
}

// Fatal implements ILogger.
func (l *zLogger) Fatal(msg string) {
	l.lClient.Fatal().Msg(msg)
}

// Fatalf implements ILogger.
func (l *zLogger) Fatalf(msg string, args ...interface{}) {
	l.lClient.Fatal().Msgf(msg, args...)
}

// Info implements ILogger.
func (l *zLogger) Info(msg string) {
	l.lClient.Info().Msg(msg)
}

// Infof implements ILogger.
func (l *zLogger) Infof(msg string, args ...interface{}) {
	l.lClient.Info().Msgf(msg, args...)
}

// Panic implements ILogger.
func (l *zLogger) Panic(msg string) {
	l.lClient.Panic().Msg(msg)
}

// Panicf implements ILogger.
func (l *zLogger) Panicf(msg string, args ...interface{}) {
	l.lClient.Panic().Msgf(msg, args...)
}

// Trace implements ILogger.
func (l *zLogger) Trace(msg string) {
	l.lClient.Trace().Msg(msg)
}

// Tracef implements ILogger.
func (l *zLogger) Tracef(msg string, args ...interface{}) {
	l.lClient.Trace().Msgf(msg, args...)
}

// Warn implements ILogger.
func (l *zLogger) Warn(msg string) {
	l.lClient.Warn().Msg(msg)
}

// Warnf implements ILogger.
func (l *zLogger) Warnf(msg string, args ...interface{}) {
	l.lClient.Warn().Msgf(msg, args...)
}

func newZeroLogger() ILogger {
	return &zLogger{
		lClient: zerolog.New(os.Stdout),
	}
}
