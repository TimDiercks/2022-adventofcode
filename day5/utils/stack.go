package utils

import (
	"errors"
)

type Stack []byte

func (s *Stack) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (s *Stack) Push(item byte) {
	*s = append(*s, item)
}

func (s *Stack) Top() (byte, error) {
	if s.IsEmpty() {
		return 0, errors.New("Empty Stack")
	}
	topIndex := len(*s) - 1
	element := (*s)[topIndex]
	return element, nil
}

func (s *Stack) Pop() (byte, error) {
	if s.IsEmpty() {
		return 0, errors.New("Empty Stack")
	}
	topIndex := len(*s) - 1
	element := (*s)[topIndex]
	*s = (*s)[:topIndex]
	return element, nil
}

func (s *Stack) PushArray(items []byte) {
	*s = append(*s, items...)
}

func (s *Stack) PopArray(amount int) ([]byte, error) {
	if len(*s) < amount {
		return []byte{}, errors.New("Not enough items")
	}
	topIndex := len(*s) - 1
	elements := (*s)[topIndex-amount+1 : topIndex+1]
	*s = (*s)[:topIndex-amount+1]
	return elements, nil
}

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
