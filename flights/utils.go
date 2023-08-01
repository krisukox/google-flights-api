package flights

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sort"
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

func getElement[T any](slice []interface{}, index int) (T, bool) {
	if len(slice) <= index {
		var empty T
		return empty, false
	}
	elem, ok := slice[index].(T)
	return elem, ok
}

func getRawElement(slice []interface{}, index int) interface{} {
	if len(slice) <= index {
		var empty interface{}
		return empty
	}

	return slice[index]
}

func readLine(body *bufio.Reader) ([]byte, error) {
	bytesToDecode, isPrefix, err := body.ReadLine()
	if err != nil {
		return nil, err
	}
	if !isPrefix {
		return bytesToDecode, nil
	}
	bytesToDecodeFinal := make([]byte, len(bytesToDecode))
	copy(bytesToDecodeFinal, bytesToDecode)
	for isPrefix {
		bytesToDecode, isPrefix, err = body.ReadLine()
		if err != nil {
			return bytesToDecode, err
		}
		bytesToDecodeFinal = append(bytesToDecodeFinal, bytesToDecode...)
	}
	return bytesToDecodeFinal, nil
}

func getInnerBytes(body *bufio.Reader) ([]byte, error) {
	bytesToDecode, err := readLine(body)
	if err != nil {
		return nil, err
	}

	innerBytes := ""

	err = json.Unmarshal(bytesToDecode, &[][]interface{}{{nil, nil, &innerBytes}})
	if err != nil {
		return nil, fmt.Errorf("error during decoding master schema: %v", err)
	}
	return []byte(innerBytes), nil
}

func sortSlice[T any](slice []T, f func(lv, rv T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return f(slice[i], slice[j])
	})
}
