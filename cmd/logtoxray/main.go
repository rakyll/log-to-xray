package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)
	var span logtoxray.Span
	for {
		if err := decoder.Decode(&span); err != nil {
			log.Fatal(err)
		}
		fmt.Println("----", span)
	}
}
