package globals

import (
	"example/app/cores"
	"fmt"
)

func GlobalJwtConfigInit() *cores.JwtConfig {
	var err error
	var jwtConfig *cores.JwtConfig
	cores.KeepVoid(err, jwtConfig)

	if jwtConfig, err = cores.ViperJwtConfigUnmarshal("jwt_auth"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	jwtSettings := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults["jwtSettings"]))
	jwtSettings.SetValue("algorithm", jwtConfig.Algorithm)
	jwtSettings.SetValue("audience", jwtConfig.Audience)
	jwtSettings.SetValue("issuer", jwtConfig.Issuer)
	jwtSettings.SetValue("secretKey", jwtConfig.SecretKey)
	jwtSettings.SetValue("expiresIn", jwtConfig.ExpiresIn)

	return jwtConfig
}

func GetGlobalJwtConfig() *cores.JwtConfig {
	jwtConfig := cores.NewJwtConfig()
	jwtSettings := cores.Unwrap(cores.Cast[cores.MapAny](ConfigDefaults["jwtSettings"]))

	jwtConfig.Algorithm = jwtSettings.GetValue("algorithm").(string)
	jwtConfig.Audience = jwtSettings.GetValue("audience").(string)
	jwtConfig.Issuer = jwtSettings.GetValue("issuer").(string)
	jwtConfig.SecretKey = jwtSettings.GetValue("secretKey").(string)
	jwtConfig.ExpiresIn = jwtSettings.GetValue("expiresIn").(string)

	return jwtConfig
}
