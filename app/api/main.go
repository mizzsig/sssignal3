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
	"time"
	"math/rand"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"./soft"
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

// 広告
type Advertisement struct {
	Html string `bson:"html"`
}

// 記事
type Column struct {
	Date string `bson:"date"`
	Url string `bson:"url"`
	Title string `bson:"title"`
	Body string `bson:"body"`
	CharacterComment string `bson:"characterComment"`
	ShowCharacter bool `bson:"showCharacter"`
	ImageUrl string `bson:"imageUrl"`
}

// /soft/endless_typing の単語
type EndlessTypingWord struct {
	Word string `bson:"word"`
}

func main() {
	e := echo.New()

	// .envファイル読み込み
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(".env loading error!")
	}

	// CORSの設定, 本番は正しいドメインからしかAPIにアクセスできないようにする
	if os.Getenv("ENVIRONMENT") == "production" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://sssignal.com", "https://ver3.sssignal.com"},
		}))
	} else {
		e.Use(middleware.CORS())
	}

	slackWebhook := os.Getenv("SLACK_WEBHOOK")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	apiPassword := os.Getenv("API_PASSWORD")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!!!")
	})
	e.GET("/gallery/images", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
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
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		movies := session.DB("sssignal3").C("movies")

		// Query発行
		var movieList []Movie
		movies.Find(bson.M{}).Sort("-_id").All(&movieList)

		session.Close()
		return c.JSON(http.StatusOK, movieList)
	})
	e.POST("/inquiry", func(c echo.Context) error {
		// パラメータを受け取る
		inquiryRequest := new(InquiryRequest)
		if err := c.Bind(inquiryRequest); err != nil {
			return c.String(http.StatusOK, "error!")
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
			return c.String(http.StatusOK, "error!")
		}

		// webhookの結果をJSONで返す
		return c.JSON(http.StatusOK, map[string]string{"result": string(responseBody)})
	})
	// ランダムでツイートを１件返す
	e.GET("/tweet", func(c echo.Context) error {
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)

		rand.Seed(time.Now().UnixNano())
		// データを取ってくる対象のCollection
		target := [...] string{"images", "movies"}

		// ランダムで選んだCollectionに対してデータを取得する
		switch target[rand.Intn(2)] {
			case "images":
				images := session.DB("sssignal3").C("images")

				// Query発行
				var imageList []Image
				images.Find(bson.M{}).Sort("-_id").All(&imageList)
				count, err := images.Count()
				if err != nil {
					return c.String(http.StatusOK, "error!")
				}
				
				result := imageList[rand.Intn(count)]

				session.Close()
				return c.JSON(http.StatusOK, result)
			case "movies":
				movies := session.DB("sssignal3").C("movies")

				// Query発行
				var movieList []Movie
				movies.Find(bson.M{}).Sort("-_id").All(&movieList)
				count, err := movies.Count()
				if err != nil {
					return c.String(http.StatusOK, "error!")
				}
				
				result := movieList[rand.Intn(count)]

				session.Close()
				return c.JSON(http.StatusOK, result)
			default:
				session.Close()
				return c.String(http.StatusOK, "error!")
		}
	})
	// 表示させる広告を取得
	e.GET("/advertisement", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		advertisements := session.DB("sssignal3").C("advertisements")
		// 表示する広告の数
		viewAdvertisementNum := 2

		var advertisementList []Advertisement
		advertisements.Find(bson.M{}).Sort("-_id").All(&advertisementList)
		count, err := advertisements.Count()
		if err != nil {
			return c.String(http.StatusOK, "error!")
		}

		var advertisementIndex []int
		var advertisementResult []Advertisement

		for len(advertisementIndex) < viewAdvertisementNum {
			index := rand.Intn(count)
			var isEqual = false

			// 同一のインデックスが既に選ばれているかを確認
			for i := 0; i < len(advertisementIndex); i++ {
				if (index == advertisementIndex[i]) {
					isEqual = true
				}
			}

			// インデックスと結果配列に格納する
			if !isEqual {
				advertisementIndex = append(advertisementIndex, index)
				advertisementResult = append(advertisementResult, advertisementList[index])
			}
		}

		session.Close()
		return c.JSON(http.StatusOK, advertisementResult)
	})
	// 記事の一覧取得
	e.GET("/column/list", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		columns := session.DB("sssignal3").C("columns")

		// Query発行
		var columnList []Column
		columns.Find(bson.M{}).Sort("-_id").All(&columnList)

		session.Close()
		return c.JSON(http.StatusOK, columnList)
	})
	// 記事の詳細取得
	e.GET("/column/:url", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		columns := session.DB("sssignal3").C("columns")

		// 一件取得
		var column Column
		columns.Find(bson.M{"url": c.Param("url")}).One(&column)

		session.Close()
		return c.JSON(http.StatusOK, column)
	})
	// 記事の登録・更新
	e.POST("/column/:url/:pass", func(c echo.Context) error {
		// POSTリクエストのパラメータを取得
		postColumn := new(Column)
		if err = c.Bind(postColumn); err != nil {
			return c.JSON(http.StatusOK, "error!")
		}

		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		columns := session.DB("sssignal3").C("columns")

		// 記事取得
		var column Column
		columns.Find(bson.M{"url": c.Param("url")}).One(&column)

		if apiPassword == c.Param("pass") {
			if column.Url == c.Param("url") {
				// 既にある場合は更新
				colQuerier := bson.M{"url": column.Url}
				updateValue := bson.M{"$set": postColumn}

				err = columns.Update(colQuerier, updateValue)
				if err != nil {
					return c.JSON(http.StatusOK, "error!")
				}
			} else {
				// 無い場合は新規登録
				err = columns.Insert(postColumn)
				if err != nil {
					return c.JSON(http.StatusOK, "error!")
				}
			}
		}

		session.Close()
		return c.JSON(http.StatusOK, "r")
	})
	// 記事の削除
	e.DELETE("/column/:url/:pass", func(c echo.Context) error {
		// MongoDBにログイン, DBとCollectionを指定
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		columns := session.DB("sssignal3").C("columns")

		if apiPassword == c.Param("pass") {
			colQuerier := bson.M{"url": c.Param("url")}
			err = columns.Remove(colQuerier)
			if err != nil {
				return c.JSON(http.StatusOK, "error!")
			}
		}

		session.Close()
		return c.JSON(http.StatusOK, "r")
	})
	// /soft/endless_typing 用の単語を取得
	e.GET("/soft/endless_typing/all_words", func(c echo.Context) error {
		credential := &mgo.Credential{Username: mongoUsername, Password: mongoPassword}
		session, _ := mgo.Dial("mongodb")
		session.Login(credential)
		words := session.DB("sssignal3").C("endless_typing_words")

		var wordList []EndlessTypingWord
		words.Find(bson.M{}).All(&wordList)

		result := []string{}
		for _, value := range wordList {
			result = append(result, value.Word)
		}

		session.Close()
		return c.JSON(http.StatusOK, result)
	})
	// /soft/endless_typing 用のランキングを取得
	e.GET("/soft/endless_typing/ranking", endlessTyping.GetRanking)
	// /soft/endless_typing 用のランキングを更新
	e.PUT("/soft/endless_typing/ranking", endlessTyping.PutRanking)
	e.Logger.Fatal(e.Start(":1323"))
}
