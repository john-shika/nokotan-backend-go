package cores

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"time"
)

var ErrJwtTokenNotFound = errors.New("exception: JWT  token not found")
var ErrJwtClaimsInvalid = errors.New("exception: Invalid JWT claims")
var ErrJwtIdentityNotFound = errors.New("exception: JWT  identity not found")
var ErrJwtIssuedAtNotFound = errors.New("exception: JWT  issued at not found")
var ErrJwtIssuerNotFound = errors.New("exception: JWT  issuer not found")
var ErrJwtSubjectNotFound = errors.New("exception: JWT  subject not found")
var ErrJwtExpiresNotFound = errors.New("exception: JWT  expires not found")
var ErrJwtSessionIdNotFound = errors.New("exception: JWT  session id not found")
var ErrJwtUserInvalid = errors.New("exception: JWT  user not found")
var ErrJwtEmailInvalid = errors.New("exception: JWT  email not found")
var ErrJwtSecretKeyNotFound = errors.New("exception: JWT  secret key not found")

type JwtConfig struct {
	Algorithm string `mapstructure:"algorithm" json:"algorithm"`
	SecretKey string `mapstructure:"secret_key" json:"secretKey"`
	Audience  string `mapstructure:"audience" json:"audience"`
	Issuer    string `mapstructure:"issuer" json:"issuer"`
	ExpiresIn string `mapstructure:"expires_in" json:"expiresIn"`
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{}
}

type JwtClaimNameTyped string

const (
	JwtClaimNameIdentity  JwtClaimNameTyped = "jti"
	JwtClaimNameSubject                     = "sub"
	JwtClaimNameSessionId                   = "sid"
	JwtClaimNameIssuedAt                    = "iat"
	JwtClaimNameIssuer                      = "iss"
	JwtClaimNameExpires                     = "exp"
	JwtClaimNameUser                        = "user"
	JwtClaimNameEmail                       = "email"
	JwtClaimNameRole                        = "role"
)

type Claims struct {
	jwt.MapClaims
}

func NewClaims(claims jwt.MapClaims) *Claims {
	return &Claims{
		claims,
	}
}

func GetClaimsFromJwtToken(token *jwt.Token) (*Claims, error) {
	var ok bool
	var claims jwt.MapClaims
	KeepVoid(ok, claims)

	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, ErrJwtClaimsInvalid
	}

	return NewClaims(claims), nil
}

func (claims *Claims) Get(key string) (any, bool) {
	var ok bool
	var value any
	KeepVoid(ok, value)

	value, ok = claims.MapClaims[key]
	return value, ok
}

func (claims *Claims) Set(key string, value any) bool {
	var ok bool
	var temp any
	KeepVoid(ok, temp)

	if temp, ok = claims.MapClaims[key]; !ok {
		return false
	}

	claims.MapClaims[key] = value
	return true
}

func NewNumericDateFromSeconds(f float64) *jwt.NumericDate {
	round, frac := math.Modf(f)
	return jwt.NewNumericDate(time.Unix(int64(round), int64(frac*1e9)))
}

func (claims *Claims) ParseNumericDate(key string) (*jwt.NumericDate, error) {
	var ok bool
	var err error
	var value any
	KeepVoid(ok, err, value)

	if value, ok = claims.Get(key); !ok {
		return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	switch exp := value.(type) {
	case float64:
		if exp == 0 {
			return nil, nil
		}

		return NewNumericDateFromSeconds(exp), nil
	case json.Number:
		if value, err = exp.Float64(); err != nil {
			return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
		}

		return NewNumericDateFromSeconds(value.(float64)), nil
	}

	return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
}

func (claims *Claims) ParseString(key string) (string, error) {
	var ok bool
	var value any
	var temp string
	KeepVoid(ok, value, temp)

	if value, ok = claims.Get(key); !ok {
		return EmptyString, ErrDataTypeInvalid
	}

	if temp, ok = value.(string); !ok {
		return EmptyString, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	return temp, nil
}

func (claims *Claims) GetIdentity() (string, error) {
	return claims.ParseString("jti")
}

func (claims *Claims) SetIdentity(identity string) bool {
	return claims.Set("jti", identity)
}

func (claims *Claims) GetSubject() (string, error) {
	return claims.ParseString("sub")
}

func (claims *Claims) SetSubject(subject string) bool {
	return claims.Set("sub", subject)
}

func (claims *Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return claims.ParseNumericDate("iat")
}

func (claims *Claims) SetIssuedAt(date time.Time) bool {
	return claims.Set("iat", &jwt.NumericDate{Time: date})
}

func (claims *Claims) GetIssuer() (string, error) {
	return claims.ParseString("iss")
}

func (claims *Claims) SetIssuer(issuer string) bool {
	return claims.Set("iss", issuer)
}

func (claims *Claims) GetExpires() (*jwt.NumericDate, error) {
	return claims.ParseNumericDate("exp")
}

func (claims *Claims) SetExpires(date time.Time) bool {
	return claims.Set("exp", &jwt.NumericDate{Time: date})
}

func (claims *Claims) GetSessionId() (string, error) {
	return claims.ParseString("sid")
}

func (claims *Claims) SetSessionId(sessionId string) bool {
	return claims.Set("sid", sessionId)
}

func (claims *Claims) GetUser() (string, error) {
	return claims.ParseString("user")
}

func (claims *Claims) SetUser(user string) bool {
	return claims.Set("user", user)
}

func (claims *Claims) GetRole() (string, error) {
	return claims.ParseString("role")
}

func (claims *Claims) SetRole(role string) bool {
	return claims.Set("role", role)
}

func (claims *Claims) GetEmail() (string, error) {
	return claims.ParseString("email")
}

func (claims *Claims) SetEmail(email string) bool {
	return claims.Set("email", email)
}
