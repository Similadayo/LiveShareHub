package main

import (
	"github.com/gin-gonic/gin"
	"github.com/similadayo/internal/user"
	"github.com/similadayo/pkg/auth"
	"github.com/similadayo/pkg/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logger := logging.NewLogger()

	//initialize sqlite database
	db, err := gorm.Open(sqlite.Open("User.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", map[string]interface{}{
			"error": err.Error(),
		})
	}

	//auto migrate db
	db.AutoMigrate(&user.User{})

	//Initialize gin router
	r := gin.Default()

	//Initialize user repository
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, logger)
	userHandler := user.NewHandler(userService)

	// initialize auth package
	authService := auth.AuthMiddleWare(userService)

	//API Routes
	api := r.Group("/api")
	{
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/", userHandler.Register)
		}
	}

	//apply middleware with the logger and auth
	r.Use(auth.LoggerMiddleWare(logger), authService)
	apiAuth := r.Group("/api/auth")
	{
		userRoutes := apiAuth.Group("/users")
		{
			userRoutes.POST("/", userHandler.GetUserByIDHandler)
		}
	}

	r.Run(":8081")
}
