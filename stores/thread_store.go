package stores

import (
	"time"

	"github.com/YO-RO/kgb/models"
)

type threadsStore []models.Thread

var ThreadStore threadsStore = make(threadsStore, 0, 30)

func (tp *threadsStore) Post(target models.Thread) {
	target.Id = len(*tp) + 1
	*tp = append(*tp, target)
}

func (threads threadsStore) Read() []models.Thread {
	for i, t := range threads {
		if t.IsDeleted {
			threads[i].Body = "削除されました"
		}
	}
	return threads
}

func (t threadsStore) Delete(id int) bool {
	index := id - 1
	if !(index >= 0 && index < len(t)) {
		return false
	}

	target := t[index]
	target.DeletedAt = time.Now()
	target.IsDeleted = true
	t[index] = target
	return true
}
