package logtoxray

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Segment struct {
	ID          string            `json:"segment_id,omitempty"`
	TraceID     string            `json:"trace_id,omitempty"`
	Name        string            `json:"name,omitempty"`
	StartTime   time.Time         `json:"start_time,omitempty"`
	EndTime     time.Time         `json:"end_time,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// TODO(jbd): Use epoc for start and end time.
// Segment should support X-Ray header directly.

func (s *Segment) Key() string {
	return s.TraceID + "-" + s.ID
}

func (s *Segment) Merge(s2 *Segment) {
	// TraceID, SpanID, StartTime and EndTime are not mutable.
	if s2.Name != "" {
		s.Name = s2.Name
	}
	for k, v := range s2.Annotations {
		s.Annotations[k] = v
	}
}

type Consumer struct {
	buffer     map[string]*Segment
	xrayClient *xray.Client
}

func NewConsumer() (*Consumer, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return &Consumer{
		xrayClient: xray.NewFromConfig(awsConfig),
		buffer:     make(map[string]*Segment, 1000),
	}, nil
}

func (c *Consumer) Start(r io.Reader) error {
	decoder := json.NewDecoder(r)
	var span Segment

	for {
		err := decoder.Decode(&span)
		switch err {
		case io.EOF:
			time.Sleep(10 * time.Millisecond)
			continue
		case nil:
			c.handleSpan(&span)
		default:
			log.Printf("error consuming entry: %v", err)
		}
	}
}

func (c *Consumer) handleSpan(s *Segment) {
	if s.TraceID == "" || s.ID == "" {
		log.Printf("Invalid entry; trace_id=%q, span_id=%q", s.TraceID, s.ID)
		return
	}

	key := s.Key()
	if (s.StartTime != time.Time{}) {
		c.buffer[key] = s
	} else {
		prev, ok := c.buffer[key]
		if ok {
			prev.Merge(s)
		}
	}
	if (s.EndTime != time.Time{}) {
		c.send(s)
		delete(c.buffer, key)
	}
}

func (c *Consumer) send(s *Segment) {
	// TODO(jbd): Implement.
	fmt.Println(s)
}
