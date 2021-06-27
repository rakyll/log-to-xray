package main

import (
	"encoding/json"
	"fmt"
	"time"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	for {
		logSpan(logtoxray.Span{
			TraceID:   "122435353",
			SpanID:    "45666",
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Attributes: map[string]string{
				"key": "value",
			},
		})
		time.Sleep(100 * time.Millisecond)
	}
}

func logSpan(span logtoxray.Span) {
	entry, _ := json.Marshal(span)
	fmt.Printf("%s\n", entry)
}
