package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)

}

const (
	NotFoundWordErr  = DictionaryErr("could not find the word")
	WordExistsErr    = DictionaryErr("word already exists")
	DoestNotExistErr = DictionaryErr("canont update word because it doest no exist")
)

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]

	if !ok {
		return "", NotFoundWordErr
	}

	return word, nil

}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundWordErr:
		d[word] = definition
	case nil:
		return WordExistsErr
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundWordErr:
		return DoestNotExistErr
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil

}

func (d Dictionary) Delete(word string) {
	delete(d, word)

}
