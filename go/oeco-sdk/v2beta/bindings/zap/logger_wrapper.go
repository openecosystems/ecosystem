package zaploggerv1

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLoggerWrapper wraps a zap.Logger to implement log-compatible interfaces like Printf for structured logging.
// It provides a convenient way to integrate zap.Logger into libraries expecting standard logging interfaces.
type ZapLoggerWrapper struct {
	logger *zap.Logger
}

// NewZapLoggerWrapper creates and returns a new instance of ZapLoggerWrapper using the provided zap.Logger instance.
func NewZapLoggerWrapper(z *zap.Logger) *ZapLoggerWrapper {
	return &ZapLoggerWrapper{
		logger: z,
	}
}

// Printf logs a formatted message at the INFO level using the zap.Logger instance.
func (z *ZapLoggerWrapper) Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	z.logger.Info(msg)
}

// Debug logs a debug-level message with key-value pairs for structured logging using a sugared logger.
func (z *ZapLoggerWrapper) Debug(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Debugw(msg, keyvals...)
}

// Info logs an informational message with structured key-value pair fields using the SugaredLogger.
func (z *ZapLoggerWrapper) Info(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Infow(msg, keyvals...)
}

// Warn logs a warning-level message with optional structured key-value pairs using the sugared logger.
func (z *ZapLoggerWrapper) Warn(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Warnw(msg, keyvals...)
}

// Error logs an error message with optional key-value pairs using the SugaredLogger for structured logging.
func (z *ZapLoggerWrapper) Error(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Errorw(msg, keyvals...)
}

func (z *ZapLoggerWrapper) Noticef(format string, v ...any) {
	z.logger.Sugar().Infof(format, v...)
}

func (z *ZapLoggerWrapper) Warnf(format string, v ...any) {
	z.logger.Sugar().Warnf(format, v...)
}

func (z *ZapLoggerWrapper) Fatalf(format string, v ...any) {
	z.logger.Sugar().Fatalf(format, v...)
}

func (z *ZapLoggerWrapper) Errorf(format string, v ...any) {
	z.logger.Sugar().Errorf(format, v...)
}

func (z *ZapLoggerWrapper) Debugf(format string, v ...any) {
	z.logger.Sugar().Debugf(format, v...)
}

func (z *ZapLoggerWrapper) Tracef(format string, v ...any) {
	z.logger.Sugar().Debugf(format, v...)
}

// ZapSugaredLoggerWrapper is a wrapper around zap.SugaredLogger to provide custom logging functionalities.
// It includes methods for structured and formatted logging leveraging the underlying zap.SugaredLogger instance.
type ZapSugaredLoggerWrapper struct {
	logger *zap.SugaredLogger
}

// NewZapSugaredLoggerWrapper initializes and returns a new ZapSugaredLoggerWrapper using the provided SugaredLogger instance.
func NewZapSugaredLoggerWrapper(z *zap.SugaredLogger) *ZapSugaredLoggerWrapper {
	return &ZapSugaredLoggerWrapper{
		logger: z,
	}
}

// Printf logs a formatted message at the Info level using the provided format and arguments.
func (z *ZapSugaredLoggerWrapper) Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	z.logger.Info(msg)
}

// Debug logs a debug-level message with key-value pairs for structured logging using a sugared logger.
func (z *ZapSugaredLoggerWrapper) Debug(msg string, keyvals ...interface{}) {
	z.logger.Debugw(msg, keyvals...)
}

// Info logs an informational message with structured key-value pair fields using the SugaredLogger.
func (z *ZapSugaredLoggerWrapper) Info(msg string, keyvals ...interface{}) {
	z.logger.Infow(msg, keyvals...)
}

// Warn logs a warning-level message with optional structured key-value pairs using the sugared logger.
func (z *ZapSugaredLoggerWrapper) Warn(msg string, keyvals ...interface{}) {
	z.logger.Warnw(msg, keyvals...)
}

// Error logs an error message with optional key-value pairs using the SugaredLogger for structured logging.
func (z *ZapSugaredLoggerWrapper) Error(msg string, keyvals ...interface{}) {
	z.logger.Errorw(msg, keyvals...)
}

func (z *ZapSugaredLoggerWrapper) Noticef(format string, v ...any) {
	z.logger.Infof(format, v...)
}

func (z *ZapSugaredLoggerWrapper) Warnf(format string, v ...any) {
	z.logger.Warnf(format, v...)
}

func (z *ZapSugaredLoggerWrapper) Fatalf(format string, v ...any) {
	z.logger.Fatalf(format, v...)
}

func (z *ZapSugaredLoggerWrapper) Errorf(format string, v ...any) {
	z.logger.Errorf(format, v...)
}

func (z *ZapSugaredLoggerWrapper) Debugf(format string, v ...any) {
	z.logger.Debugf(format, v...)
}

func (z *ZapSugaredLoggerWrapper) Tracef(format string, v ...any) {
	z.logger.Debugf(format, v...)
}
