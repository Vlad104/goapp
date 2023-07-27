package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	//Логика извлечения и проверки токена будет реализована в https://github.com/users/Vlad104/projects/1?pane=issue&itemId=34022506
		next.ServeHTTP(w, r)
	})
}