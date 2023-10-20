package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
