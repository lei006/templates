package jwt

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

type TokenClaims struct {
	jwt.StandardClaims
	Data interface{}
}

type CustomClaims struct {
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

type JwtToken struct {
	KeepTime time.Duration //token 保持时间
	Issuer   string        //签发人
	SignKey  []byte
}

func MakeJwtToken(keepTime time.Duration, issuer string, signKey string) *JwtToken {

	tmp := &JwtToken{
		KeepTime: keepTime,
		Issuer:   issuer,
		SignKey:  []byte(signKey),
	}

	return tmp
}

// GenToken 生成JWT
func (jwtToken *JwtToken) GenToken(data interface{}) (string, error) {

	// 创建一个我们自己的声明
	c := TokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(jwtToken.KeepTime).Unix(), // 过期时间
			Issuer:    jwtToken.Issuer,
		},
		data,
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(jwtToken.SignKey)
}

// ParseToken 解析JWT
func (jwtToken *JwtToken) ParseToken(tokenString string) (interface{}, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtToken.SignKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
