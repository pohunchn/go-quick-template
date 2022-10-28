package middleware

import (
	"go-quick-template/internal/model"
	"go-quick-template/pkg/app"
	"go-quick-template/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exist := c.Get("USER"); exist {
			if userModel, ok := user.(*model.User); ok {
				if userModel.Status == model.UserStatusNormal && userModel.IsAdmin {
					c.Next()
					return
				}
			}
		}

		response := app.NewResponse(c)
		response.ToErrorResponse(errcode.NoAdminPermission)
		c.Abort()
	}
}
