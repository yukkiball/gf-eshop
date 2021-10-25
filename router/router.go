package router

import (
	"gf-eshop/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Use(MiddlewareCORS)
	s.Group("/item", func(group *ghttp.RouterGroup) {
		group.POST("/create", api.Item.CreateItem)
		group.GET("/get", api.Item.GetItem)
		group.GET("/list", api.Item.ListItem)
		group.GET("/publishPromo", api.Item.PublishPromo)
	})
	s.Group("/order", func(group *ghttp.RouterGroup) {
		group.ALL("/generateVerifyCode", api.Order.GenerateVerifyCode)
		group.POST("/generateToken", api.Order.GenerateToken)
		group.POST("/create", api.Order.CreateOrder)
	})
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.POST("/login", api.User.Login)
		group.POST("/register", api.User.Register)
		group.POST("getOtp", api.User.GetOtp)
		group.GET("/getUser", api.User.GetUser)
	})
}
