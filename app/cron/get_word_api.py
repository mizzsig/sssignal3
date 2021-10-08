### /soft/endless_typing に使用する英単語を取得し、DBに保存する
import os
import requests
from pymongo import MongoClient
from dotenv import load_dotenv

### mongoDBに接続
load_dotenv()
mongoDB = MongoClient('mongodb')
mongoDB['admin'].authenticate(os.getenv('MONGO_USERNAME'), os.getenv('MONGO_PASSWORD'))

### APIからランダムで単語を取得
url = "https://wordsapiv1.p.rapidapi.com/words/"

querystring = {"random":"true"}

headers = {
    'x-rapidapi-key': os.getenv('RAPIDAPI_KEY'),
    'x-rapidapi-host': "wordsapiv1.p.rapidapi.com"
}

response = requests.request("GET", url, headers=headers, params=querystring)
responseWord = response.json()['word']

### APIから取得したものをDBに保存(スタイル崩れの原因になる文字は抜く)
if ' ' not in responseWord and '-' not in responseWord:
    data = mongoDB['sssignal3']['endless_typing_words'].find_one({'word': responseWord})
    if data is None:
        mongoDB['sssignal3']['endless_typing_words'].insert_one({'word': responseWord})
