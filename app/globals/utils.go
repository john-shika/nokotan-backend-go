package globals

import (
	"example/app/cores"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func GlobalJwtConfigInit() *cores.JwtConfig {
	var err error
	var jwtConfig *cores.JwtConfig
	cores.KeepVoid(err, jwtConfig)

	if jwtConfig, err = cores.ViperJwtConfigUnmarshal("jwt_auth"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	jwtSettings := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults["jwtSettings"]))
	jwtSettings.SetValueByKey("algorithm", jwtConfig.Algorithm)
	jwtSettings.SetValueByKey("audience", jwtConfig.Audience)
	jwtSettings.SetValueByKey("issuer", jwtConfig.Issuer)
	jwtSettings.SetValueByKey("secretKey", jwtConfig.SecretKey)
	jwtSettings.SetValueByKey("expiresIn", jwtConfig.ExpiresIn)

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

func GetGlobalJwtConfig() *cores.JwtConfig {
	jwtConfig := cores.NewJwtConfig()
	jwtSettings := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults["jwtSettings"]))

	jwtConfig.Algorithm = jwtSettings.GetValueByKey("algorithm").(string)
	jwtConfig.Audience = jwtSettings.GetValueByKey("audience").(string)
	jwtConfig.Issuer = jwtSettings.GetValueByKey("issuer").(string)
	jwtConfig.SecretKey = jwtSettings.GetValueByKey("secretKey").(string)
	jwtConfig.ExpiresIn = jwtSettings.GetValueByKey("expiresIn").(string)

	return jwtConfig
}
