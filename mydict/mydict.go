package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not found")
	errWoedExist  = errors.New("that word already exist")
	errCantUpdate = errors.New("can't update non-existing word")
)

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

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
