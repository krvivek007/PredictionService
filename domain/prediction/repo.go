package prediction

import (
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/krvivek007/predictionservice/mongo"
)

type Repository struct {
	mongoClient *mongo.Client
}

func NewRepository(mongoClient *mongo.Client) *Repository {
	repo := Repository{mongoClient}
	return &repo
}

const (
	databaseName   = "betfair_rapid_api_mapping"
	collectionName = "rapidapi_predictions"
)

func (r *Repository) GetPredictionByFixtureId(fixtureId string) (*PredictionResponse, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	intVar, _ := strconv.Atoi(fixtureId)
	query := bson.M{"fixtureId": intVar}

	var pred *PredictionResponse
	err := session.
		DB(databaseName).
		C(collectionName).
		Find(query).
		One(&pred)

	if err != nil {
		return nil, err
	}

	return pred, nil
}
