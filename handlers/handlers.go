package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/YO-RO/kgb/models"
	"github.com/YO-RO/kgb/stores"
)

func ThreadsViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	formHtml := "<form action='/post' method='post'>" +
		"<label>ニックネーム<br><input type='text' name='name'/></label><br>" +
		"<textarea name='body' placeholder='今日の出来事を書いてみましょう' rows='4' cols='40'></textarea><br>" +
		"<button>投稿</button>" +
		"</form>"

	makePresentThreadHtml := func(t models.Thread) string {
		return fmt.Sprintf(
			"<div>"+
				"(%d) %s --- %s<p style='white-space: pre-wrap'>%s</p>"+
				"<form action='/delete' method='post'>"+
				"<button name='id' value='%d'>勝手に削除する🤪</button>"+
				"</form></div>",
			t.Id,
			html.EscapeString(t.Name),
			t.CreatedAt.Format("2006/1/2 15:04"),
			html.EscapeString(t.Body),
			t.Id,
		)
	}

	makeDeletedThreadHtml := func(t models.Thread) string {
		return fmt.Sprintf(
			"<div>(%d) %s --- %s<p>%sに誰かに削除されました。</p></div>",
			t.Id, t.Name, t.CreatedAt.Format("2006/1/2 15:04"), t.DeletedAt.Format("2006/1/2 15:04"),
		)
	}

	var threadsHtml string
	threads := stores.ThreadStore.Read()
	for _, t := range threads {
		switch t.IsDeleted {
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

func PostThreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	newThread := models.NewThread(
		r.PostFormValue("name"),
		r.PostFormValue("body"),
		time.Now(),
	)
	stores.ThreadStore.Create(newThread)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteThreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	target, err := stores.ThreadStore.FindById(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	target.Delete(time.Now())

	if err := stores.ThreadStore.Update(target); err != nil {
		log.Printf("delete failed - id: %d\n", id)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
