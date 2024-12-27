package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"web/utils/config"
	"web/utils/errmsg"
)

var JwtKey = []byte(config.JwtKey)
var code int

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

	// 生成签名
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 2. 验证token
func CheckToken(token string) (*LoginClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*LoginClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		// token不能为空
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		// 验证token格式
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 验证token
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 判断token是否过期
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_EXPIRED
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
