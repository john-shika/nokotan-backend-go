package cores

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"time"
)

var ErrJwtTokenNotFound = errors.New("exception: JWT token not found")
var ErrJwtClaimsInvalid = errors.New("exception: Invalid JWT claims")
var ErrJwtIdentityNotFound = errors.New("exception: JWT identity not found")
var ErrJwtIssuedAtNotFound = errors.New("exception: JWT issued at not found")
var ErrJwtIssuerNotFound = errors.New("exception: JWT issuer not found")
var ErrJwtSubjectNotFound = errors.New("exception: JWT subject not found")
var ErrJwtExpiresNotFound = errors.New("exception: JWT expires not found")
var ErrJwtSessionIdNotFound = errors.New("exception: JWT session id not found")
var ErrJwtUserInvalid = errors.New("exception: JWT user not found")
var ErrJwtEmailInvalid = errors.New("exception: JWT email not found")
var ErrJwtSecretKeyNotFound = errors.New("exception: JWT secret key not found")

var JwtSigningMethod jwt.SigningMethod = jwt.SigningMethodHS256

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
	JwtClaimNameUser      JwtClaimNameTyped = "name"
	JwtClaimNameEmail     JwtClaimNameTyped = "email"
	JwtClaimNameRole      JwtClaimNameTyped = "role"
)

type JwtTokenImpl interface {
	SignedString(key interface{}) (string, error)
	SigningString() (string, error)
	EncodeSegment(seg []byte) string
}

type TimeOrJwtNumericDateImpl interface {
	*jwt.NumericDate | time.Time | string | int64 | int
}

func GetTimeUtcFromAny(value any) time.Time {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return Default[time.Time]()
	}
	value = val.Interface()
	switch value.(type) {
	case jwt.NumericDate:
		return value.(jwt.NumericDate).Time.UTC()
	case time.Time:
		return value.(time.Time).UTC()
	case string:
		return Unwrap(ParseTimeUtcByStringISO8601(value.(string)))
	case int64:
		return GetTimeUtcByTimeStamp(value.(int64))
	case int:
		return GetTimeUtcByTimeStamp(int64(value.(int)))
	default:
		return Default[time.Time]()
	}
}

func GetTimeUtcFromStrict[V TimeOrJwtNumericDateImpl](value V) time.Time {
	return GetTimeUtcFromAny(value)
}

func GetJwtNumericDateFromAny(value any) *jwt.NumericDate {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return nil
	}
	value = val.Interface()
	switch value.(type) {
	case jwt.NumericDate:
		t := value.(jwt.NumericDate)
		return &t
	case time.Time:
		return jwt.NewNumericDate(value.(time.Time))
	case string:
		t := Unwrap(ParseTimeUtcByStringISO8601(value.(string)))
		return jwt.NewNumericDate(t)
	case int64:
		t := GetTimeUtcByTimeStamp(value.(int64))
		return jwt.NewNumericDate(t)
	case int:
		t := GetTimeUtcByTimeStamp(int64(value.(int)))
		return jwt.NewNumericDate(t)
	default:
		return nil
	}
}

func GetJwtNumericDateFromStrict[V TimeOrJwtNumericDateImpl](value V) *jwt.NumericDate {
	return GetJwtNumericDateFromAny(value)
}

type JwtClaimsImpl interface {
	GetDataAccess() *JwtClaimsDataAccess
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
	GetIssued() (*jwt.NumericDate, error)
	SetIssuedAt(date any) bool
	GetIssuer() (string, error)
	SetIssuer(issuer string) bool
	GetAudience() ([]string, error)
	SetAudience(audience []string) bool
	GetExpires() (*jwt.NumericDate, error)
	SetExpiresAt(date any) bool
	GetSessionId() (string, error)
	SetSessionId(sessionId string) bool
	GetUser() (string, error)
	SetUser(user string) bool
	GetRole() (string, error)
	SetRole(role string) bool
	GetEmail() (string, error)
	SetEmail(email string) bool
}

type JwtClaimsDataAccessImpl interface {
	GetIdentity() string
	SetIdentity(identity string)
	GetSubject() string
	SetSubject(subject string)
	GetIssuer() string
	SetIssuer(issuer string)
	GetAudience() []string
	SetAudience(audience []string)
	GetIssued() *jwt.NumericDate
	SetIssuedAt(date any)
	GetExpires() *jwt.NumericDate
	SetExpiresAt(date any)
	GetSessionId() string
	SetSessionId(sessionId string)
	GetUser() string
	SetUser(user string)
	GetRole() string
	SetRole(role string)
	GetEmail() string
	SetEmail(email string)
}

type JwtClaimsDataAccess struct {
	ID        string           `json:"jti,omitempty"`
	Issuer    string           `json:"iss,omitempty"`
	Subject   string           `json:"sub,omitempty"`
	Audience  []string         `json:"aud,omitempty"`
	NotBefore *jwt.NumericDate `json:"nbf,omitempty"`
	IssuedAt  *jwt.NumericDate `json:"iat,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"exp,omitempty"`
	SessionId string           `json:"sid,omitempty"`
	User      string           `json:"name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Role      string           `json:"role,omitempty"`
}

func NewJwtClaimsDataAccess(claims *jwt.RegisteredClaims) JwtClaimsDataAccessImpl {
	return &JwtClaimsDataAccess{
		ID:        claims.ID,
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
		Audience:  claims.Audience,
		NotBefore: claims.NotBefore,
		IssuedAt:  claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}
}

func NewEmptyJwtClaimsDataAccess() JwtClaimsDataAccessImpl {
	return new(JwtClaimsDataAccess)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIdentity() string {
	return claimsDataAccess.ID
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIdentity(identity string) {
	claimsDataAccess.ID = identity
}

func (claimsDataAccess *JwtClaimsDataAccess) GetSubject() string {
	return claimsDataAccess.Subject
}

func (claimsDataAccess *JwtClaimsDataAccess) SetSubject(subject string) {
	claimsDataAccess.Subject = subject
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIssuer() string {
	return claimsDataAccess.Issuer
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIssuer(issuer string) {
	claimsDataAccess.Issuer = issuer
}

func (claimsDataAccess *JwtClaimsDataAccess) GetAudience() []string {
	return claimsDataAccess.Audience
}

func (claimsDataAccess *JwtClaimsDataAccess) SetAudience(audience []string) {
	claimsDataAccess.Audience = audience
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIssued() *jwt.NumericDate {
	return claimsDataAccess.IssuedAt
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIssuedAt(date any) {
	claimsDataAccess.IssuedAt = GetJwtNumericDateFromAny(date)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetExpires() *jwt.NumericDate {
	return claimsDataAccess.ExpiresAt
}

func (claimsDataAccess *JwtClaimsDataAccess) SetExpiresAt(date any) {
	claimsDataAccess.ExpiresAt = GetJwtNumericDateFromAny(date)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetSessionId() string {
	return claimsDataAccess.SessionId
}

func (claimsDataAccess *JwtClaimsDataAccess) SetSessionId(sessionId string) {
	claimsDataAccess.SessionId = sessionId
}

func (claimsDataAccess *JwtClaimsDataAccess) GetUser() string {
	return claimsDataAccess.User
}

func (claimsDataAccess *JwtClaimsDataAccess) SetUser(user string) {
	claimsDataAccess.User = user
}

func (claimsDataAccess *JwtClaimsDataAccess) GetRole() string {
	return claimsDataAccess.Role
}

func (claimsDataAccess *JwtClaimsDataAccess) SetRole(role string) {
	claimsDataAccess.Role = role
}

func (claimsDataAccess *JwtClaimsDataAccess) GetEmail() string {
	return claimsDataAccess.Email
}

func (claimsDataAccess *JwtClaimsDataAccess) SetEmail(email string) {
	claimsDataAccess.Email = email
}

type JwtClaims struct {
	jwt.MapClaims
}

func NewJwtClaims(claims jwt.MapClaims) JwtClaimsImpl {
	return &JwtClaims{
		claims,
	}
}

func EmptyJwtClaims() JwtClaimsImpl {
	return &JwtClaims{
		jwt.MapClaims{},
	}
}

func GetJwtClaimsFromJwtToken(token JwtTokenImpl) (JwtClaimsImpl, error) {
	var ok bool
	var claims jwt.MapClaims
	KeepVoid(ok, claims)

	jwtToken := token.(*jwt.Token)
	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, ErrJwtClaimsInvalid
	}

	return NewJwtClaims(claims), nil
}

func (claims *JwtClaims) GetDataAccess() *JwtClaimsDataAccess {
	return &JwtClaimsDataAccess{
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

func (claims *JwtClaims) Get(key string) (any, bool) {
	var ok bool
	var value any
	KeepVoid(ok, value)

	value, ok = claims.MapClaims[key]
	return value, ok
}

func (claims *JwtClaims) Set(key string, value any) bool {
	var ok bool
	var temp any
	KeepVoid(ok, temp)

	if temp, ok = claims.MapClaims[key]; !ok {
		return false
	}

	claims.MapClaims[key] = value
	return true
}

func (claims *JwtClaims) ToJwtToken() JwtTokenImpl {
	return jwt.NewWithClaims(JwtSigningMethod, claims.MapClaims)
}

func (claims *JwtClaims) ToJwtTokenString(secretKey string) (string, error) {
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

func (claims *JwtClaims) ParseNumericDate(key string) (*jwt.NumericDate, error) {
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

func (claims *JwtClaims) ParseString(key string) (string, error) {
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

func (claims *JwtClaims) ParseStringMany(key string) ([]string, error) {
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

func (claims *JwtClaims) GetIdentity() (string, error) {
	return claims.ParseString(string(JwtClaimNameIdentity))
}

func (claims *JwtClaims) SetIdentity(identity string) bool {
	return claims.Set(string(JwtClaimNameIdentity), identity)
}

func (claims *JwtClaims) GetSubject() (string, error) {
	return claims.ParseString(string(JwtClaimNameSubject))
}

func (claims *JwtClaims) SetSubject(subject string) bool {
	return claims.Set(string(JwtClaimNameSubject), subject)
}

func (claims *JwtClaims) GetIssued() (*jwt.NumericDate, error) {
	var err error
	var temp *jwt.NumericDate
	KeepVoid(err, temp)

	temp, err = claims.ParseNumericDate(string(JwtClaimNameIssuedAt))
	return temp, err
}

func (claims *JwtClaims) SetIssuedAt(date any) bool {
	return claims.Set(string(JwtClaimNameIssuedAt), GetJwtNumericDateFromAny(date))
}

func (claims *JwtClaims) GetIssuer() (string, error) {
	return claims.ParseString(string(JwtClaimNameIssuer))
}

func (claims *JwtClaims) SetIssuer(issuer string) bool {
	return claims.Set(string(JwtClaimNameIssuer), issuer)
}

func (claims *JwtClaims) GetAudience() ([]string, error) {
	return claims.ParseStringMany(string(JwtClaimNameAudience))
}

func (claims *JwtClaims) SetAudience(audience []string) bool {
	return claims.Set(string(JwtClaimNameAudience), audience)
}

func (claims *JwtClaims) GetExpires() (*jwt.NumericDate, error) {
	var err error
	var temp *jwt.NumericDate
	KeepVoid(err, temp)

	temp, err = claims.ParseNumericDate(string(JwtClaimNameExpiresAt))
	return temp, err
}

func (claims *JwtClaims) SetExpiresAt(date any) bool {
	return claims.Set(string(JwtClaimNameExpiresAt), GetJwtNumericDateFromAny(date))
}

func (claims *JwtClaims) GetSessionId() (string, error) {
	return claims.ParseString(string(JwtClaimNameSessionId))
}

func (claims *JwtClaims) SetSessionId(sessionId string) bool {
	return claims.Set(string(JwtClaimNameSessionId), sessionId)
}

func (claims *JwtClaims) GetUser() (string, error) {
	return claims.ParseString(string(JwtClaimNameUser))
}

func (claims *JwtClaims) SetUser(user string) bool {
	return claims.Set(string(JwtClaimNameUser), user)
}

func (claims *JwtClaims) GetRole() (string, error) {
	return claims.ParseString(string(JwtClaimNameRole))
}

func (claims *JwtClaims) SetRole(role string) bool {
	return claims.Set(string(JwtClaimNameRole), role)
}

func (claims *JwtClaims) GetEmail() (string, error) {
	return claims.ParseString(string(JwtClaimNameEmail))
}

func (claims *JwtClaims) SetEmail(email string) bool {
	return claims.Set(string(JwtClaimNameEmail), email)
}

func ParseJwtTokenUnverified(token string) (JwtTokenImpl, error) {
	var err error
	var parts []string
	var jwtToken JwtTokenImpl
	KeepVoid(err, parts, jwtToken)

	parser := jwt.NewParser()
	claims := make(jwt.MapClaims)

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

func CvtJwtClaimsAccessDataToJwtClaims(claimsDataAccess *JwtClaimsDataAccess) JwtClaimsImpl {
	claims := NewJwtClaims(jwt.MapClaims{})
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

func CvtJwtClaimsToJwtClaimsAccessData(claims JwtClaimsImpl) *JwtClaimsDataAccess {
	claimsDataAccess := new(JwtClaimsDataAccess)
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
