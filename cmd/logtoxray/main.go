package main

import (
	"log"
	"os"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	c := logtoxray.NewConsumer()
	if err := c.Start(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
