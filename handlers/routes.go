package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krvivek007/predictionservice/domain/prediction"
	"github.com/krvivek007/predictionservice/handlers/prediction_handlers"
)

func InitializeRoutes(e *gin.Engine, rp *prediction.Repository) {
	//Predicton API
	e.GET("v1/match/prediction/:fixtureId", prediction_handlers.GetFixturePredictionHandler(rp))
}
