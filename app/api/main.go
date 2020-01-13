package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/joho/godotenv"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"os"
	"fmt"
)

type Slack struct {
	Text string `json:"text"`
}

func main() {
	e := echo.New()

	// .envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env loading error!")
	}

	slackWebhook := os.Getenv("SLACK_WEBHOOK")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!!!")
	})
	e.GET("/inquiry", func(c echo.Context) error {
		// パラメータを受け取る
		address := c.FormValue("address")
		text := c.FormValue("text")

		// TODO : 見た目もう少し良くしたい 
		// https://api.slack.com/tools/block-kit-builder
		slack := new(Slack)
		slack.Text = "*【問い合わせフォーム】*\naddress : *" + address + "*\ntext : " + text

		// 構造体をJSONに変換
		json, _:= json.Marshal(slack)

		// JSONちゃんと作れてるかの確認用
		// return c.JSON(http.StatusOK, url.Values{"payload": {string(json)}})

		// Slack webhookにPOSTリクエスト送る
		response, err := http.PostForm(slackWebhook, url.Values{"payload": {string(json)}})
		// レスポンス読み込み
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return c.String(http.StatusOK, "err!")
		}

		// webhookの結果をJSONで返す
		return c.JSON(http.StatusOK, map[string]string{"result": string(body)})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
