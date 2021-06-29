package main

import (
	"encoding/json"
	"fmt"
	"time"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	for {
		logSegment(logtoxray.Segment{
			TraceID:   "122435353",
			ID:        "45666",
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Annotations: map[string]string{
				"key": "value",
			},
		})
		time.Sleep(100 * time.Millisecond)
	}
}

func logSegment(s logtoxray.Segment) {
	entry, _ := json.Marshal(s)
	fmt.Printf("%s\n", entry)
}
