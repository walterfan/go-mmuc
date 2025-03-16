package message

import (
	"fmt"
	"time"
)

// Json Messaging and Presence Protocol
type Request struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Seq     int32  `json:"seq"`
	TrackId string `json:"trackId"`
	//msec milliseconds since January 1, 1970 UTC.
	Timestamp int64  `json:"timestamp"`
	From      string `json:"from"`
	To        string `json:"to"`
	Message   string `json:"message"`
}

type Response struct {
	Request
	Code int32  `json:"code"`
	Desc string `json:"desc"`
}

// String method for Request
func (r Request) String() string {
	return fmt.Sprintf("Request{Name: %s, Type: %s, Seq: %d, TrackId: %s, Timestamp: %s, From: %s, To: %s, Message: %s}",
		r.Name, r.Type, r.Seq, r.TrackId, time.UnixMilli(r.Timestamp).String(), r.From, r.To, r.Message)
}

// String method for Response
func (r Response) String() string {
	return fmt.Sprintf("Response{ %s, Code: %d, Desc: %s}",
		r.Request.String(), r.Code, r.Desc)
}
