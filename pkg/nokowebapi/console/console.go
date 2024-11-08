package console

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"nokowebapi/cores"
	"nokowebapi/globals"
	"nokowebapi/xterm"
)

var WriterSyncer zapcore.WriteSyncer

var Logger *zap.Logger
var Locker cores.LockerImpl

func GetLocker() cores.LockerImpl {
	if Locker != nil {
		return Locker
	}
	Locker = cores.NewLocker()
	return Locker
}

func updateWriterSyncer(stdout io.Writer) zapcore.WriteSyncer {
	var ok bool
	cores.KeepVoid(ok)

	if stdout, ok = stdout.(zapcore.WriteSyncer); ok {
		WriterSyncer = zapcore.Lock(stdout.(zapcore.WriteSyncer))
	} else {
		WriterSyncer = zapcore.AddSync(stdout)
	}

	return WriterSyncer
}

func NewWriterSyncer(stdout io.Writer) zapcore.WriteSyncer {
	if WriterSyncer != nil {
		return WriterSyncer
	}

	return updateWriterSyncer(stdout)
}

func makeLogger() *zap.Logger {
	var logger *zap.Logger
	cores.KeepVoid(logger)

	loggerConfig := globals.Globals().GetLoggerConfig()
	writerSyncer := NewWriterSyncer(xterm.Stdout)
	level := loggerConfig.GetLevel()

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(4),
	}

	if loggerConfig.StackTraceEnabled {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	options = append(options, zap.IncreaseLevel(level))

	if loggerConfig.Development {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		logger = zap.New(core, options...)
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		logger = zap.New(core, options...)
	}

	logger = logger.Named("[NokoWebApi]")
	zap.ReplaceGlobals(logger)
	return logger
}

func NewLogger() *zap.Logger {
	if Logger != nil {
		return Logger
	}
	Logger = makeLogger()
	return Logger
}

func Dir(obj any, fields ...zap.Field) {
	locker := GetLocker()
	locker.Lock(func() {
		logger := NewLogger()
		data := "\n" + cores.ShikaYamlEncode(obj)
		logger.Info(data, fields...)
	})
}

func Log(msg string, fields ...zap.Field) {
	locker := GetLocker()
	locker.Lock(func() {
		logger := NewLogger()
		logger.Info(msg, fields...)
	})
}

func Logf(format string, args ...any) {
	locker := GetLocker()
	locker.Lock(func() {
		logger := NewLogger()
		logger.Info(fmt.Sprintf(format, args...))
	})
}

func Warn(msg string, fields ...zap.Field) {
	locker := GetLocker()
	locker.Lock(func() {
		logger := NewLogger()
		logger.Warn(msg, fields...)
	})
}

func Error(msg string, fields ...zap.Field) {
	locker := GetLocker()
	locker.Lock(func() {
		defer updateWriterSyncer(xterm.Stdout)
		updateWriterSyncer(xterm.Stderr)

		logger := NewLogger()
		logger.Error(msg, fields...)
	})
}

func Fatal(msg string, fields ...zap.Field) {
	locker := GetLocker()
	locker.Lock(func() {
		defer updateWriterSyncer(xterm.Stdout)
		updateWriterSyncer(xterm.Stderr)

		logger := NewLogger()
		logger.Fatal(msg, fields...)
	})
}
