package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateSystemToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Check if the token is valid and not expired
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		// Check token expiration
		if time.Now().Unix() > claims.ExpiresAt {
			return echo.ErrUnauthorized
		}

		if claims.Type != "system" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}

func ValidateParkingToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Check if the token is valid and not expired
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		// Check token expiration
		if time.Now().Unix() > claims.ExpiresAt {
			return echo.ErrUnauthorized
		}

		if claims.Type != "parking" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}

func ValidateUserToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Check if the token is valid and not expired
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		// Check token expiration
		if time.Now().Unix() > claims.ExpiresAt {
			return echo.ErrUnauthorized
		}

		if claims.Type != "user" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}

func ValidateSystemTokenExpired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		if claims.Type != "system" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}

func ValidateParkingTokenExpired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		if claims.Type != "parking" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}

func ValidateUserTokenExpired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.ErrUnauthorized
			}
		}

		claims, ok := token.Claims.(*entity.CustomClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		if claims.Type != "user" {
			return echo.ErrUnauthorized
		}

		// Store the user information in the context for further use
		c.Set("user", claims)

		return next(c)
	}
}
