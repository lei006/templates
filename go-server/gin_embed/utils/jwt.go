package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("Token is expired")            //过期.
	TokenNotValidYet = errors.New("Token not active yet")        // 无效,
	TokenMalformed   = errors.New("That's not even a token")     //不是一个令牌
	TokenInvalid     = errors.New("Couldn't handle this token:") //无法处理
)

type JwtToken struct {
	claims jwt.StandardClaims
}

func MakeJwtToken(subject string, issuer string) *JwtToken {

	tmp := &JwtToken{
		claims: jwt.StandardClaims{
			Issuer:  issuer,
			Subject: subject,
		},
	}

	return tmp
}

type JwtClaims struct {
	Info interface{} `json:"info"`
	jwt.StandardClaims
}

func (jwtToken *JwtToken) JwtEncode(info interface{}, secret string, timeout time.Duration) (string, error) {

	expireTime := time.Now().Add(timeout) //过期时间取值
	claims := JwtClaims{
		info,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间 int64
			Issuer:    jwtToken.claims.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (jwtToken *JwtToken) JwtDecode(tokenString string, secret string) (*JwtClaims, error) {

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
