package main

import (
	"fmt"
	"log"
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
		// fmt.Println(string(c))
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
	a := []byte{8, 28, 16, 2, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 51, 106, 12, 8, 2, 18, 8, 47, 109, 47, 48, 56, 52, 53, 98, 114, 12, 8, 3, 18, 8, 47, 109,
		47, 48, 53, 121, 119, 103, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 55, 106, 12, 8, 3, 18, 8, 47, 109, 47, 48, 53, 121, 119, 103, 114, 12, 8, 2, 18, 8,
		47, 109, 47, 48, 56, 52, 53, 98, 64, 1, 72, 1, 112, 1, 130, 1, 11, 8, 255, 255, 255, 255, 255, 255, 255, 255, 255, 1, 152, 1, 1}
	log.Print(createURL(a))
	// log.Println("Test")

	// requestURL := "https://www.google.com/_/TravelFrontendUi/data/batchexecute?rpcids=H028ib&source-path=%2Ftravel%2Fflights%2Fsearch&f.sid=-8421128425468344897&bl=boq_travel-frontend-ui_20230613.06_p0&hl=pl&soc-app=162&soc-platform=1&soc-device=1&_reqid=444052&rt=c"

	// jsonBody := []byte(`f.req=%5B%5B%5B%22H028ib%22%2C%22%5B%5C%22Pr%5C%22%2C%5B1%2C2%2C3%2C5%2C4%5D%2Cnull%2C%5B1%2C1%2C1%5D%2C1%5D%22%2Cnull%2C%22generic%22%5D%5D%5D&at=AAuQa1qJpLKW2Hl-i40OwJyzmo22%3A1687083247610&`)
	// bodyReader := bytes.NewReader(jsonBody)

	// req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	// req.Header.Set("authority", "www.google.com")
	// req.Header.Set("accept", "*/*")
	// req.Header.Set("accept-language", "pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7,hr;q=0.6")
	// req.Header.Set("cache-control", "no-cache")
	// req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	// req.Header.Set("cookie", "CONSENT=YES+srp.gws-20211208-0-RC2.pl+FX+371")
	// req.Header.Set("origin", "https://www.google.com")
	// req.Header.Set("pragma", "no-cache")
	// req.Header.Set("referer", "https://www.google.com/travel/flights/search?tfs=CBwQAhonEgoyMDIzLTA3LTA0agwIAhIIL20vMDg0NWJyCwgDEgcvbS8wbjJ6GicSCjIwMjMtMDctMDhqCwgDEgcvbS8wbjJ6cgwIAhIIL20vMDg0NWJAAUgBcAGCAQsI____________AZgBAQ")
	// req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	// req.Header.Set("sec-ch-ua-arch", "\"x86\"")
	// req.Header.Set("sec-ch-ua-bitness", "\"64\"")
	// req.Header.Set("sec-ch-ua-full-version", "\"113.0.5672.92\"")
	// req.Header.Set("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"113.0.5672.92\", \"Chromium\";v=\"113.0.5672.92\", \"Not-A.Brand\";v=\"24.0.0.0\"")
	// req.Header.Set("sec-ch-ua-mobile", "?0")
	// req.Header.Set("sec-ch-ua-model", "")
	// req.Header.Set("sec-ch-ua-platform", "Linux")
	// req.Header.Set("sec-ch-ua-platform-version", "5.19.0")
	// req.Header.Set("sec-ch-ua-wow64", "?0")
	// req.Header.Set("sec-fetch-dest", "empty")
	// req.Header.Set("sec-fetch-mode", "cors")
	// req.Header.Set("sec-fetch-site", "same-origin")
	// req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	// req.Header.Set("x-client-data", "CKS1yQEIj7bJAQimtskBCKmdygEIivHKAQiSocsBCIegzQEI8abNAQj9qs0BCKCtzQEItrLNAQicuc0BCNm5zQEI3L3NAQ==")
	// req.Header.Set("x-goog-ext-190139975-jspb", "[\"PL\",\"ZZ\",\"Xkt6eg==\"]")
	// req.Header.Set("x-goog-ext-259736195-jspb", "[\"pl\",\"PL\",\"PLN\",1,null,[-120],null,[[48764690,48627726,47907128,48480739,48710756,48676280,48593234,48707378]],1,[]]")
	// req.Header.Set("x-same-domain", "1")

	// client := http.Client{
	// 	Timeout: 30 * time.Second,
	// }

	// resp, err := client.Do(req)

	// // resp, err := http.Post("https://www.google.com/_/TravelFrontendUi/data/batchexecute?rpcids=H028ib&source-path=%2Ftravel%2Fflights%2Fsearch&f.sid=-8421128425468344897&bl=boq_travel-frontend-ui_20230613.06_p0&hl=pl&soc-app=162&soc-platform=1&soc-device=1&_reqid=444052&rt=c", "application/json; charset=utf-8", &buf)
	// if err != nil {
	// 	log.Fatalf("Couldn't send a request")
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// fmt.Print(string(body))
}
