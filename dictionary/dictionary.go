package dictionary

import "errors"

var ErrNotFound = errors.New("could not find word in dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
