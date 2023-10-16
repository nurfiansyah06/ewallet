package routes

import (
	"ewalletgolang/db"
	"ewalletgolang/handler"
	"ewalletgolang/middleware"
	"ewalletgolang/repository"
	"ewalletgolang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	db := db.ConnectDB()

	userRepository := repository.NewRepository(db)
	userUsecase := usecase.NewUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	walletRepository := repository.NewWalletRepository(db)
	walletUsecase := usecase.NewWalletUsecase(walletRepository)
	walletHandler := handler.NewWalletHandler(walletUsecase, userUsecase)

	transactionRepository := repository.NewRepository(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "page not found",
		})
	})

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.POST("/reset", userHandler.ResetPassword)

	router.GET("/profile", middleware.Authenticate(), userHandler.FindUserById)
	router.PUT("/topup", middleware.Authenticate(), walletHandler.TopUpWallet)
	router.POST("/transaction/:id", transactionHandler.AddAmount)

	router.Run(":8888")
}