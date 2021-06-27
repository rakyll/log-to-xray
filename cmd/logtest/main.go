package main

import (
	"encoding/json"
	"fmt"
	"time"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	logSpan(logtoxray.Span{
		TraceID:   "122435353",
		SpanID:    "45666",
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Attributes: map[string]string{
			"key": "value",
		},
	})
}

func logSpan(span logtoxray.Span) {
	entry, _ := json.Marshal(span)
	fmt.Printf("%s\n", entry)
}
