package jwtutil

import (
	"fmt"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// ParseTokenFromHeader parses token from Authorization header
func (j *Service) parseTokenFromHeader(c echo.Context) (*jwt.Token, error) {
	// Verify JWT
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("token not found")
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		return nil, fmt.Errorf("token invalid")
	}

	return j.parseToken(parts[1])
}

// ParseToken parses token from string
func (j *Service) parseToken(input string) (*jwt.Token, error) {
	return jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, fmt.Errorf("token method mismatched")
		}
		return j.key, nil
	})
}
