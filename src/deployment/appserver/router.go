package appserver

import (
	"github.com/Mrityunjoy99/sample-go/src/application"
	"github.com/Mrityunjoy99/sample-go/src/application/admin"
	"github.com/Mrityunjoy99/sample-go/src/application/healthcheck"
	"github.com/Mrityunjoy99/sample-go/src/application/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.Engine, s application.Service) {
	registerHealthCheckRoutes(g)
	registerUserRoutes(g, s)
	adminRouteGroup(g, s)
}

func adminRouteGroup(g *gin.Engine, s application.Service) {
	if s.AdminService == nil {
		panic("AdminService is required for admin routes")
	}

	adminGroup := g.Group("/admin")
	// admin.Use(middleware.AdminAuth())
	// TODO: Implement and uncomment admin authentication middleware
	// adminGroup.Use(middleware.AdminAuth())
	adminController := admin.NewController(s.AdminService)
	adminGroup.POST("/generate-token", adminController.GenerateToken)
	adminGroup.POST("/validate-token", adminController.ValidateToken)
}

func registerHealthCheckRoutes(g *gin.Engine) {
	healthCheckController := healthcheck.NewController()
	g.GET("/health-check", healthCheckController.HealthCheck)
}

func registerUserRoutes(g *gin.Engine, s application.Service) {
	userController := user.NewController(s.UserService)
	g.GET("/user/:id", userController.GetUserById)
	g.POST("/user", userController.CreateUser)
	g.PUT("/user/:id", userController.UpdateUser)
	g.DELETE("/user/:id", userController.DeleteUser)
}
