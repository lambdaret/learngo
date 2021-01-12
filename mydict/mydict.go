package mydict

import "errors"

// Dictionary ...
type Dictionary map[string] string


var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-exising word")
	errWordExists = errors.New("That word already exists")
)

// Search ...
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add (word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

// Update ...
func (d Dictionary) Update (word, definition string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete ...
func (d Dictionary) Delete(word string) {
	delete(d, word)
}