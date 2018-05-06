package main

import (
	"fmt"
	"math/rand"
)

const (
	OK      = "SUCCESS"
	FAIL    = "FAILURE"
	ABORTED = "ABORTED"
)

const (
	STARTED  = "started"
	FINISHED = "finished"
	WAITING  = "waiting"
)

type EmojiProvider struct{}

func (EmojiProvider) giveOne(emojiList []string) string {
	return emojiList[rand.Intn(len(emojiList))]
}
func (em EmojiProvider) Start() string {
	emojis := []string{"🏁", "🤖", "👐", "🤞", "👩‍💻", "🦊", "🐉", "💫", "🚀"}
	return em.giveOne(emojis)
}
func (em EmojiProvider) Ok() string {
	emojis := []string{"🦄", "🐳", "🤘", "🙌", "🐙", "✨", "⭐️", "💯", "🔝", "🍻", "🎉"}
	return em.giveOne(emojis)
}

func (em EmojiProvider) Fail() string {
	emojis := []string{"🔥", "🥀", "💥", "⚡️", "⛈", "🥐", "🥃", "🚒", "🛸", "⚰️", "💔", "❌"}
	return em.giveOne(emojis)
}

func (em EmojiProvider) Aborted() string {
	emojis := []string{"👼", "❓", "🚮", "😔", "🤨", "🤔"}
	return em.giveOne(emojis)
}

func (em EmojiProvider) Waiting() string {
	emojis := []string{"👆", "🖕"}
	return em.giveOne(emojis)
}

func (em EmojiProvider) PullRequest() string {
	emojis := []string{"😱", "🧐", "🤓", "☠️", "🤝", "💡", "🍴"}
	return em.giveOne(emojis)
}

func (em EmojiProvider) Unknown() string {
	emojis := []string{"💩"}
	return em.giveOne(emojis)
}

type BuildNotification struct {
	Phase   string `json:"phase"`
	Result  string `json:"result"`
	Project string `json:"project"`
	FullUrl string `json:"build_url"`
}

func (bn BuildNotification) ToText() string {
	var prefix string
	switch bn.Phase {
	case STARTED:
		prefix = EmojiProvider{}.Start()
		return fmt.Sprintf("%s Inició el build de **%s** [%s]", prefix, bn.Project, bn.FullUrl)
	case FINISHED:
		switch bn.Result {
		case OK:
			prefix = EmojiProvider{}.Ok()
		case FAIL:
			prefix = EmojiProvider{}.Fail()
		case ABORTED:
			prefix = EmojiProvider{}.Aborted()
		}
		return fmt.Sprintf("%s Terminó el build de **%s** con resultado: %s [%s]", prefix, bn.Project, bn.Result, bn.FullUrl)
	case WAITING:
		prefix = EmojiProvider{}.Waiting()
		return fmt.Sprintf("%s El build de **%s** está esperando input! [%s]", prefix, bn.Project, bn.FullUrl)
	default:
		prefix = EmojiProvider{}.Unknown()
		return fmt.Sprintf("%s Evento desconocido: %+v", prefix, bn)
	}
}

type PullRequestNotification struct {
	Project   string
	Target    string
	ChangeId  string
	Author    string
	ChangeUrl string
}

func (prn PullRequestNotification) ToText() string {
	prefix := EmojiProvider{}.PullRequest()
	return fmt.Sprintf("%s Nuevo PullRequest de %s para el branch **%s** de *%s* (PR #%s) [%s]",
		prefix, prn.Author, prn.Target, prn.Project, prn.ChangeId, prn.ChangeUrl)
}

type Pronounceable interface {
	ToText() string
}
