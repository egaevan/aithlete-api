package middleware

import (
	"net/http"
	"strings"

	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func Auth(tokenSvc service.TokenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgMissingAuthHeader)
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgInvalidAuthHeader)
			}

			userID, err := tokenSvc.ValidateAccessToken(token)
			if err != nil {
				return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgInvalidToken)
			}

			c.Set("user_id", userID)
			return next(c)
		}
	}
}
