package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krvivek007/predictionservice/domain/prediction"
	"github.com/krvivek007/predictionservice/handlers"
	"github.com/krvivek007/predictionservice/mongo"
)

func main() {
	conn := os.Getenv("MONGO_ADDRESS")

	mongoClient, err := mongo.NewClient(conn)
	if err != nil {
		panic("Connection could not be established")
	}
	repoPrediction := prediction.NewRepository(mongoClient)
	engine := gin.New()
	engine.Use(gin.Recovery())

	handlers.InitializeRoutes(engine, repoPrediction)

	engine.Run(":8080")
}
