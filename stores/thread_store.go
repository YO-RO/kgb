package stores

import (
	"errors"

	"github.com/YO-RO/kgb/models"
)

type threadsStore []models.Thread

var ThreadStore threadsStore = make(threadsStore, 0, 30)

func (sp *threadsStore) Create(target models.Thread) {
	target.Id = len(*sp) + 1
	*sp = append(*sp, target)
}

func (s threadsStore) Read() []models.Thread {
	copied := make(threadsStore, len(s))
	copy(copied, s)
	return copied
}

func (s threadsStore) FindById(id int) (models.Thread, error) {
	index := id - 1
	if !(index >= 0 && index < len(s)) {
		return models.Thread{}, errors.New("id is invalid")
	}

	return s[index], nil
}

func (s threadsStore) Update(updatedThread models.Thread) error {
	index := updatedThread.Id - 1
	if !(index >= 0 && index < len(s)) {
		return errors.New("id is invalid")
	}

	s[index] = updatedThread
	return nil
}
