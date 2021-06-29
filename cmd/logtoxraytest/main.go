package main

import (
	"encoding/json"
	"fmt"
	"time"

	logtoxray "github.com/rakyll/log-to-xray"
)

func main() {
	// Start segment.
	logSegment(logtoxray.Segment{
		ID:        "0f910026178b71eb",
		TraceID:   "1-5880168b-fd515828bs07678a3bb5a78c",
		Name:      "foo",
		StartTime: float64(time.Now().UnixNano()) / 1000,
	})

	// Add metadata.
	logSegment(logtoxray.Segment{
		ID:      "0f910026178b71eb",
		TraceID: "1-5880168b-fd515828bs07678a3bb5a78c",
		Annotations: map[string]string{
			"service": "auth",
		},
	})
	// Finish segment.
	logSegment(logtoxray.Segment{
		ID:      "0f910026178b71eb",
		TraceID: "1-5880168b-fd515828bs07678a3bb5a78c",
		EndTime: float64(time.Now().UnixNano()) / 1000,
	})

	// Alternatively, log the entire segment as a single entry.
	logSegment(logtoxray.Segment{
		ID:        "0f910026178b71ef",
		TraceID:   "1-5880168b-fd515828bs07678a3bb5a78c",
		Name:      "foo2",
		StartTime: float64(time.Now().UnixNano()) / 1000,
		EndTime:   float64(time.Now().UnixNano()) / 1000,
		Annotations: map[string]string{
			"service": "auth",
		},
	})
}

func logSegment(s logtoxray.Segment) {
	entry, _ := json.Marshal(s)
	fmt.Printf("%s\n", entry)
}
