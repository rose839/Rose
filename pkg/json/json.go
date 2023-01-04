//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

type RawMessage = json.RawMessage

var (
	Marshal       = json.Marshal
	UnMarshal     = json.UnMarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
