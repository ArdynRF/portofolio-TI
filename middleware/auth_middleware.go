package middleware

import (
	"net/http"
	"os"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/web"
	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) unauthorized(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.Response{
		Status: "unauthorized",
	}
	helper.WriteToBody(writer, webResponse)
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if request.RequestURI == "/api/v1/auth" || request.RequestURI == "/" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		tokenAuth := request.Header.Get("Authorization")
		if tokenAuth == "" {
			middleware.unauthorized(writer, request)
			return
		}

		var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
		claims := &web.TokenClaims{}

		token, err := jwt.ParseWithClaims(tokenAuth, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				middleware.unauthorized(writer, request)
				return
			}
		}
		if !token.Valid {
			middleware.unauthorized(writer, request)
			return
		}
		middleware.Handler.ServeHTTP(writer, request)
	}
}
