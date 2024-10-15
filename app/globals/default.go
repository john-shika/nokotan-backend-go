package globals

import (
	"example/app/cores"
)

var GlobalDefaults = cores.MapAny{
	"jwtSettings": cores.MapAny{
		"algorithm": "HS256",
		"secretKey": "your-secret-key",
		"audience":  "your-audience",
		"issuer":    "your-issuer",
		"expiresIn": "1h",
	},
}
