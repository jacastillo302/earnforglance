package middleware

import (
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/middleware"
	"earnforglance/server/internal/tokenutil"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareController struct {
	MiddlewareUsecase domain.MiddlewareUsecase
}

func (lc *MiddlewareController) GetPermissionsCustumer(c *gin.Context, custumerID string) ([]domain.Middleware, error) {
	permissions, err := lc.MiddlewareUsecase.GetPermissionsCustumer(c, custumerID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			fmt.Println("authorized", c.Request.URL.Path)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				/*
					slugs, err := tokenutil.ExtractSlugsFromToken(authToken, secret)
					if err != nil {
						c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: err.Error()})
						c.Abort()
						return
					}
					fmt.Println("slugs", slugs)
				*/

				// Check if the user has permission to access the requested path
				c.Set("x-user-id", userID)

				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
