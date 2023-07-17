package models

import "time"

type Thread struct {
	Id        int
	Name      string
	Body      string
	CreatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func NewThread(name string, body string) Thread {
	if name == "" {
		name = "匿名さん"
	}
	if body == "" {
		body = "本文なし...\n何か書いてよ〜"
	}
	return Thread{
		Name:      name,
		Body:      body,
		CreatedAt: time.Time{},
	}
}
