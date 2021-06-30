package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerImpl struct {
	logger *zap.SugaredLogger
}

// NewLogger get the new logger based on environment.
func NewLogger(environment string) Logger {
	var conf zap.Config

	if environment == "release" {
		conf = newProductionConfig()
	} else {
		conf = newDevelopmentConfig()
	}

	conf.DisableStacktrace = true
	log, err := conf.Build()
	if err != nil {
		panic(err)
	}

	return &loggerImpl{
		logger: log.WithOptions(zap.AddCallerSkip(1)).Sugar(),
	}
}

// newProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// newProductionConfig is a reasonable production logging configuration.
// Logging is enabled at InfoLevel and above.
//
// It uses a JSON encoder, writes to standard error, and enables sampling.
// Stacktraces are automatically included on logs of ErrorLevel and above.
func newProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// newDevelopmentEncoderConfig returns an opinionated EncoderConfig for
// development environments.
func newDevelopmentEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// newDevelopmentConfig is a reasonable development logging configuration.
// Logging is enabled at DebugLevel and above.
//
// It enables development mode (which makes DPanicLevel logs panic), uses a
// console encoder, writes to standard error, and disables sampling.
// Stacktraces are automatically included on logs of WarnLevel and above.
func newDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    newDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Debug uses fmt.Sprint to construct and log a message
func (l *loggerImpl) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *loggerImpl) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With
func (l *loggerImpl) Debugw(msg string, keysValues ...interface{}) {
	l.logger.Debugw(msg, keysValues...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *loggerImpl) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message
func (l *loggerImpl) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *loggerImpl) Infow(msg string, keysValues ...interface{}) {
	l.logger.Infow(msg, keysValues...)
}

// Warn uses fmt.Sprint to construct and log a message
func (l *loggerImpl) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnf uses fmt.Sprintf to log a templated message
func (l *loggerImpl) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *loggerImpl) Warnw(msg string, keysValues ...interface{}) {
	l.logger.Warnw(msg, keysValues...)
}

// Error uses fmt.Sprint to construct and log a message
func (l *loggerImpl) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message
func (l *loggerImpl) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *loggerImpl) Errorw(msg string, keysValues ...interface{}) {
	l.logger.Errorw(msg, keysValues...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit
func (l *loggerImpl) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit
func (l *loggerImpl) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With
func (l *loggerImpl) Fatalw(msg string, keysValues ...interface{}) {
	l.logger.Fatalw(msg, keysValues...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics
func (l *loggerImpl) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *loggerImpl) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With
func (l *loggerImpl) Panicw(msg string, keysValues ...interface{}) {
	l.logger.Panicw(msg, keysValues...)
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *loggerImpl) With(args ...interface{}) Logger {
	return &loggerImpl{
		logger: l.logger.With(args...),
	}
}
