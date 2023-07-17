package models

import (
	"strings"
	"time"
	"unicode"
)

type Thread struct {
	Id        int
	Name      string
	Body      string
	CreatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func NewThread(name string, body string, now time.Time) Thread {
	name = strings.TrimSpace(name)
	body = strings.TrimRightFunc(body, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	if name == "" {
		name = "匿名さん"
	}
	if body == "" {
		body = "本文なし...\n何か書いてよ〜"
	}
	return Thread{
		Name:      name,
		Body:      body,
		CreatedAt: now,
	}
}

func (t *Thread) Delete(now time.Time) {
	t.DeletedAt = now
	t.IsDeleted = true
}
