package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWoedExist = errors.New("That word already exist")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	// if - else로 표현했을 때
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWoedExist
	// }
	// return nil

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWoedExist
	}
	return nil
}
