package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/joho/godotenv"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"os"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 問い合わせフォームから送られるリクエスト
type InquiryRequest struct {
	Address string `json:"address"`
	Body string `json:"body"`
}

// Slack webhookに渡す用の構造体
type Slack struct {
	Text string `json:"text"`
}

// ギャラリーに出す画像
type Image struct {
	ImageUrl string `bson:"imageUrl"`
	TweetUrl string `bson:"tweetUrl"`
	Text string `bson:"text"`
	Comment string `bson:"comment"`
	Date string `bson:"date"`
}

// ギャラリーに出す動画
type Movie struct {
	MovieUrl string `bson:"movieUrl"`
	TweetUrl string `bson:"tweetUrl"`
	Text string `bson:"text"`
	Comment string `bson:"comment"`
	Date string `bson:"date"`
}

func main() {
	e := echo.New()

	// CORS用の設定
	e.Use(middleware.CORS())

	// .envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env loading error!")
	}

	slackWebhook := os.Getenv("SLACK_WEBHOOK")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!!!")
	})
	e.GET("/gallery/images", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: "root", Password: "pass"}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		images := session.DB("sssignal3").C("images")

		// Query発行
		var imageList []Image
		images.Find(bson.M{}).Sort("-_id").All(&imageList)

		session.Close()
		return c.JSON(http.StatusOK, imageList)
	})
	e.GET("/gallery/movies", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: "root", Password: "pass"}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		images := session.DB("sssignal3").C("movies")

		// Query発行
		var movieList []Movie
		images.Find(bson.M{}).Sort("-_id").All(&movieList)

		session.Close()
		return c.JSON(http.StatusOK, movieList)
	})
	e.POST("/inquiry", func(c echo.Context) error {
		// パラメータを受け取る
		inquiryRequest := new(InquiryRequest)
		if err := c.Bind(inquiryRequest); err != nil {
			return c.String(http.StatusOK, "err!")
		}
		if inquiryRequest.Address != "" {
			inquiryRequest.Address = "*" + inquiryRequest.Address + "*"
		}

		// TODO : 見た目もう少し良くしたい 
		// https://api.slack.com/tools/block-kit-builder
		slack := new(Slack)
		slack.Text = "*【問い合わせフォーム】*\naddress : " + inquiryRequest.Address + "\ntext : " + inquiryRequest.Body

		// 構造体をJSONに変換
		json, _:= json.Marshal(slack)

		// JSONちゃんと作れてるかの確認用
		// return c.JSON(http.StatusOK, url.Values{"payload": {string(json)}})

		// Slack webhookにPOSTリクエスト送る
		response, err := http.PostForm(slackWebhook, url.Values{"payload": {string(json)}})
		// レスポンス読み込み
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return c.String(http.StatusOK, "err!")
		}

		// webhookの結果をJSONで返す
		return c.JSON(http.StatusOK, map[string]string{"result": string(responseBody)})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
