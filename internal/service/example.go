package service

import (
	"go-quick-template/internal/model"
	"go-quick-template/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// DoLogin 用户认证
func DoLogin(ctx *gin.Context, param *AuthRequest) (*model.User, error) {

	user, err := ds.GetUserByUsername(param.Username)

	if err != nil {
		return user, nil
	}

	return nil, errcode.UnauthorizedAuthNotExist
}
