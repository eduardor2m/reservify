package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		urlsNotNeedAuthorization := []string{
			"/api/user/login",
		}

		for _, url := range urlsNotNeedAuthorization {
			if url == context.Request().URL.Path {
				return next(context)
			}
		}

		authHeader := context.Request().Header.Get("Authorization")

		if authHeader == "" {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(authHeader[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		if !token.Valid {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		return next(context)
	}
}
