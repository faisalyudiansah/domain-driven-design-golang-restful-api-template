package route

import (
	"server/internal/auth/controller"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RefreshTokenControllRoute(c *controller.RefreshTokenController, r *gin.Engine, authMiddleware *middleware.AuthMiddleware) {
	g := r.Group("/auth")
	{
		g.POST("/logout", authMiddleware.Authorization(), c.Logout)
	}
}

func OauthControllerRoute(c *controller.OauthController, r *gin.Engine) {
	g := r.Group("/auth")
	{
		g.GET("/:provider/login", c.Login)
		g.GET("/:provider/callback", c.Callback)
		g.GET("/:provider/logout", c.Logout)
	}
}

func UserControllerRoute(c *controller.UserController, r *gin.Engine) {
	g := r.Group("/auth")
	{
		g.POST("/login", c.Login)
		g.POST("/register", c.Register)
		g.POST("/forgot-password", c.ForgotPassword)
		g.POST("/reset-password", c.ResetPassword)
		g.POST("/send-verification", c.SendVerification)
		g.POST("/verify-email", c.VerifyAccount)
	}
}
