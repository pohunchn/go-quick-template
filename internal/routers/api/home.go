package api

import (
	_ "embed"

	"go-quick-template/pkg/app"
	"go-quick-template/pkg/debug"

	"github.com/gin-gonic/gin"
)

//go:embed assets/comic.ttf
var comic []byte

const MAX_PHONE_CAPTCHA = 10

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"BuildInfo": debug.ReadBuildInfo(),
	})
}
