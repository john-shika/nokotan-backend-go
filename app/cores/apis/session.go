package apis

import (
	"crypto/sha256"
	"errors"
	"example/app/cores"
	"example/app/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
	"time"
)

func NewSession(db *gorm.DB, user *models.User) (*models.Session, string, error) {
	var err error
	var token string
	var claims jwt.MapClaims
	var secretKey string
	var secretKeyBytes []byte
	var jwtExpirationTime time.Duration
	//var jwtRefreshExpirationTime time.Duration

	if jwtExpirationTime = viper.GetDuration("jwt_expiration_time"); jwtExpirationTime == 0 {
		return nil, token, errors.New("jwt expiration time not found")
	}

	//if jwtRefreshExpirationTime = viper.GetDuration("jwt_refresh_expiration_time"); jwtRefreshExpirationTime == 0 {
	//	return nil, errors.New("jwt refresh expiration time not found")
	//}

	sessionId := cores.NewUuid()
	timeUtcNow := time.Now().UTC()
	timeExpiredAt := timeUtcNow.Add(jwtExpirationTime * time.Second)
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// set claims
	if claims, err = GetJwtClaims(jwtToken); err != nil {
		return nil, token, err
	}

	// set claims
	claims["iat"] = timeUtcNow.UnixMilli()
	claims["exp"] = timeExpiredAt.UnixMilli()
	claims["did"] = sessionId
	claims["user"] = user.Username

	if secretKeyBytes, err = NewSecretKey(); err != nil {
		return nil, token, err
	}

	secretKey = cores.NewBase64EncodeToString(secretKeyBytes)

	// signed jwt token
	if token, err = jwtToken.SignedString(secretKeyBytes); err != nil {
		return nil, token, err
	}

	// create a new session
	session := models.Session{}
	session.Model = models.Model{}
	session.UUID = sessionId
	session.UserID = user.ID
	session.SecretKey = secretKey
	session.ExpiredAt = timeExpiredAt

	// save session
	if tx := db.Create(&session); tx.Error != nil {
		panic(tx.Error.Error())
	}

	// return session
	return &session, token, nil
}

func NewSecretKey() ([]byte, error) {
	var err error
	var secretKey string
	var buff []byte
	var salted []byte

	// get jwt secret key
	if secretKey = strings.Trim(viper.GetString("jwt_secret_key"), " "); secretKey == "" {
		return nil, cores.ErrJwtSecretKeyNotFound
	}

	// convert secret key to bytes
	buff = []byte(secretKey)

	// create new salt
	if salted, err = cores.NewRandomBytes(32); err != nil {
		return nil, err
	}

	// combine secret key and salt
	hash := sha256.New()
	hash.Write(buff)
	hash.Write(salted)

	// assign new secret key in bytes
	buff = hash.Sum(nil)

	// return
	return buff, nil
}

func GetJwtToken(c echo.Context) (*jwt.Token, error) {
	var ok bool
	var err error
	var token string
	var jwtToken *jwt.Token
	if token = strings.Trim(c.Request().Header.Get("Authorization"), " "); token == "" {
		return nil, cores.ErrJwtTokenNotFound
	}
	if token, ok = strings.CutPrefix(token, "Bearer "); !ok {
		return nil, cores.ErrJwtTokenNotFound
	}
	if token = strings.Trim(token, " "); token == "" {
		return nil, cores.ErrJwtTokenNotFound
	}
	if jwtToken, err = ParseJwtToken(token); err != nil {
		return nil, err
	}
	return jwtToken, nil
}

func ParseJwtToken(token string) (*jwt.Token, error) {
	var err error
	var parts []string
	var jwtToken *jwt.Token
	cores.KeepVoid(err, parts, jwtToken)

	parser := jwt.NewParser()
	claims := jwt.MapClaims{}

	if jwtToken, parts, err = parser.ParseUnverified(token, claims); err != nil {
		return nil, cores.ErrJwtTokenNotFound
	}

	return jwtToken, nil
}

func ValidateJwtToken(db *gorm.DB, jwtToken *jwt.Token) (*models.Session, *jwt.Token, error) {
	var ok bool
	var err error
	var claims jwt.MapClaims
	var secretKeyBytes []byte
	var session *models.Session
	var sessionId string
	if claims, err = GetJwtClaims(jwtToken); err != nil {
		return nil, nil, err
	}
	if sessionId, err = GetSessionIdFromJwtClaims(claims); err != nil {
		return nil, nil, err
	}
	if session, err = GetSessionModel(db, sessionId); err != nil {
		return nil, nil, err
	}
	// get secret key
	if secretKeyBytes, err = cores.NewBase64DecodeToBytes(session.SecretKey); err != nil {
		return nil, nil, err
	}
	// parse jwt
	parser := jwt.NewParser()
	if jwtToken, err = parser.Parse(jwtToken.Raw, func(jwtToken *jwt.Token) (interface{}, error) {
		var hmac *jwt.SigningMethodHMAC
		cores.KeepVoid(hmac)

		if hmac, ok = jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return secretKeyBytes, nil
	}); err != nil {
		return nil, nil, err
	}
	// check validation token
	if !jwtToken.Valid {
		// token invalid
		return nil, nil, jwt.ErrTokenSignatureInvalid
	}
	return session, jwtToken, nil
}

func GetJwtClaims(jwtToken *jwt.Token) (jwt.MapClaims, error) {
	var ok bool
	var claims jwt.MapClaims
	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, cores.ErrJwtClaimsInvalid
	}
	return claims, nil
}

// GetIatFromJwtClaims method get iat from jwt claims, prefer use `claims.GetIssued()` method
func GetIatFromJwtClaims(claims jwt.MapClaims) (time.Time, error) {
	var ok bool
	var iat int64
	if iat, ok = claims["iat"].(int64); !ok {
		return time.Time{}, ErrJwtIatNotFound
	}
	return time.UnixMilli(iat), nil
}

// GetExpFromJwtClaims method get exp from jwt claims, prefer use `claims.GetExpirationTime()` method
func GetExpFromJwtClaims(claims jwt.MapClaims) (time.Time, error) {
	var ok bool
	var exp int64
	if exp, ok = claims["exp"].(int64); !ok {
		return time.Time{}, cores.ErrJwtExpiresNotFound
	}
	return time.UnixMilli(exp), nil
}

func GetSessionIdFromJwtClaims(claims jwt.MapClaims) (string, error) {
	var ok bool
	var sessionId string
	if sessionId, ok = claims["session_id"].(string); !ok {
		return "", cores.ErrJwtSessionIdNotFound
	}
	if sessionId = strings.Trim(sessionId, " "); sessionId == "" {
		return "", cores.ErrJwtSessionIdNotFound
	}
	return sessionId, nil
}

func GetUserIdFromJwtClaims(claims jwt.MapClaims) (string, error) {
	var ok bool
	var userId string
	if userId, ok = claims["user_id"].(string); !ok {
		return "", cores.ErrJwtUserInvalid
	}
	if userId = strings.Trim(userId, " "); userId == "" {
		return "", cores.ErrJwtUserInvalid
	}
	return userId, nil
}

func GetSessionModel(db *gorm.DB, sessionId string) (*models.Session, error) {
	var session models.Session
	if tx := db.First(&session, "uuid = ?", sessionId); tx.Error != nil {
		// db error
		return nil, tx.Error
	}
	// get current time in UTC
	timeUtcNow := time.Now().UTC()
	// check token expired
	if timeUtcNow.After(session.ExpiredAt) {
		// delete current session
		if tx := db.Delete(&session); tx.Error != nil {
			// db error
			return nil, tx.Error
		}
		// reject token
		return nil, jwt.ErrTokenExpired
	}
	// return
	return &session, nil
}
