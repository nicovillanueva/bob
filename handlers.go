package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
)

func reply(w http.ResponseWriter, code int, message string) {
	response, err := json.Marshal(message)
	if err != nil {
		log.Errorf("could not marshal message: %v", err)
		response = []byte{}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func BotAnnounce(message Pronounceable) {
	text := message.ToText()
	if !dryrun {
		_, err := bot.Send(&tb.Chat{
			ID: groupId,
		}, text, tb.ModeMarkdown)
		if err != nil {
			log.Errorf("error dispatching event: %v", err)
		} else {
			log.Debug("dispatched ok")
		}
	} else {
		log.Debugf("would have said: \"%s\"", text)
	}
}

func NotifyBuildHandler(w http.ResponseWriter, r *http.Request) {
	var bn BuildNotification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bn); err != nil {
		log.Errorf("error decoding: %v", err)
		reply(w, 400, "invalid payload")
		return
	}
	defer r.Body.Close()

	log.Debugf("Got build: %+v", bn)
	BotAnnounce(bn)
}

func PullRequestHandler(w http.ResponseWriter, r *http.Request) {
	var prn PullRequestNotification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&prn); err != nil {
		log.Errorf("error decoding: %v", err)
		reply(w, 400, "invalid payload")
		return
	}
	defer r.Body.Close()
	log.Debugf("Got PR: %+v", prn)
	BotAnnounce(prn)
}
