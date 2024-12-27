package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
	"web/utils/config"
	"web/utils/errmsg"
)

var JwtKey = []byte(config.JwtKey)

type LoginClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 1.生成token
func SetToken(username string) (string, int) {
	// 过期时间
	expireTime := time.Now().Add(8 * time.Hour)
	SetClaims := LoginClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: &jwt.Time{Time: expireTime},
			Issuer:    "web",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 2. 验证token

// jwt中间件
