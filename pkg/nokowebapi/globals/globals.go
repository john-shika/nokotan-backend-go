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
		"expires":   "1h",
	},
	"logger": cores.MapAny{
		"development":       true,
		"level":             "debug",
		"encoding":          "console",
		"stackTraceEnabled": true,
	},
}

func Globals() []any {
	return []any{
		JwtConfigGlobals(),
		LoggerConfigGlobals(),
	}
}