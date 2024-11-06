package globals

import (
	"nokowebapi/cores"
)

var ConfigDefaults = cores.MapAny{
	"jwtSettings": cores.MapAny{
		"algorithm": "HS256",
		"audience":  "your-audience",
		"issuer":    "your-issuer",
		"secretKey": "your-secret-key",
		"expiresIn": "1h",
	},
	"logger": cores.MapAny{
		"development": true,
		"encoding":     "console",
		"level":       "debug",
	},
}

func Globals() []any {
	return []any{
		JwtConfigGlobals(),
		LoggerConfigGlobals(),
	}
}
