package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cakehandler "github.com/rohmatsyam/mongo_golang/cake/delivery/http"
	cakerepository "github.com/rohmatsyam/mongo_golang/cake/repository/mongo"
	cakeusecase "github.com/rohmatsyam/mongo_golang/cake/usecase"
	"github.com/rohmatsyam/mongo_golang/database"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("unable to load env")
	}
}

func main() {
	router := gin.Default()

	//	run db
	client, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	cakeCollection := database.GetCollection(client, os.Getenv("MONGO_COLLECTION"))

	cakeRepo := cakerepository.NewCakeRepository(cakeCollection)
	cakeUseCase := cakeusecase.NewCakeUseCase(cakeRepo)
	cakehandler.NewCakeHandler(router, cakeUseCase)

	router.Run(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")))
}
