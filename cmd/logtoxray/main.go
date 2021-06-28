package main

import (
	"log"
	"os"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	c, err := logtoxray.NewConsumer()
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Start(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
