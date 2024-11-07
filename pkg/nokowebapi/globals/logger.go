package globals

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"nokowebapi/cores"
)

type LoggerConfig struct {
	Development       bool   `mapstructure:"development" json:"development"`
	Encoding          string `mapstructure:"encoding" json:"encoding"`
	Level             string `mapstructure:"level" json:"level"`
	StackTraceEnabled bool   `mapstructure:"stack_trace_enabled" json:"stackTraceEnabled"`
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{}
}

func (LoggerConfig) GetName() string {
	return "Logger"
}

func GetLoggerConfigLevel(loggerConfig *LoggerConfig) zapcore.Level {
	switch cores.ToCamelCase(loggerConfig.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func customEncoderConfig(encoderConfig zapcore.EncoderConfig) zapcore.EncoderConfig {
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}

func GetLoggerConfigEncoder(loggerConfig *LoggerConfig, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	switch cores.ToSnakeCase(loggerConfig.Encoding) {
	case "console":
		return zapcore.NewConsoleEncoder(customEncoderConfig(encoderConfig))
	case "json":
		return zapcore.NewJSONEncoder(customEncoderConfig(encoderConfig))
	default:
		return zapcore.NewConsoleEncoder(customEncoderConfig(encoderConfig))
	}
}

func ViperLoggerConfigUnmarshal() (*LoggerConfig, error) {
	var err error
	cores.KeepVoid(err)

	loggerConfig := NewLoggerConfig()
	keyName := cores.ToSnakeCase(cores.GetName(loggerConfig))
	if err = viper.UnmarshalKey(keyName, loggerConfig); err != nil {
		return nil, err
	}

	return loggerConfig, nil
}

func LoggerConfigGlobals() *LoggerConfig {
	var err error
	var loggerConfig *LoggerConfig
	cores.KeepVoid(err, loggerConfig)

	if loggerConfig, err = ViperLoggerConfigUnmarshal(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if loggerConfig == nil {
		return GetLoggerConfigGlobals()
	}

	keyName := cores.ToCamelCase(cores.GetName(loggerConfig))
	config := Defaults.GetValueByKey(keyName).(cores.MapAny)

	config.SetValueByKey("development", loggerConfig.Development)
	config.SetValueByKey("level", loggerConfig.Level)
	config.SetValueByKey("encoding", loggerConfig.Encoding)
	config.SetValueByKey("stackTraceEnabled", loggerConfig.StackTraceEnabled)

	return loggerConfig
}

func GetLoggerConfigGlobals() *LoggerConfig {
	loggerConfig := NewLoggerConfig()
	keyName := cores.ToCamelCase(cores.GetName(loggerConfig))
	config := Defaults.GetValueByKey(keyName).(cores.MapAny)

	loggerConfig.Development = config.GetValueByKey("development").(bool)
	loggerConfig.Level = config.GetValueByKey("level").(string)
	loggerConfig.Encoding = config.GetValueByKey("encoding").(string)
	loggerConfig.StackTraceEnabled = config.GetValueByKey("stackTraceEnabled").(bool)

	return loggerConfig
}
