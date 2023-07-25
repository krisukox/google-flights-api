package flights

import (
	"bufio"
	"encoding/json"
	"fmt"
)

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
