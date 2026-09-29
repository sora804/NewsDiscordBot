package main

import (
	"DiscordBot/discord"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID string
	dgs     *discordgo.Session
)

// discordボットのエントリーポイント
func main() {
	sessionManager := &discord.DiscordSessionManager{}
	dgs = sessionManager.InitializeSession(os.Getenv("DISCORD_BOT_TOKEN"))

	// Botに権限を付与している
	dgs.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsAll | discordgo.PermissionSendMessages

	if err := dgs.Open(); err != nil {
		log.Fatalf("Discordセッションのオープンに失敗: %v", err)
	}
	defer dgs.Close()

	log.Println("ボットが起動しました。Ctrl+Cで終了します。")

	waitForExitSignal()
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
