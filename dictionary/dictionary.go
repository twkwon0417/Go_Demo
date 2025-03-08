package dictionary

import (
	"errors"
)

type Dictionary map[string]string

var (
	errNoMatchingWord = errors.New("no matching word")
	errAlreadyExist   = errors.New("word already exists")
	errCantUpdate     = errors.New("can't update")
)

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if ok {
		return def, nil
	}
	return "", errNoMatchingWord
}

func (d Dictionary) Add(word, def string) error {
	if _, err := d.Search(word); err == nil {
		return errAlreadyExist
	}
	d[word] = def
	return nil
}

func (d Dictionary) Update(word, def string) error {
	switch _, err := d.Search(word); {
	case err == nil:
		d[word] = def
	case errors.Is(err, errNoMatchingWord):
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
