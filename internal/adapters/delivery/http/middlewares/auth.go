package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"regexp"
	"strings"
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		urlsNotNeedAuthorization := []string{
			"/api/user/login",
			"/api/user",
			"/api/docs/",
		}

		currentURL := context.Request().URL.Path

		for _, urlPattern := range urlsNotNeedAuthorization {
			if strings.HasPrefix(currentURL, urlPattern) {
				return next(context)
			}

			// Handle pattern with wildcard
			if strings.HasSuffix(urlPattern, "/*") {
				urlPrefix := strings.TrimSuffix(urlPattern, "/*")
				matched, err := regexp.MatchString("^"+regexp.QuoteMeta(urlPrefix), currentURL)
				if err == nil && matched {
					return next(context)
				}
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
