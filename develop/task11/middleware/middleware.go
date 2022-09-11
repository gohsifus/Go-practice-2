package middleware

import (
	"fmt"
	"net/http"
	"task11/logger"
)

func Logging(f http.Handler, logger *logger.Log) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Запрос: " + r.URL.String())
		logger.Info("Запрос: " + r.URL.String())
		f.ServeHTTP(w, r)
	})
}
