package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// MyClaims 自定义 结构体并内嵌jwt.StandardClaims
// Payload 载荷
// 想添加什么，随便来点
type MyClaims struct {
	Username           string `json:"username"`
	jwt.StandardClaims        // 标准jwt结构体
}

// 定义JWT过期时间
const TokenExpireDuration = time.Hour * 12

//const TokenExpireDuration = time.Second * 30

// 定义秘钥
var MySecret = []byte("jkflsdkljfdsjfsdfjlsdjfklsjdfkewrueiowruioweueiuriowuio@dfas")

// GenToken 生成JWT
func GenToken(username string) string {
	myClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "elkPorter",
		},
	}
	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	// 使用指定的secret签名
	// 获得完整的编码后的字符串token
	stringToken, err := token.SignedString(MySecret)
	if err != nil {
		log.Fatalf("GenToken failed, err: %v \n", err)
		return ""
	}
	return stringToken
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	fmt.Println(token.Claims)
	return nil, errors.New("invalid token")
}
