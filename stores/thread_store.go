package stores

import (
	"errors"

	"github.com/YO-RO/kgb/models"
)

type threadsStore []models.Thread

var ThreadStore threadsStore = make(threadsStore, 0, 30)

func (tp *threadsStore) Create(target models.Thread) {
	target.Id = len(*tp) + 1
	*tp = append(*tp, target)
}

func (threads threadsStore) Read() []models.Thread {
	copied := make(threadsStore, len(threads))
	copy(copied, threads)
	return copied
}

func (threads threadsStore) FindById(id int) (models.Thread, error) {
	index := id - 1
	if !(index >= 0 && index < len(threads)) {
		return models.Thread{}, errors.New("id is invalid")
	}

	return threads[index], nil
}

func (t threadsStore) Update(updatedThread models.Thread) error {
	index := updatedThread.Id - 1
	if !(index >= 0 && index < len(t)) {
		return errors.New("id is invalid")
	}

	t[index] = updatedThread
	return nil
}
