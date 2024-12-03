package jwtutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userid string) (string, error) {
	// 创建 Access Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    userid,
		"realm":      "algohub", // 添加 realm 字段
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	})

	// 签名并返回令牌
	signedToken, err := token.SignedString([]byte("helloworld"))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

// VerifyToken 验证 JWT 令牌
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// 解析和验证令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("helloworld"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	return token, nil
}
