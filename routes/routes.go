package routes

import (
	"ewalletgolang/db"
	"ewalletgolang/handler"
	"ewalletgolang/repository"
	"ewalletgolang/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	db := db.ConnectDB()

	userRepository := repository.NewRepository(db)
	userUsecase := usecase.NewUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	router.Run(":8888")
}