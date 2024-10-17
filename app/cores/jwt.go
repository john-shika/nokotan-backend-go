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

var JwtSigningMethod = jwt.SigningMethodHS256

// TODO: not implemented yet

type JwtConfig struct {
	Algorithm string `mapstructure:"algorithm" json:"algorithm"`
	SecretKey string `mapstructure:"secret_key" json:"secretKey"`
	Audience  string `mapstructure:"audience" json:"audience"`
	Issuer    string `mapstructure:"issuer" json:"issuer"`
	ExpiresIn string `mapstructure:"expires_in" json:"expiresIn"`
}

func NewJwtConfig() *JwtConfig {
	return new(JwtConfig)
}

// TODO: not implemented yet

type JwtClaimNameTyped string

const (
	JwtClaimNameIdentity  JwtClaimNameTyped = "jti"
	JwtClaimNameSubject   JwtClaimNameTyped = "sub"
	JwtClaimNameSessionId JwtClaimNameTyped = "sid"
	JwtClaimNameIssuedAt  JwtClaimNameTyped = "iat"
	JwtClaimNameIssuer    JwtClaimNameTyped = "iss"
	JwtClaimNameAudience  JwtClaimNameTyped = "aud"
	JwtClaimNameExpiresAt JwtClaimNameTyped = "exp"
	JwtClaimNameUser      JwtClaimNameTyped = "user"
	JwtClaimNameEmail     JwtClaimNameTyped = "email"
	JwtClaimNameRole      JwtClaimNameTyped = "role"
)

type JwtTokenImpl interface {
	SignedString(key interface{}) (string, error)
	SigningString() (string, error)
	EncodeSegment(seg []byte) string
}

type TimeOrNumericDateImpl interface {
	*jwt.NumericDate | time.Time
}

func GetTimeFromTimeOrNumericDateType[V TimeOrNumericDateImpl](date V) time.Time {
	var ok bool
	var t time.Time
	var d *jwt.NumericDate
	KeepVoid(ok, t, d)

	if date == nil {
		return time.Time{}
	}

	if t, ok = Cast[time.Time](date); !ok {
		if d, ok = Cast[*jwt.NumericDate](date); ok {
			return d.Time
		}
		return time.Time{}
	}

	return t
}

func GetNumericDateFromTimeOrNumericDateType[V TimeOrNumericDateImpl](date V) *jwt.NumericDate {
	var ok bool
	var t *jwt.NumericDate
	var d time.Time
	KeepVoid(ok, t, d)

	if date == nil {
		return nil
	}

	if t, ok = Cast[*jwt.NumericDate](date); !ok {
		if d, ok = Cast[time.Time](date); ok {
			return jwt.NewNumericDate(d)
		}
		return nil
	}

	return t
}

type JwtClaimsImpl[V TimeOrNumericDateImpl] interface {
	GetDataAccess() *JwtClaimsDataAccess[V]
	Get(key string) (any, bool)
	Set(key string, value any) bool
	ToJwtToken() JwtTokenImpl
	ToJwtTokenString(secretKey string) (string, error)
	ParseNumericDate(key string) (*jwt.NumericDate, error)
	ParseString(key string) (string, error)
	ParseStringMany(key string) ([]string, error)
	GetIdentity() (string, error)
	SetIdentity(identity string) bool
	GetSubject() (string, error)
	SetSubject(subject string) bool
	GetIssued() (V, error)
	SetIssuedAt(date V) bool
	GetIssuer() (string, error)
	SetIssuer(issuer string) bool
	GetAudience() ([]string, error)
	SetAudience(audience []string) bool
	GetExpires() (V, error)
	SetExpiresAt(date V) bool
	GetSessionId() (string, error)
	SetSessionId(sessionId string) bool
	GetUser() (string, error)
	SetUser(user string) bool
	GetRole() (string, error)
	SetRole(role string) bool
	GetEmail() (string, error)
	SetEmail(email string) bool
}

type JwtClaimsDataAccessImpl[V TimeOrNumericDateImpl] interface {
	GetIdentity() string
	SetIdentity(identity string)
	GetSubject() string
	SetSubject(subject string)
	GetIssuer() string
	SetIssuer(issuer string)
	GetAudience() []string
	SetAudience(audience []string)
	GetIssued() V
	SetIssuedAt(date V)
	GetExpires() V
	SetExpiresAt(date V)
	GetSessionId() string
	SetSessionId(sessionId string)
	GetUser() string
	SetUser(user string)
	GetRole() string
	SetRole(role string)
	GetEmail() string
	SetEmail(email string)
}

type JwtClaimsDataAccess[V TimeOrNumericDateImpl] struct {
	ID        string   `json:"jti,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	Audience  []string `json:"aud,omitempty"`
	NotBefore V        `json:"nbf,omitempty"`
	IssuedAt  V        `json:"iat,omitempty"`
	ExpiresAt V        `json:"exp,omitempty"`
	SessionId string   `json:"sid,omitempty"`
	User      string   `json:"user,omitempty"`
	Role      string   `json:"role,omitempty"`
	Email     string   `json:"email,omitempty"`
}

func NewJwtClaimsDataAccess[V TimeOrNumericDateImpl](claims *jwt.RegisteredClaims) JwtClaimsDataAccessImpl[V] {
	return &JwtClaimsDataAccess[V]{
		ID:        claims.ID,
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
		Audience:  claims.Audience,
		NotBefore: Unwrap(CastAny(claims.NotBefore)),
		IssuedAt:  Unwrap(CastAny(claims.IssuedAt)),
		ExpiresAt: Unwrap(CastAny(claims.ExpiresAt)),
	}
}

func NewEmptyJwtClaimsDataAccess[V TimeOrNumericDateImpl]() JwtClaimsDataAccessImpl[V] {
	return new(JwtClaimsDataAccess[V])
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetIdentity() string {
	return claimsDataAccess.ID
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetIdentity(identity string) {
	claimsDataAccess.ID = identity
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetSubject() string {
	return claimsDataAccess.Subject
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetSubject(subject string) {
	claimsDataAccess.Subject = subject
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetIssuer() string {
	return claimsDataAccess.Issuer
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetIssuer(issuer string) {
	claimsDataAccess.Issuer = issuer
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetAudience() []string {
	return claimsDataAccess.Audience
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetAudience(audience []string) {
	claimsDataAccess.Audience = audience
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetIssued() V {
	return claimsDataAccess.IssuedAt
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetIssuedAt(date V) {
	claimsDataAccess.IssuedAt = date
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetExpires() V {
	return claimsDataAccess.ExpiresAt
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetExpiresAt(date V) {
	claimsDataAccess.ExpiresAt = date
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetSessionId() string {
	return claimsDataAccess.SessionId
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetSessionId(sessionId string) {
	claimsDataAccess.SessionId = sessionId
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetUser() string {
	return claimsDataAccess.User
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetUser(user string) {
	claimsDataAccess.User = user
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetRole() string {
	return claimsDataAccess.Role
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetRole(role string) {
	claimsDataAccess.Role = role
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) GetEmail() string {
	return claimsDataAccess.Email
}

func (claimsDataAccess *JwtClaimsDataAccess[V]) SetEmail(email string) {
	claimsDataAccess.Email = email
}

type JwtClaims[V TimeOrNumericDateImpl] struct {
	jwt.MapClaims
}

func NewJwtClaims[V TimeOrNumericDateImpl](claims jwt.MapClaims) JwtClaimsImpl[V] {
	return &JwtClaims[V]{
		claims,
	}
}

func EmptyJwtClaims[V TimeOrNumericDateImpl]() JwtClaimsImpl[V] {
	return &JwtClaims[V]{
		jwt.MapClaims{},
	}
}

func GetJwtClaimsFromJwtToken[V TimeOrNumericDateImpl](token JwtTokenImpl) (JwtClaimsImpl[V], error) {
	var ok bool
	var claims jwt.MapClaims
	KeepVoid(ok, claims)

	jwtToken := token.(*jwt.Token)
	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, ErrJwtClaimsInvalid
	}

	return NewJwtClaims[V](claims), nil
}

func (claims *JwtClaims[V]) GetDataAccess() *JwtClaimsDataAccess[V] {
	return &JwtClaimsDataAccess[V]{
		ID:        Unwrap(claims.GetIdentity()),
		Subject:   Unwrap(claims.GetSubject()),
		Issuer:    Unwrap(claims.GetIssuer()),
		Audience:  Unwrap(claims.GetAudience()),
		IssuedAt:  Unwrap(claims.GetIssued()),
		ExpiresAt: Unwrap(claims.GetExpires()),
		SessionId: Unwrap(claims.GetIdentity()),
		User:      Unwrap(claims.GetUser()),
		Role:      Unwrap(claims.GetRole()),
		Email:     Unwrap(claims.GetEmail()),
	}
}

func (claims *JwtClaims[V]) Get(key string) (any, bool) {
	var ok bool
	var value any
	KeepVoid(ok, value)

	value, ok = claims.MapClaims[key]
	return value, ok
}

func (claims *JwtClaims[V]) Set(key string, value any) bool {
	var ok bool
	var temp any
	KeepVoid(ok, temp)

	if temp, ok = claims.MapClaims[key]; !ok {
		return false
	}

	claims.MapClaims[key] = value
	return true
}

func (claims *JwtClaims[V]) ToJwtToken() JwtTokenImpl {
	return jwt.NewWithClaims(JwtSigningMethod, claims.MapClaims)
}

func (claims *JwtClaims[V]) ToJwtTokenString(secretKey string) (string, error) {
	var err error
	var tokenString string
	KeepVoid(err, tokenString)

	token := claims.ToJwtToken()
	if tokenString, err = token.SignedString([]byte(secretKey)); err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewNumericDateFromSeconds(f float64) *jwt.NumericDate {
	round, frac := math.Modf(f)
	return jwt.NewNumericDate(time.Unix(int64(round), int64(frac*1e9)))
}

func (claims *JwtClaims[V]) ParseNumericDate(key string) (*jwt.NumericDate, error) {
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

func (claims *JwtClaims[V]) ParseString(key string) (string, error) {
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

func (claims *JwtClaims[V]) ParseStringMany(key string) ([]string, error) {
	var ok bool
	var value any
	var temp []string
	KeepVoid(ok, value, temp)

	if value, ok = claims.Get(key); !ok {
		return nil, ErrDataTypeInvalid
	}

	switch value.(type) {
	case string:
		temp = []string{value.(string)}
	case []string:
		temp = value.([]string)
	case []any:
		values := value.([]any)
		temp = make([]string, len(values))
		for i, v := range values {
			if temp[i], ok = v.(string); !ok {
				return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
			}
		}
	default:
		return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	return temp, nil
}

func (claims *JwtClaims[V]) GetIdentity() (string, error) {
	return claims.ParseString(string(JwtClaimNameIdentity))
}

func (claims *JwtClaims[V]) SetIdentity(identity string) bool {
	return claims.Set(string(JwtClaimNameIdentity), identity)
}

func (claims *JwtClaims[V]) GetSubject() (string, error) {
	return claims.ParseString(string(JwtClaimNameSubject))
}

func (claims *JwtClaims[V]) SetSubject(subject string) bool {
	return claims.Set(string(JwtClaimNameSubject), subject)
}

func (claims *JwtClaims[V]) GetIssued() (V, error) {
	var err error
	var temp any
	KeepVoid(err, temp)

	temp, err = claims.ParseNumericDate(string(JwtClaimNameIssuedAt))
	return temp, err
}

func (claims *JwtClaims[V]) SetIssuedAt(date V) bool {
	return claims.Set(string(JwtClaimNameIssuedAt), GetNumericDateFromTimeOrNumericDateType(date))
}

func (claims *JwtClaims[V]) GetIssuer() (string, error) {
	return claims.ParseString(string(JwtClaimNameIssuer))
}

func (claims *JwtClaims[V]) SetIssuer(issuer string) bool {
	return claims.Set(string(JwtClaimNameIssuer), issuer)
}

func (claims *JwtClaims[V]) GetAudience() ([]string, error) {
	return claims.ParseStringMany(string(JwtClaimNameAudience))
}

func (claims *JwtClaims[V]) SetAudience(audience []string) bool {
	return claims.Set(string(JwtClaimNameAudience), audience)
}

func (claims *JwtClaims[V]) GetExpires() (V, error) {
	var err error
	var temp any
	KeepVoid(err, temp)

	temp, err = claims.ParseNumericDate(string(JwtClaimNameExpiresAt))
	return temp, err
}

func (claims *JwtClaims[V]) SetExpiresAt(date V) bool {
	return claims.Set(string(JwtClaimNameExpiresAt), GetNumericDateFromTimeOrNumericDateType(date))
}

func (claims *JwtClaims[V]) GetSessionId() (string, error) {
	return claims.ParseString(string(JwtClaimNameSessionId))
}

func (claims *JwtClaims[V]) SetSessionId(sessionId string) bool {
	return claims.Set(string(JwtClaimNameSessionId), sessionId)
}

func (claims *JwtClaims[V]) GetUser() (string, error) {
	return claims.ParseString(string(JwtClaimNameUser))
}

func (claims *JwtClaims[V]) SetUser(user string) bool {
	return claims.Set(string(JwtClaimNameUser), user)
}

func (claims *JwtClaims[V]) GetRole() (string, error) {
	return claims.ParseString(string(JwtClaimNameRole))
}

func (claims *JwtClaims[V]) SetRole(role string) bool {
	return claims.Set(string(JwtClaimNameRole), role)
}

func (claims *JwtClaims[V]) GetEmail() (string, error) {
	return claims.ParseString(string(JwtClaimNameEmail))
}

func (claims *JwtClaims[V]) SetEmail(email string) bool {
	return claims.Set(string(JwtClaimNameEmail), email)
}

func ParseJwtTokenUnverified(token string) (JwtTokenImpl, error) {
	var err error
	var parts []string
	var jwtToken JwtTokenImpl
	KeepVoid(err, parts, jwtToken)

	parser := jwt.NewParser()
	claims := jwt.MapClaims{}

	if jwtToken, parts, err = parser.ParseUnverified(token, claims); err != nil {
		return nil, ErrJwtTokenNotFound
	}

	return jwtToken, nil
}

func ParseJwtToken(token string, secretKey string) (JwtTokenImpl, error) {
	var err error
	var jwtToken JwtTokenImpl
	KeepVoid(err, jwtToken)

	parser := jwt.NewParser()
	claims := jwt.MapClaims{}
	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != JwtSigningMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	}

	if jwtToken, err = parser.ParseWithClaims(token, claims, keyFunc); err != nil {
		return nil, ErrJwtTokenNotFound
	}

	return jwtToken, nil
}

func ConvertJwtClaimsAccessDataToJwtClaims[V TimeOrNumericDateImpl](claimsDataAccess *JwtClaimsDataAccess[V]) JwtClaimsImpl[V] {
	claims := NewJwtClaims[V](jwt.MapClaims{})
	claims.SetIdentity(claimsDataAccess.GetIdentity())
	claims.SetSubject(claimsDataAccess.GetSubject())
	claims.SetIssuer(claimsDataAccess.GetIssuer())
	claims.SetAudience(claimsDataAccess.GetAudience())
	claims.SetIssuedAt(claimsDataAccess.GetIssued())
	claims.SetExpiresAt(claimsDataAccess.GetExpires())
	claims.SetSessionId(claimsDataAccess.GetSessionId())
	claims.SetUser(claimsDataAccess.GetUser())
	claims.SetRole(claimsDataAccess.GetRole())
	claims.SetEmail(claimsDataAccess.GetEmail())
	return claims
}

func ConvertJwtClaimsToJwtClaimsAccessData[V TimeOrNumericDateImpl](claims JwtClaimsImpl[V]) *JwtClaimsDataAccess[V] {
	claimsDataAccess := new(JwtClaimsDataAccess[V])
	claimsDataAccess.SetIdentity(Unwrap(claims.GetIdentity()))
	claimsDataAccess.SetSubject(Unwrap(claims.GetSubject()))
	claimsDataAccess.SetIssuer(Unwrap(claims.GetIssuer()))
	claimsDataAccess.SetAudience(Unwrap(claims.GetAudience()))
	claimsDataAccess.SetIssuedAt(Unwrap(claims.GetIssued()))
	claimsDataAccess.SetExpiresAt(Unwrap(claims.GetExpires()))
	claimsDataAccess.SetSessionId(Unwrap(claims.GetSessionId()))
	claimsDataAccess.SetUser(Unwrap(claims.GetUser()))
	claimsDataAccess.SetRole(Unwrap(claims.GetRole()))
	claimsDataAccess.SetEmail(Unwrap(claims.GetEmail()))
	return claimsDataAccess
}
