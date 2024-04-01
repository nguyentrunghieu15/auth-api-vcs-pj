package server

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyentrunghieu15/common-vcs-prj/apu/auth"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateAcessToken(msg *auth.LoginMessage) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(EXPIRY_TIME * time.Minute)
	claims["authorized"] = true
	claims["username"] = msg.Username
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		log.Printf("Auth Func: Cant create access token error %v\n", err)
		return "", err
	}

	return tokenString, nil
}
