package globals

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"nokowebapi/cores"
	"strings"
)

func JwtConfigGlobals() *cores.JwtConfig {
	var err error
	var jwtConfig *cores.JwtConfig
	cores.KeepVoid(err, jwtConfig)

	if jwtConfig, err = cores.ViperJwtConfigUnmarshal(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	keyName := cores.ToCamelCase(cores.GetNameReflection(jwtConfig))
	config := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults.GetValueByKey(keyName)))

	config.SetValueByKey("algorithm", jwtConfig.Algorithm)
	config.SetValueByKey("audience", jwtConfig.Audience)
	config.SetValueByKey("issuer", jwtConfig.Issuer)
	config.SetValueByKey("secretKey", jwtConfig.SecretKey)
	config.SetValueByKey("expires", jwtConfig.Expires)

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

	keyName := cores.ToCamelCase(cores.GetNameReflection(jwtConfig))
	config := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults.GetValueByKey(keyName)))

	jwtConfig.Algorithm = config.GetValueByKey("algorithm").(string)
	jwtConfig.Audience = config.GetValueByKey("audience").(string)
	jwtConfig.Issuer = config.GetValueByKey("issuer").(string)
	jwtConfig.SecretKey = config.GetValueByKey("secretKey").(string)
	jwtConfig.Expires = config.GetValueByKey("expires").(string)

	return jwtConfig
}