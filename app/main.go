package main

import (
	"github.com/gin-gonic/gin"
	cakehandler "github.com/rohmatsyam/mongo_golang/cake/delivery/http"
	cakerepository "github.com/rohmatsyam/mongo_golang/cake/repository/mongo"
	cakeusecase "github.com/rohmatsyam/mongo_golang/cake/usecase"
	"github.com/rohmatsyam/mongo_golang/database"
	"log"
)

func main() {
	router := gin.Default()

	//	run db
	client, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	cakeCollection := database.GetCollection(client, "cakes")

	cakeRepo := cakerepository.NewCakeRepository(cakeCollection)
	cakeUseCase := cakeusecase.NewCakeUseCase(cakeRepo)
	cakehandler.NewCakeHandler(router, cakeUseCase)

	router.Run("localhost:8000")
}
