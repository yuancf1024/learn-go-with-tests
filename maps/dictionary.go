package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("counld not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)


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