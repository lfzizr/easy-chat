package ctxdata

import "github.com/golang-jwt/jwt/v4"

const Identify = "easy-chat"

func GetJwtToken(secretKey string,iat,seconds int64,uid string) (string ,error) {
	// 定义claims
	claims := make(jwt.MapClaims)
	claims["exp"]=iat+seconds
	claims["identify"]=uid
	claims["iat"]=iat
	// 生成token 并签名 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	// 加密
	return token.SignedString([]byte(secretKey))
}