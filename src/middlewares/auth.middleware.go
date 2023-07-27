package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      // логика извлечения и проверки токена
		next.ServeHTTP(w, r)
	})
}