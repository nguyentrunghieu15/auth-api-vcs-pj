package server

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyentrunghieu15/auth-api-vcs-pj/pkg/repository"
	"github.com/nguyentrunghieu15/common-vcs-prj/apu/auth"
)

type AuthServer struct {
	auth.AuthServiceServer
	UserRepo repository.UserRepositoryInterface
}

const (
	EXPIRY_TIME    = 30
	JWT_SECRET_KEY = "5TJZWKEWniPve33vlWYSoGYefn3gLQRzjfDlhSJ9ZKM"
)

func (a *AuthServer) Login(ctx context.Context, msg *auth.LoginMessage) (*auth.LoginResponseMessage, error) {
	if err := ctx.Err(); err == context.DeadlineExceeded {
		log.Printf("Auth Server - Deadline Exceeded : Request Login email %v\n", msg.Username)
	}
	isExistsUser, user, err := a.UserRepo.IsExsitsUserByEmail(msg.Username)
	if err != nil {
		log.Printf("Auth Server : Check exists error %v\n", err)
		return nil, err
	}

	if !isExistsUser {
		return nil, fmt.Errorf("Not found")
	}

	if CheckPasswordHash(msg.Password, user.Password) {
		accessToken, err := CreateAcessToken(msg)
		if err != nil {
			log.Printf("Auth Server : Create access token error %v\n", err)
			return nil, err
		}
		return &auth.LoginResponseMessage{AccessToken: accessToken}, nil
	} else {
		log.Printf("Auth Server : Unauthorized %v\n", err)
		return nil, fmt.Errorf("Auth Server : Unauthorized %v", err)
	}
}

func (a *AuthServer) ExtensionToken(ctx context.Context, token *auth.ExtensionAccessToken) (*auth.ReplyExtensionAccessToken, error) {
	if token != nil {
		oldAccessToken, err := jwt.Parse(token.OldAccessToken, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				log.Printf("Auth Server: there's an error with the signing method\n")
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return JWT_SECRET_KEY, nil

		})

		if err != nil {
			log.Printf("Auth Server: Parsing token errorn", err)
			return nil, err
		}
		claims, ok := oldAccessToken.Claims.(jwt.MapClaims)
		if ok && oldAccessToken.Valid {
			username := claims["username"].(string)

			newAccessToken, err := CreateAcessToken(&auth.LoginMessage{Username: username})
			if err != nil {
				log.Printf("Auth Server: Extension token error", err)
				return nil, err
			}
			return &auth.ReplyExtensionAccessToken{NewAccessToken: newAccessToken}, nil
		}

		log.Printf("Auth Server: unable to extract claims\n", err)
		return nil, err
	}
	return nil, nil
}
