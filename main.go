package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var bot *tb.Bot
var groupId int64
var dryrun bool

func main() {
	botToken := os.Getenv("TG_TOKEN")
	envGroupId := os.Getenv("TG_GROUP_ID")
	dbg := os.Getenv("BOT_DEBUG")

	log.SetOutput(os.Stdout)
	rand.Seed(time.Now().Unix())

	if dbg != "" {
		log.Debug("debug logging enabled")
		log.SetLevel(log.DebugLevel)
	}

	if botToken != "" {
		var err error
		bot, err = tb.NewBot(tb.Settings{
			Token:  botToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Panicln(err)
		}
	} else {
		log.SetLevel(log.DebugLevel)
		log.Debugln("no token provided. running in dry-run mode and enabling debug mode.")
		dryrun = true
	}

	if envGroupId == "" && !dryrun {
		log.Info("no group id provided, will wait to be added to one")
		bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
			groupId = m.Chat.ID
			log.Printf("added to: %v", groupId)
		})
	} else if envGroupId != "" && !dryrun {
		var err error
		groupId, err = strconv.ParseInt(envGroupId, 10, 64)
		if err != nil {
			log.Fatalf("could not parse provided group id: %v", err)
		}
		log.Infof("sending to group: %v", groupId)
	} else {
		log.Debugln("would be handling OnAddedToGroup event")
	}

	if !dryrun {
		// it's alive...
		go bot.Start()
	} else {
		log.Debugln("would have started bot")
	}
	// ... and listening for events
	router := mux.NewRouter()
	router.HandleFunc("/notify/build", NotifyBuildHandler)
	router.HandleFunc("/notify/pr", PullRequestHandler)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Debugln("api online")
	log.Fatal(http.ListenAndServe(":8888", router))
}
