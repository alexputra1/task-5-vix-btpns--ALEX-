package main

import (
	"github.com/alexputra1/golang_api/config"
	"github.com/alexputra1/golang_api/controller"
	"github.com/alexputra1/golang_api/repo"
	"github.com/alexputra1/golang_api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	jwtService service.JWTService = service.NewJWTService()
	userRepository repo.UserRepository = repo.NewUserRepository(db)
	authService service.AuthService = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}