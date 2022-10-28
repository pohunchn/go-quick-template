package routers

import (
	"net/http"
	"path/filepath"

	"go-quick-template/internal/conf"
	"go-quick-template/internal/middleware"
	"go-quick-template/internal/routers/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// 按需注册 docs、静态资源、LocalOSS 路由
	{
		registerStatick(e)
		routeLocalOSS(e)
	}

	// v1 group api
	r := e.Group("/v1")

	// 获取version
	r.GET("/", api.Version)

	// 无鉴权路由组
	noAuthApi := r.Group("/")
	{
		// 获取用户基本信息
		noAuthApi.GET("/", api.Version)
	}

	// 宽松鉴权路由组
	looseApi := r.Group("/").Use(middleware.JwtLoose())
	{
		// 获取用户动态列表
		looseApi.GET("/loose", api.Version)
	}

	// 鉴权路由组
	authApi := r.Group("/").Use(middleware.JWT())
	adminApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Admin())
	{
		// 同步索引
		authApi.GET("/auth", api.Version)

		adminApi.GET("/admin", api.Version)

	}

	// 其他路由注册
	otherRoute(r, authApi)

	// 默认404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	return e
}

// routeLocalOSS register LocalOSS route if neeed
func routeLocalOSS(e *gin.Engine) {
	if !conf.CfgIf("LocalOSS") {
		return
	}

	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("get localOSS save path err: %v", err)
	}
	e.Static("/oss", savePath)

	logrus.Infof("register LocalOSS route in /oss on save path: %s", savePath)
}

func otherRoute(public gin.IRoutes, authApi gin.IRoutes) {
	if !conf.CfgIf("Alipay") {
		return
	}

	// 支付宝回调
	public.POST("/other/public", api.Version)

	// 用户充值
	authApi.POST("/other/auth", api.Version)
}
