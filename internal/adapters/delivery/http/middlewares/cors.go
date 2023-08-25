package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func VerifyOrigin(origin string) (bool, error) {
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:8080",
	}

	for _, allowedOrigin := range allowedOrigins {
		if strings.Compare(origin, allowedOrigin) == 0 {
			return true, nil
		}
	}

	return false, &echo.HTTPError{Code: 401, Message: "Você não tem permissão para acessar essa aplicação CORS"}
}

func OriginInspectSkipper(context echo.Context) bool {
	return false
}
