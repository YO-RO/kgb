package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"time"
)

const port = ":8888"

type thread struct {
	id        int
	name      string
	body      string
	createdAt time.Time
	deletedAt time.Time
	isDeleted bool
}

func newThread(name string, body string) thread {
	if name == "" {
		name = "匿名さん"
	}
	if body == "" {
		body = "本文なし...\n何か書いてよ〜"
	}
	return thread{name: name, body: body, createdAt: time.Now()}
}

type threadsStore []thread

var store threadsStore = make(threadsStore, 0, 30)

func (tp *threadsStore) post(target thread) {
	target.id = len(*tp) + 1
	*tp = append(*tp, target)
}

func (threads threadsStore) read() []thread {
	for i, t := range threads {
		if t.isDeleted {
			threads[i].body = "削除されました"
		}
	}
	return threads
}

func (t threadsStore) delete(id int) bool {
	index := id - 1
	if !(index >= 0 && index < len(t)) {
		return false
	}

	target := t[index]
	target.deletedAt = time.Now()
	target.isDeleted = true
	t[index] = target
	return true
}

func handleThreadView(w http.ResponseWriter, r *http.Request) {
	formHtml := "<form action='/post' method='post'>" +
		"<label>ニックネーム<br><input type='text' name='name'/></label><br>" +
		"<textarea name='body' placeholder='今日の出来事を書いてみましょう' rows='4' cols='40'></textarea><br>" +
		"<button>投稿</button>" +
		"</form>"

	makePresentThreadHtml := func(t thread) string {
		return fmt.Sprintf(
			"<div>"+
				"(%d) %s --- %s<p>%s</p>"+
				"<form action='/delete' method='post'>"+
				"<button name='id' value='%d'>勝手に削除する🤪</button>"+
				"</form></div>",
			t.id,
			html.EscapeString(t.name),
			t.createdAt.Format("2006/1/2 15:04"),
			html.EscapeString(t.body),
			t.id,
		)
	}

	makeDeletedThreadHtml := func(t thread) string {
		return fmt.Sprintf(
			"<div>(%d) %s --- %s<p>%sに誰かに削除されました。</p></div>",
			t.id, t.name, t.createdAt.Format("2006/1/2 15:04"), t.deletedAt.Format("2006/1/2 15:04"),
		)
	}

	threads := store.read()
	var threadsHtml string
	for _, t := range threads {
		switch t.isDeleted {
		case true:
			threadsHtml += makeDeletedThreadHtml(t)
		case false:
			threadsHtml += makePresentThreadHtml(t)
		}
	}
	threadsHtml = "<div>" + threadsHtml + "</div>"

	pageHtml := "<html><h1>匿名掲示板</h1>" + formHtml + threadsHtml + "</html>"
	w.Write([]byte(pageHtml))
}

func handleThreadPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	store.post(newThread(r.PostFormValue("name"), r.PostFormValue(("body"))))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func handleThreadDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	if ok := store.delete(id); !ok {
		log.Printf("delete failed - id: %d\n", id)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	log.Printf("Start Server - http://localhost%s\n", port)

	http.HandleFunc("/", handleThreadView)
	http.HandleFunc("/post", handleThreadPost)
	http.HandleFunc("/delete", handleThreadDelete)
	log.Fatal(http.ListenAndServe(port, nil))
}
