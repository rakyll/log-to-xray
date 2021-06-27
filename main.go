package logtoxray

import (
	"time"
)

type Span struct {
	TraceID    string            `json:"trace_id,omitempty"`
	SpanID     string            `json:"span_id,omitempty"`
	Name       string            `json:"name,omitempty"`
	StartTime  time.Time         `json:"start_time,omitempty"`
	EndTime    time.Time         `json:"end_time,omitempty"`
	Attributes map[string]string `json:"attrs,omitempty"`
}
