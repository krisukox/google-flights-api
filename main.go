// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// func main() {
// 	body := `<html>
//     <body>
//         <h1>Main title</h1>
//         <a href="https://code-maven.com/">Code Maven</a>
//         <h2 id="subtitle" class="important">Some subtle title</h2>
//     </body>
//     </html>`

// 	reader := strings.NewReader(body)
// 	tokenizer := html.NewTokenizer(reader)
// 	for {
// 		tt := tokenizer.Next()
// 		if tt == html.ErrorToken {
// 			if tokenizer.Err() == io.EOF {
// 				return
// 			}
// 			fmt.Printf("Error: %v", tokenizer.Err())
// 			return
// 		}
// 		tag, hasAttr := tokenizer.TagName()
// 		fmt.Printf("Tag: %v\n", string(tag))
// 		if hasAttr {
// 			for {
// 				attrKey, attrValue, moreAttr := tokenizer.TagAttr()
// 				// if string(attrKey) == "" {
// 				//     break
// 				// }
// 				fmt.Printf("Attr: %v\n", string(attrKey))
// 				fmt.Printf("Attr: %v\n", string(attrValue))
// 				fmt.Printf("Attr: %v\n", moreAttr)
// 				if !moreAttr {
// 					break
// 				}
// 			}
// 		}
// 	}
// }

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
	"golang.org/x/text/currency"
)

func main() {
	departureDate, _ := time.Parse("2006-01-02", "2023-08-12")
	returnDate, _ := time.Parse("2006-01-02", "2023-08-19")
	url, err := api.CreateFullURL(departureDate, returnDate, "Wrocław", "Rzym")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(url)
	// bodyReader := bytes.NewReader([]byte(``))

	flights, err := api.GetOffers(url, currency.PLN)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, flight := range flights {
		fmt.Printf("%+v\n", flight)
	}

}
