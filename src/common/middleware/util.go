package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResError(c *gin.Context, err error) {
	panic(err.Error())
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
