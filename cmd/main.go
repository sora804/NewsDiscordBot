package main

import (
	"fmt"
	"log"
	"newsSite/api"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	GuildID string
	dgs     *discordgo.Session
)

/*
・Discordボットの起動
・自動物理削除（48時間以前削除）
・コマンドなしで、自動配信

*/

// discordボットのエントリーポイント
func main() {
	//環境変数
	if err := godotenv.Load(); err != nil {
		log.Println(".env を読み込めませんでした")
		//エラーのため終了
		return
	}
	token := os.Getenv("DISCORD_TOKEN")

	// DB初期化と情報登録の処理
	err := api.DBinit()
	if err != nil {
		log.Fatal(err)
	}
	api.RegisterRSS_Handler() //これは時間駆動にする
	// ここでエラーハンドリングを追加
	if err != nil {
		fmt.Println(err)
	}

	/*------------------------------------------------
	----bot起動処理------------------------------------
	------------------------------------------------*/

	//token処理
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	//discord.AddHandler(api.GetData) //メッセージ受信時に呼ぶFunc

	err = discord.Open()

	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	err = discord.Close()

	//return

	/*------------------------------------------------------
	----time処理------------------------------------
	------------------------------------------------*/
	//now := time.Now()
	//log.Println("Current time:", now)

	log.SetFlags(log.Lmicroseconds)
	ticker := time.NewTicker(time.Minute * 10)
	defer ticker.Stop()
	count := 0
	for {
		select {
		case <-ticker.C:
			log.Printf("count=%d\n", count)
			count++
			api.GetRSS_Handler() //
		case <-stopBot:
			log.Print("stop")
			return
		}
	}

}

// waitForExitSignalは終了シグナルを待機してボットを安全にシャットダウンする
func waitForExitSignal() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

/*
webサイト用main関数

import (
	"net/http"
	// "newsSite/api"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//処理
			//この辺でdbのfuncを呼び出したり
		default:
			//処理
		}
	})
	http.HandleFunc("/ss/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//処理
		default:
			//処理
		}
	})
	http.ListenAndServe(":8080", nil)
}
*/
