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
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		logger.Fatal("failed to migrate database", map[string]interface{}{
			"error": err.Error(),
		})
	}

	//Initialize gin router
	r := gin.Default()

	//Initialize user repository
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, logger)
	userHandler := user.NewHandler(userService)

	// initialize auth package
	authService := auth.AuthMiddleware()

	//API Routes
	r.Use(auth.LoggerMiddleWare(logger))
	api := r.Group("/api")
	{
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/", userHandler.Register)
			userRoutes.POST("/login", userHandler.Login)
		}
	}

	//apply middleware with the logger and auth
	r.Use(auth.LoggerMiddleWare(logger), authService)
	apiAuth := r.Group("/api/auth")
	{
		userRoutes := apiAuth.Group("/users")
		{
			userRoutes.GET("/user/:username", userHandler.GetUserByUserNameHandler)
			userRoutes.GET("/:id", userHandler.GetUserByIDHandler)
			userRoutes.PUT("/:id", userHandler.UpdateUserHandler)
			userRoutes.DELETE("/:id", userHandler.DeleteUserHandler)
			userRoutes.GET("/profile", userHandler.GetUserProfileHandler)
			userRoutes.GET("/filter/:user", userHandler.FilterUserByNameHandler)
		}
	}

	r.Run(":8081")
}
