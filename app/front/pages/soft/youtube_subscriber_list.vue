<template>
    <div v-if="access_token === 'none'">
        <p>
            youtube subscriber list
        </p>
        <p>
            このページでは、Youtubeの登録チャンネルを管理することができます<br>
            まずはYoutubeにログインしてください！
        </p>
        <h1>
            <a :href="'https://accounts.google.com/o/oauth2/auth?client_id=260053394039-9pgbce0boctusgg974ai62bsqls5ituq.apps.googleusercontent.com&redirect_uri=' + location_protocol + '//' + location_host + '/soft/youtube_subscriber_list&scope=https://www.googleapis.com/auth/youtube.readonly&response_type=token'">
                Youtubeログイン
            </a>
        </h1>
    </div>
    <div v-else>
        <!-- 左側 -->
        <div class="left-menu">
            <div style="padding: 10px 0px" v-bind:class="{ 'list-select': isListSelect === 0 }" v-on:click="setListSelect(0)">
                登録チャンネル
            </div>
            <div v-for="(value, index) in list" :key="'category-' + index" style="padding: 10px 0px" v-bind:class="{ 'list-select': isListSelect === 1 + index }" v-on:click="setListSelect(1 + index)">
                <div v-if="index === isListEdit">
                    <input type="text" style="width: 80%;" v-model="editCategory" @keydown.enter="setCategory">
                </div>
                <div v-else>
                    <div style="display: inline-block;">{{ value.name }}</div>
                    <div v-on:click="deleteList(index)" class="list-action">X</div>
                    <div v-on:click="setListEdit(index)" class="list-action">E</div>
                </div>
            </div>
            <div v-on:click="setNewListEdit()" style="padding: 10px 0px">
                +
            </div>
        </div>
        <!-- 右側 -->
        <div class="right-menu">
            <div v-show="isListSelect === 0" style="width: 100%;">
                <div class="right-header" style="padding: 4px 10px;">クリックでチャンネルページに飛びます</div>
                <div class="right-content" style="display: flex; flex-wrap: wrap;">
                    <div v-for="channel in subscribe_channel" :key="channel.id" class="listed-channel">
                        <a target="_blank" :href="'https://www.youtube.com/channel/' + channel.snippet.resourceId.channelId" style="text-decoration: none;">
                            <img :src="channel.snippet.thumbnails.default.url">
                            <div class="edit-channel-title">{{ channel.snippet.title }}</div>
                        </a>
                    </div>
                </div>
            </div>
            <div v-show="isListSelect > 0" style="width: 100%;">
                <div class="right-header">
                    <div style="padding: 5px 10px" v-on:click="setDetailAction('edit_list')" v-bind:class="{'action-select': listDetailAction === 'edit_list'}">edit_list</div>
                    <div style="padding: 5px 10px" v-on:click="setDetailAction('edit_detail')" v-bind:class="{'action-select': listDetailAction === 'edit_detail'}">edit_detail</div>
                    <div style="padding: 5px 10px" v-on:click="setDetailAction('show_list')" v-bind:class="{'action-select': listDetailAction === 'show_list'}">show_list</div>
                </div>
                <div v-show="listDetailAction === 'edit_list'" class="right-content" style="display: flex; flex-wrap: wrap;">
                    <div v-for="channel in subscribe_channel" :key="channel.id" class="listed-channel" v-bind:class="{'registed-channel': list[isListSelect - 1] === undefined ? false : list[isListSelect - 1].channelIds.includes(channel.snippet.resourceId.channelId)}">
                        <div style="cursor: pointer;" v-on:click="toggleListChannel(channel.snippet.resourceId.channelId)">
                            <img :src="channel.snippet.thumbnails.default.url">
                            <div class="edit-channel-title">{{ channel.snippet.title }}</div>
                        </div>
                    </div>
                </div>
                <div v-show="listDetailAction === 'edit_detail'" class="right-content" style="display: flex; flex-wrap: wrap;">
                    <div v-for="channel in subscribe_channel" :key="channel.id">
                        <div v-show="list[isListSelect - 1] === undefined ? false : list[isListSelect - 1].channelIds.includes(channel.snippet.resourceId.channelId)" class="listed-channel">
                            <img :src="channel.snippet.thumbnails.default.url">
                            <div class="edit-channel-title">{{ channel.snippet.title }}</div>
                        </div>
                    </div>
                </div>
                <div v-show="listDetailAction === 'show_list'" class="right-content">
                    <div v-for="video in subscribe_videoList" :key="video.etag">
                        <a v-show="list[isListSelect - 1] === undefined ? false : list[isListSelect - 1].channelIds.includes(video.snippet.channelId)" target="_blank" :href="'https://www.youtube.com/watch?v=' + video.contentDetails.videoId" style="display: flex; margin: 2px; padding: 3px; text-decoration: none; font-size: 14px; max-width: 100%; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
                            <img :src="video.snippet.thumbnails.default.url">
                            <div style="display: flex; flex-direction: column; justify-content: center; width: calc(100% - 120px); margin: 5px;">
                                <div class="show-text" style="font-size: 16px;">{{ video.snippet.title }}</div>
                                <div class="show-text" style="font-size: 12px;">{{ video.snippet.channelTitle }}</div>
                                <div class="show-text" style="font-size: 12px;">{{ new Date(video.contentDetails.videoPublishedAt).toLocaleString() }}</div>
                                <div class="show-text" style="font-size: 12px;">{{ video.videoDetail !== undefined ? video.videoDetail.contentDetails.duration : '' }}</div>
                            </div>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return ({
            access_token: '',
            token_type: '',
            // 全ての登録チャンネル
            subscribe_channel: [],
            // 登録チャンネルのプレイリストID
            subscribe_playlistIds: [],
            // 登録チャンネルの動画一覧
            subscribe_videoList: [],
            // 作成したリスト一覧
            list: [],
            // ログインしたユーザのチャンネルID
            my_channelId: '',
            isListSelect: 0,
            isListEdit: -1,
            listDetailAction: 'edit_list',
            // 編集中のリスト名
            editCategory: '',
            location_protocol: '',
            location_host: ''
        });
    },
    async mounted() {
        // リロードなどの時にbeforeDestroyを起こす
        window.addEventListener('beforeunload', this.setMyListCookie);

        this.$store.commit("character/setShow", false);
        this.location_protocol = location.protocol;
        this.location_host = location.host;

        // URL末尾のハッシュを分解
        const urlHash = this.$route.hash.substr(1);
        const urlHashParams = urlHash.split('&').reduce(function (res, item) {
            const parts = item.split('=');
            res[parts[0]] = parts[1];
            return res;
        }, {});

        // cookieを分解
        const cookieParams = document.cookie.split(';').reduce(function (res, item) {
            const parts = item.split('=');
            res[parts[0].trim()] = parts[1];
            return res;
        }, {});

        if (urlHashParams.access_token !== undefined) {
            /* URLにaccess_tokenがあるとき, cookieに入れて処理続行 */
            // googleからリダイレクトされた返り値を受け取る
            document.cookie = `access_token=${urlHashParams.access_token}; max-age=${60*60*24*365}`;
            document.cookie = `scope=${urlHashParams.scope}; max-age=${60*60*24*365}`;
            document.cookie = `token_type=${urlHashParams.token_type}; max-age=${60*60*24*365}`;
            this.access_token = urlHashParams.access_token;
            this.token_type = urlHashParams.token_type;
        } else if (cookieParams.access_token !== undefined) {
            /* cookieある時, 処理続行 */
            this.access_token = cookieParams.access_token;
            this.token_type = cookieParams.token_type;
        } else {
            /* 何も無い時は, googleログインのリンクを出して終了 */
            this.access_token = 'none';
            return;
        }

        // https://developers.google.com/youtube/v3/guides/auth/client-side-web-apps
        // 4. ユーザのトークンを検証する
        await fetch(`https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=${this.access_token}`, {
            "method": "GET"
        })
        .then(response => response.json())
        .then(data => {
            if (data.error !== undefined) {
                // セッション有効期限切れなどのエラー
                this.access_token = 'none';
                return;
            }
        })
        .catch(err => {
            console.log(err);
        });

        // 自分のチャンネルIDを取得
        this.my_channelId = await this.getMyChannelId();
        if (cookieParams[`list_${this.my_channelId}`] !== undefined) {
            this.list = JSON.parse(cookieParams[`list_${this.my_channelId}`]);
        }

        // 登録チャンネル一覧取得
        let data = await this.setSubscribeChannel();

        // 登録チャンネル一覧取得
        while (data.nextPageToken !== undefined) {
            data = await this.setSubscribeChannel(data.nextPageToken);
        }
        
        // 登録チャンネルのプレイリストID一覧取得
        const setPlaylistFunctions = [];
        this.subscribe_channel.forEach(channel => {
            setPlaylistFunctions.push(this.setPlaylistId(channel.snippet.resourceId.channelId));
        })
        await Promise.all(setPlaylistFunctions);

        // プレイリストIDから動画一覧取得
        const setVideolistFunctions = [];
        this.subscribe_playlistIds.forEach(playlistId => {
            setVideolistFunctions.push(this.setVideoList(playlistId));
        });
        await Promise.all(setVideolistFunctions);
        
        // 動画の詳細情報を取得
        const setVideoDetailFunctions = [];
        let videoIdsString = '';
        this.subscribe_videoList.forEach((video, index) => {
            if (index % 50 === 0) {
                if (videoIdsString !== '') {
                    setVideoDetailFunctions.push(this.setVideoDetail(videoIdsString));
                }
                videoIdsString = video.snippet.resourceId.videoId;
            } else {
                videoIdsString += `,${video.snippet.resourceId.videoId}`;
            }
        });
        if (videoIdsString !== '') {
            setVideoDetailFunctions.push(this.setVideoDetail(videoIdsString));
        }
        await Promise.all(setVideoDetailFunctions);

        // 動画を公開日時順に並び替え
        this.subscribe_videoList = this.subscribe_videoList.sort((prev, next) => {
            const prevAt = Date.parse(prev.contentDetails.videoPublishedAt);
            const nextAt = Date.parse(next.contentDetails.videoPublishedAt);
            if (prevAt > nextAt) {
                return -1;
            } else {
                return 1;
            }
        });
    },
    beforeDestroy() {
        this.setMyListCookie();
        window.removeEventListener('beforeunload', this.setMyListCookie);
    },
    methods: {
        getMyChannelId() {
            return fetch('https://www.googleapis.com/youtube/v3/channels?part=id&mine=true', {
                'method': 'GET',
                'headers': {
                    'Authorization': `${this.token_type} ${this.access_token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                if (data.error !== undefined) {
                    // セッション有効期限切れなどのエラー
                    this.access_token = 'none';
                    return data;
                }

                return data.items[0].id;
            });
        },
        setSubscribeChannel(nextPageToken) {
            return fetch(`https://www.googleapis.com/youtube/v3/subscriptions?part=snippet&mine=true&maxResults=50&order=alphabetical${nextPageToken !== undefined ? `&pageToken=${nextPageToken}` : ''}`, {
                'method': 'GET',
                'headers': {
                    'Authorization': `${this.token_type} ${this.access_token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                if (data.error !== undefined) {
                    // セッション有効期限切れなどのエラー
                    this.access_token = 'none';
                    return data;
                }
                
                // チャンネル一覧に追加
                data.items.forEach(item => {
                    this.subscribe_channel.push(item);
                });

                return data;
            });
        },
        setPlaylistId(channelId) {
            return fetch(`https://www.googleapis.com/youtube/v3/channels?part=contentDetails&id=${channelId}&maxResults=1`, {
                'method': 'GET',
                'headers': {
                    'Authorization': `${this.token_type} ${this.access_token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                if (data.error !== undefined) {
                    // セッション有効期限切れなどのエラー
                    this.access_token = 'none';
                    return;
                }
                
                const playlistId = data.items[0].contentDetails.relatedPlaylists.uploads;
                this.subscribe_playlistIds.push(playlistId);
            })
        },
        setVideoList(playlistId) {
            return fetch(`https://www.googleapis.com/youtube/v3/playlistItems?part=snippet,contentDetails,status&playlistId=${playlistId}&maxResults=3`, {
                'method': 'GET',
                'headers': {
                    'Authorization': `${this.token_type} ${this.access_token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                if (data.error !== undefined) {
                    // セッション有効期限切れなどのエラー
                    // this.access_token = 'none';
                    return;
                }

                data.items.forEach(item => {
                    this.subscribe_videoList.push(item);
                });
            })
        },
        setVideoDetail(videoIdsString) {
            return fetch(`https://www.googleapis.com/youtube/v3/videos?part=id,contentDetails&id=${videoIdsString}&maxResults=50`, {
                'method': 'GET',
                'headers': {
                    'Authorization': `${this.token_type} ${this.access_token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                if (data.error !== undefined) {
                    // セッション有効期限切れなどのエラー
                    // this.access_token = 'none';
                    return;
                }

                data.items.forEach(video => {
                    this.subscribe_videoList.forEach((value, index) => {
                        if (value.snippet.resourceId.videoId === video.id) {
                            this.subscribe_videoList[index]['videoDetail'] = video;
                        }
                    });
                });

                return data;
            });
        },
        setNewListEdit() {
            if (this.isListEdit < 0) {
                this.list.push({ name: '', channelIds: [] });
                this.isListEdit = this.list.length - 1;
                this.isListSelect = this.isListEdit + 1;
            }
        },
        setListEdit(index) {
            this.isListEdit = index;
            this.editCategory = this.list[index].name;
        },
        deleteList(index) {
            this.list = this.list.filter((value, i) => index !== i);
            this.setListSelect(index);
        },
        setListSelect(index) {
            if (this.isListSelect !== index) {
                this.isListEdit = -1;
                this.isListSelect = index;
                if (this.isListSelect > this.list.length) {
                    this.isListSelect = this.list.length;
                }
            }
        },
        setCategory(event) {
            if (event.keyCode === 13) {
                this.list[this.isListEdit].name = this.editCategory;
                this.isListEdit = -1;
                this.editCategory = '';
            }
            setMyListCookie();
        },
        setDetailAction(action) {
            this.listDetailAction = action;
        },
        toggleListChannel(channelId) {
            if (this.list[this.isListSelect - 1].channelIds.includes(channelId)) {
                this.list[this.isListSelect - 1].channelIds = this.list[this.isListSelect - 1].channelIds.filter(id => id !== channelId);
            } else {
                this.list[this.isListSelect - 1].channelIds.push(channelId);
            }
            setMyListCookie();
        },
        setMyListCookie() {
            document.cookie = `list_${this.my_channelId}=${JSON.stringify(this.list)}; max-age=${60*60*24*365}`;
        }
    }
}
</script>

<style lang="scss" scoped>
.list-select {
    background: #446;
}

.list-action {
    display: inline-block;
    float: right;
    margin-right: 5px;
    cursor: pointer;
}

.left-menu {
    position: fixed; 
    width: 150px; 
    background: #222;
}

.right-menu {
    display: flex; 
    flex-wrap: wrap; 
    width: calc(100% - 150px); 
    margin-left: 150px; 
    background: #444;
}

.right-header {
    display: flex; 
    justify-content: center; 
    position: fixed; 
    width: calc(100% - 150px); 
    background: rgba(64, 64, 64, 0.9);
}

.right-content {
    margin-top: 32px;
}

@media screen and (max-width: 500px) {
  .left-menu {
    width: 100%;
    height: 100px;
    overflow: scroll;
    opacity: 0.95;
  }

  .right-menu {
      width: 100%;
      margin-left: 0px;
  }

  .right-header {
      width: 100%;
      margin-top: 100px;
  }

  .right-content {
      margin-top: 132px;
  }
}

.action-select {
    color: #44F;
}

.registed-channel {
    background: #88F;
}

.listed-channel {
    width: 100px; 
    margin: 1px; 
    padding: 3px; 
    border-radius: 2px
}

.edit-channel-title {
    font-size: 12px; 
    max-width: 100%; 
    overflow: hidden; 
    text-overflow: ellipsis; 
    white-space: nowrap;
}

.show-text {
    width: 100%; 
    overflow: hidden; 
    text-overflow: ellipsis; 
    white-space: nowrap; 
    text-align: left; 
    margin: 3px;
}
</style>