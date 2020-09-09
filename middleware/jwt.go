package middleware

import (
	"errors"
	"server/global"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrTokenExpired Token过期
	ErrTokenExpired = errors.New("Token过期")
	// ErrTokenInvalid Token错误
	ErrTokenInvalid = errors.New("Token错误")
)

// JWT 存储秘钥
type JWT struct {
	JWTKey []byte
}

// JWTClaim 存储claim,即用户信息
type JWTClaim struct {
	jwt.StandardClaims
	NickName string
}

// NewJWT 使用默认key创建jwt
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GCONFIG.JWTKey),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims JWTClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JWTKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JWTKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
			return nil, ErrTokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}
