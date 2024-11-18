package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密密码，防止数据库泄漏
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 安全、无差错地传递数据，判断用户是否登录状态
func GenerateJWT(Userid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": Userid,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), //token有效期为7天
	})
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err //添加token前缀
}

// 验证密码是否正确
func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	return err == nil
}
