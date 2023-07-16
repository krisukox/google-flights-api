package flights

import (
	"bufio"
	"fmt"
)

func skipPrefix(body *bufio.Reader) error {
	var isPrefix bool
	var err error
	for i := 0; i < 2; i++ {
		_, isPrefix, err = body.ReadLine()
		if err != nil || isPrefix {
			return fmt.Errorf("error when reading response with the serialized city names: %w", err)
		}
	}
	return nil
}

func getElement[T any](slice []interface{}, index int) T {
	elem, ok := getRawElement(slice, index).(T)
	if !ok {
		var empty T
		return empty
	}
	return elem
}

func getRawElement(slice []interface{}, index int) interface{} {
	if len(slice) <= index {
		var empty interface{}
		return empty
	}

	return slice[index]
}