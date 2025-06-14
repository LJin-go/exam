package Jwt

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/golang-jwt/jwt/v5"
//	"log"
//	"net/http"
//	"time"
//)
//
//var mySigningKey = []byte("AllYourBase")
//
//type MyCustomClaims struct {
//	UserId string `json:"user_id"`
//	jwt.RegisteredClaims
//}
//
//// GenerateToken 加密
//func GenerateToken(UserID string) (string, error) {
//	claims := MyCustomClaims{
//		UserID,
//		jwt.RegisteredClaims{
//			// Also fixed dates can be used for the NumericDate
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
//			Issuer:    "test",
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString(mySigningKey)
//}
//
//// ParseToken 解密
//func ParseToken(tokenString string) (*MyCustomClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte("AllYourBase"), nil
//	})
//	if err != nil {
//		log.Fatal(err)
//	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
//		fmt.Println(claims.UserId, claims.RegisteredClaims.Issuer)
//		return claims, nil
//	} else {
//		log.Fatal("unknown claims type, cannot proceed")
//	}
//	return nil, err
//}
//
//// GinAuthMiddleware 验证
//func GinAuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("token")
//		if tokenString == "" {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
//			return
//		}
//
//		claims, err := ParseToken(tokenString)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
//			return
//		}
//
//		// 将用户信息存入上下文
//		c.Set("user_id", claims.UserId)
//		c.Next()
//	}
//}
