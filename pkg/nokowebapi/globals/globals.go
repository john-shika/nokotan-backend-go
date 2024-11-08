package globals

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"nokowebapi/cores"
	"strings"
)

var defaultConfig = cores.MapAny{
	getNameKeyType[cores.JwtConfig](): cores.MapAny{
		"algorithm": "HS256",
		"audience":  "your-audience",
		"issuer":    "your-issuer",
		"secretKey": "your-super-secret-key-keep-it-mind-dont-tell-anyone",
		"expiresIn": "1h",
	},
	getNameKeyType[cores.LoggerConfig](): cores.MapAny{
		"development":       true,
		"level":             "debug",
		"encoding":          "console",
		"stackTraceEnabled": true,
		"colorable":         true,
	},
	"tasks": cores.ArrayAny{},
}

func getNameKeyType[T any]() string {
	return getNameKey(new(T))
}

func getNameKey(obj any) string {
	var ok bool
	cores.KeepVoid(ok)

	name := cores.ToCamelCase(cores.GetNameType(obj))
	name = cores.Unwrap(strings.CutSuffix(name, "Config"))
	return name
}

type ConfigurationImpl interface {
	GetJwtConfig() *cores.JwtConfig
	GetLoggerConfig() *cores.LoggerConfig
	Keys() []string
	Values() []any
	Get(key string) any
}

type Configuration struct {
	jwtConfig    *cores.JwtConfig
	loggerConfig *cores.LoggerConfig
	locker       cores.LockerImpl
}

func NewConfiguration() ConfigurationImpl {
	return &Configuration{
		jwtConfig:    nil,
		loggerConfig: nil,
		locker:       cores.NewLocker(),
	}
}

func (c *Configuration) GetJwtConfig() *cores.JwtConfig {
	c.locker.Lock(func() {
		// pass as you go
		if c.jwtConfig == nil {
			c.jwtConfig = JwtConfigGlobals()
		}
	})
	return c.jwtConfig
}

func (c *Configuration) GetLoggerConfig() *cores.LoggerConfig {
	c.locker.Lock(func() {
		// pass as you go
		if c.loggerConfig == nil {
			c.loggerConfig = LoggerConfigGlobals()
		}
	})
	return c.loggerConfig
}

func (c *Configuration) Keys() []string {
	return []string{
		cores.GetNameType(c.GetJwtConfig()),
		cores.GetNameType(c.GetLoggerConfig()),
	}
}

func (c *Configuration) Values() []any {
	return []any{
		c.GetJwtConfig(),
		c.GetLoggerConfig(),
	}
}

func (c *Configuration) Get(key string) any {
	keys := c.Keys()
	values := c.Values()

	// out of case
	if len(keys) != len(values) {
		panic("invalid configuration")
	}

	// search value by key
	for i, k := range keys {
		if k == key {
			return values[i]
		}
	}

	// not found
	return nil
}

var globals ConfigurationImpl
var locker = cores.NewLocker()

func Globals() ConfigurationImpl {
	locker.Lock(func() {
		if globals == nil {
			globals = NewConfiguration()
		}
	})
	return globals
}

func JwtConfigGlobals() *cores.JwtConfig {
	var err error
	var jwtConfig *cores.JwtConfig
	cores.KeepVoid(err, jwtConfig)

	if jwtConfig, err = cores.ViperJwtConfigUnmarshal(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if jwtConfig == nil {
		return GetJwtConfigGlobals()
	}

	keyName := cores.ToCamelCase(cores.GetNameType(jwtConfig))
	config := defaultConfig.GetValueByKey(keyName).(cores.MapAny)

	config.SetValueByKey("algorithm", jwtConfig.Algorithm)
	config.SetValueByKey("audience", jwtConfig.Audience)
	config.SetValueByKey("issuer", jwtConfig.Issuer)
	config.SetValueByKey("secretKey", jwtConfig.SecretKey)
	config.SetValueByKey("expiresIn", jwtConfig.ExpiresIn)

	switch strings.ToUpper(jwtConfig.Algorithm) {
	case "ES256":
		cores.JwtSigningMethod = jwt.SigningMethodES256
	case "ES384":
		cores.JwtSigningMethod = jwt.SigningMethodES384
	case "ES512":
		cores.JwtSigningMethod = jwt.SigningMethodES512
	case "HS256":
		cores.JwtSigningMethod = jwt.SigningMethodHS256
	case "HS384":
		cores.JwtSigningMethod = jwt.SigningMethodHS384
	case "HS512":
		cores.JwtSigningMethod = jwt.SigningMethodHS512
	case "PS256":
		cores.JwtSigningMethod = jwt.SigningMethodPS256
	case "PS384":
		cores.JwtSigningMethod = jwt.SigningMethodPS384
	case "PS512":
		cores.JwtSigningMethod = jwt.SigningMethodPS512
	case "RS256":
		cores.JwtSigningMethod = jwt.SigningMethodRS256
	case "RS384":
		cores.JwtSigningMethod = jwt.SigningMethodRS384
	case "RS512":
		cores.JwtSigningMethod = jwt.SigningMethodRS512
	}

	return jwtConfig
}

func GetJwtConfigGlobals() *cores.JwtConfig {
	jwtConfig := cores.NewJwtConfig()

	keyName := cores.ToCamelCase(cores.GetNameType(jwtConfig))
	config := defaultConfig.GetValueByKey(keyName).(cores.MapAny)

	jwtConfig.Algorithm = config.GetValueByKey("algorithm").(string)
	jwtConfig.Audience = config.GetValueByKey("audience").(string)
	jwtConfig.Issuer = config.GetValueByKey("issuer").(string)
	jwtConfig.SecretKey = config.GetValueByKey("secretKey").(string)
	jwtConfig.ExpiresIn = config.GetValueByKey("expiresIn").(string)

	return jwtConfig
}

func LoggerConfigGlobals() *cores.LoggerConfig {
	var err error
	var loggerConfig *cores.LoggerConfig
	cores.KeepVoid(err, loggerConfig)

	if loggerConfig, err = cores.ViperLoggerConfigUnmarshal(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if loggerConfig == nil {
		return GetLoggerConfigGlobals()
	}

	keyName := cores.ToCamelCase(cores.GetNameType(loggerConfig))
	config := defaultConfig.GetValueByKey(keyName).(cores.MapAny)

	config.SetValueByKey("development", loggerConfig.Development)
	config.SetValueByKey("level", loggerConfig.Level)
	config.SetValueByKey("encoding", loggerConfig.Encoding)
	config.SetValueByKey("stackTraceEnabled", loggerConfig.StackTraceEnabled)
	config.SetValueByKey("colorable", loggerConfig.Colorable)

	return loggerConfig
}

func GetLoggerConfigGlobals() *cores.LoggerConfig {
	loggerConfig := cores.NewLoggerConfig()
	keyName := cores.ToCamelCase(cores.GetNameType(loggerConfig))
	config := defaultConfig.GetValueByKey(keyName).(cores.MapAny)

	loggerConfig.Development = config.GetValueByKey("development").(bool)
	loggerConfig.Level = config.GetValueByKey("level").(string)
	loggerConfig.Encoding = config.GetValueByKey("encoding").(string)
	loggerConfig.StackTraceEnabled = config.GetValueByKey("stackTraceEnabled").(bool)
	loggerConfig.Colorable = config.GetValueByKey("colorable").(bool)

	return loggerConfig
}
