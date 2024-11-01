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
}
