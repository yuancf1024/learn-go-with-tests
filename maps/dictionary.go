package main

// import "errors"

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("counld not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// 定义字典的搜索函数
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok { // ok = false 没有成功找到key
		return "", ErrNotFound
	}
	return definition, nil
}

// 定义字典的添加函数功能
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}