package main

import (
	"fmt"
	"log"

	"github.com/krisukox/google-flights-api/api"
)

func createURL(a []byte) string {
	var chars = [...]byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_'}

	var c []byte
	var d byte // or chars[64]
	e, f := 0, 0
	for ; e < len(a)-2; e += 3 {
		h := a[e]
		k := a[e+1]
		l := a[e+2]
		m := chars[h>>2]
		h = chars[(h&3)<<4|k>>4]
		k = chars[(k&15)<<2|l>>6]
		l = chars[l&63]
		c = append(c, m)
		c = append(c, h)
		c = append(c, k)
		c = append(c, l)
		f += 1
	}

	m := 0
	l := d

	switch len(a) - e {
	case 2:
		m := a[e+1]
		element_id := int((m & 15) << 2)
		if len(chars) > element_id {
			l = chars[element_id]
		} else {
			l = d
		}
	case 1:
		a_prim := int(a[e])
		c = append(c, chars[a_prim>>2])
		c = append(c, chars[(a_prim&3)<<4|m>>4])
		c = append(c, l)
		c = append(c, d)
	}

	cos := string(c)

	fmt.Println(cos)

	return string(c[:])
}

func main() {
	// a := []byte{8, 28, 16, 2, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 51, 106, 12, 8, 2, 18, 8, 47, 109, 47, 48, 56, 52, 53, 98, 114, 12, 8, 3, 18, 8, 47, 109,
	// 	47, 48, 53, 121, 119, 103, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 55, 106, 12, 8, 3, 18, 8, 47, 109, 47, 48, 53, 121, 119, 103, 114, 12, 8, 2, 18, 8,
	// 	47, 109, 47, 48, 56, 52, 53, 98, 64, 1, 72, 1, 112, 1, 130, 1, 11, 8, 255, 255, 255, 255, 255, 255, 255, 255, 255, 1, 152, 1, 1}
	// log.Print(createURL(a))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api.GetSerializedCityName()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Couldn't read body")
	// }
	// fmt.Print(body)
	// var j interface{}
	// err = json.NewDecoder(resp.Body).Decode(&j)
	// if err != nil {
	// 	log.Fatalf("Couldn't decode response")
	// }
	// fmt.Print(j)
}
