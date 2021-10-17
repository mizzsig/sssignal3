package endlessTyping

import (
	"os"
	"time"
	"context"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// /soft/endless_typing のランキング
type EndlessTypingRanking struct {
	Time json.Number `bson:"time"`
	Success json.Number `bson:"success"`
	Failure json.Number `bson:"failure"`
	SuccessRate json.Number `bson:"successRate"`
	TypeRate json.Number `bson:"typeRate"`
}

func GetRanking(c echo.Context) error {
	// .envファイル読み込み
	err := godotenv.Load("../../.env")
	if err != nil { return err }
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb").SetAuth(credential))
	if err != nil { return err }

	collection := client.Database("sssignal3").Collection("endless_typing_ranking")
	allRanking := make(map[string]map[int]EndlessTypingRanking)
	allRanking["time"] = make(map[int]EndlessTypingRanking)
	allRanking["success"] = make(map[int]EndlessTypingRanking)
	allRanking["failure"] = make(map[int]EndlessTypingRanking)
	allRanking["successRate"] = make(map[int]EndlessTypingRanking)
	allRanking["typeRate"] = make(map[int]EndlessTypingRanking)
	
	// Time の上位10位を取得
	findOptions := options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"time", -1}})
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	index := 0
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		allRanking["time"][index] = result
		index++
	}
	
	// Success の上位10位を取得
	findOptions = options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"success", -1}})
	cur, err = collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	index = 0
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		allRanking["success"][index] = result
		index++
	}

	// Failure の上位10位を取得
	findOptions = options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"failure", 1}, {"success", -1}})
	cur, err = collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	index = 0
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		allRanking["failure"][index] = result
		index++
	}

	// SuccessRate の上位10位を取得
	findOptions = options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"successRate", -1}, {"success", -1}})
	cur, err = collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	index = 0
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		allRanking["successRate"][index] = result
		index++
	}

	// TypeRate の上位10位を取得
	findOptions = options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"typeRate", -1}})
	cur, err = collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	index = 0
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		allRanking["typeRate"][index] = result
		index++
	}

	return c.JSON(http.StatusOK, allRanking)
}

func PutRanking(c echo.Context) error {
	// .envファイル読み込み
	err := godotenv.Load("../../.env")
	if err != nil { return err }
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb").SetAuth(credential))
	if err != nil { return err }

	jsonBody := new(EndlessTypingRanking)
	if err := c.Bind(&jsonBody); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	collection := client.Database("sssignal3").Collection("endless_typing_ranking")

	insertFlag := false

	// Time の上位10位を取得
	findOptions := options.Find()
	findOptions.SetProjection(bson.D{{"_id", 0}})
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.D{{"time", -1}})
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil { return err }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result EndlessTypingRanking
		err := cur.Decode(&result)
		if err != nil { return err }
		jsonValue, _ := jsonBody.Time.Float64()
		mongoValue, _ := result.Time.Float64()
		if jsonValue > mongoValue {
			insertFlag = true
			break
		}
	}

	// Success の上位10位を取得
	if !insertFlag {
		findOptions = options.Find()
		findOptions.SetProjection(bson.D{{"_id", 0}})
		findOptions.SetLimit(10)
		findOptions.SetSort(bson.D{{"success", -1}})
		cur, err = collection.Find(ctx, bson.D{}, findOptions)
		if err != nil { return err }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result EndlessTypingRanking
			err := cur.Decode(&result)
			if err != nil { return err }
			jsonValue, _ := jsonBody.Success.Float64()
			mongoValue, _ := result.Success.Float64()
			if jsonValue > mongoValue {
				insertFlag = true
				break
			}
		}
	}

	// Failure の上位10位を取得
	if !insertFlag {
		findOptions = options.Find()
		findOptions.SetProjection(bson.D{{"_id", 0}})
		findOptions.SetLimit(10)
		findOptions.SetSort(bson.D{{"failure", 1}, {"success", -1}})
		cur, err = collection.Find(ctx, bson.D{}, findOptions)
		if err != nil { return err }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result EndlessTypingRanking
			err := cur.Decode(&result)
			if err != nil { return err }
			jsonValue, _ := jsonBody.Failure.Float64()
			mongoValue, _ := result.Failure.Float64()
			if jsonValue < mongoValue {
				insertFlag = true
				break
			}
		}
	}

	// SuccessRate の上位10位を取得
	if !insertFlag {
		findOptions = options.Find()
		findOptions.SetProjection(bson.D{{"_id", 0}})
		findOptions.SetLimit(10)
		findOptions.SetSort(bson.D{{"successRate", -1}, {"success", -1}})
		cur, err = collection.Find(ctx, bson.D{}, findOptions)
		if err != nil { return err }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result EndlessTypingRanking
			err := cur.Decode(&result)
			if err != nil { return err }
			jsonValue, _ := jsonBody.SuccessRate.Float64()
			mongoValue, _ := result.SuccessRate.Float64()
			if jsonValue > mongoValue {
				insertFlag = true
				break
			}
		}
	}

	// TypeRate の上位10位を取得
	if !insertFlag {
		findOptions = options.Find()
		findOptions.SetProjection(bson.D{{"_id", 0}})
		findOptions.SetLimit(10)
		findOptions.SetSort(bson.D{{"typeRate", -1}})
		cur, err = collection.Find(ctx, bson.D{}, findOptions)
		if err != nil { return err }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result EndlessTypingRanking
			err := cur.Decode(&result)
			if err != nil { return err }
			jsonValue, _ := jsonBody.TypeRate.Float64()
			mongoValue, _ := result.TypeRate.Float64()
			if jsonValue > mongoValue {
				insertFlag = true
				break
			}
		}
	}

	// ランクインした場合に、DBにINSERT
	if !insertFlag {
		return c.JSON(http.StatusOK, "not rank in")
	}
	result, err := collection.InsertOne(context.TODO(), jsonBody)
	if err != nil { return err }

	return c.JSON(http.StatusOK, result)
}
