package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// CreateToken 函数用于创建 JWT，其中包含指定用户ID，并返回生成的令牌字符串
func CreateToken(user_id int) string {
	// 创建一个新的 JWT，并指定签名方法为 HS256。
	token := jwt.New(jwt.SigningMethodHS256)

	// 创建 JWT 的声明，包括过期时间（exp）、签发时间（iat）和用户ID
	claims := make(jwt.MapClaims)
	token_exp, _ := beego.AppConfig.String("Tokenexp")
	fmt.Println("token_exp:", token_exp)
	tokenexp, _ := strconv.Atoi(token_exp)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix() // 设置过期时间为当前时间加上 Tokenexp 的值
	claims["iat"] = time.Now().Unix()                                          // 设置签发时间为当前时间
	claims["user_id"] = user_id                                                // 设置用户ID
	token.Claims = claims

	// 使用配置中的密钥对 JWT 进行签名，并返回生成的令牌字符串。
	token_secrets, _ := beego.AppConfig.String("TokenSecrets")
	tokenString, _ := token.SignedString([]byte(token_secrets))
	return tokenString
}

// CheckToken 函数用于检查 JWT 的有效性，并返回其中包含的用户ID
func CheckToken(tokenString string) (error) {
	// 解析传入的 JWT，验证签名是否有效，并获取其中的声明。
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		token_secrets, _ := beego.AppConfig.String("TokenSecrets")
		return []byte(token_secrets), nil
	})
	if err != nil {
		return err
	}
	// 从声明中获取用户ID，并转换为整数后返回。
	// claims, _ := token.Claims.(jwt.MapClaims)
	// user_id := int(claims["user_id"].(float64))
	// return user_id, nil
	return nil
}

// GetUserId 函数用于获取 JWT 中的用户ID，不进行令牌的有效性检查
func GetUserId(tokenString string) int {
	// 解析传入的 JWT，并获取其中的声明。
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		token_secrets, _ := beego.AppConfig.String("TokenSecrets")
		return []byte(token_secrets), nil
	})

	// 从声明中获取用户ID，并转换为整数后返回。
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	return user_id
}
