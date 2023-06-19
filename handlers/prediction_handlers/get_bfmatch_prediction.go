package prediction_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krvivek007/predictionservice/domain/prediction"
)

var GetFixturePredictionHandler = func(r *prediction.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		fixtureId := c.Param("fixtureId")
		if len(fixtureId) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be empty!"})
			return
		}

		pred, err := r.GetPredictionByFixtureId(fixtureId)
		if err != nil {
			panic(err)
		}

		if pred == nil {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(http.StatusOK, pred)
	}
}
