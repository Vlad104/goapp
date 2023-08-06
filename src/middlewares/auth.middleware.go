package middlewares

import (
	"net/http"
	"log"
	"app/src/common"
	"app/src/entities"
	"github.com/golang-jwt/jwt/v5"
	"context"
	"time"
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	accessTokenHeader := r.Header.Get("Authorization") // получение данных из заголовка

		if len(accessTokenHeader) == 0 || accessTokenHeader[:7] != "Bearer " { // проверка, что токен начинается с корректного обозначения типа
			log.Printf("Could not get token %s", accessTokenHeader)
			common.HandleHttpError(w, common.ForbiddenError)
			return
		}

		accessTokenString := accessTokenHeader[7:] // извлечение самой строки токена
		token, err := jwt.ParseWithClaims(accessTokenString, &entities.AuthData{}, func(token *jwt.Token) (interface{}, error) {
			return common.SecretKey, nil
		})

		ctx := r.Context()

		if authData, ok := token.Claims.(*entities.AuthData); ok && token.Valid {
			expirationTime := time.Now().Add(common.ExpirationTime).Unix()

			if authData.CreatedAt > expirationTime {
				log.Print("accessToken timed out")
				common.HandleHttpError(w, common.ForbiddenError)
				return 
			}
			ctx = context.WithValue(ctx, "authData", authData)
		} else {
			log.Printf("%v", err)
			common.HandleHttpError(w, common.ForbiddenError)
      		return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}