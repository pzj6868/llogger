package llogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	logruslog "github.com/sirupsen/logrus"
	"time"
)

// MyFormatter 自定义 formatter
type MyFormatter struct {
	PrettyPrint bool
}

type data struct {
	Time    string `json:"time"`
	Err     error  `json:"err,omitempty"`
	File    string `json:"file"`
	Level   string `json:"level"`
	Msg     string `json:"msg"`
}

// Format implement the Formatter interface
func (mf *MyFormatter) Format(entry *logruslog.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var data data
	for k, v := range entry.Data {
		switch k {
		case "files":
			data.File = v.(string)
		case "errors":
			data.Err = v.(error)
		default:

		}
	}
	data.Time = entry.Time.Format(time.RFC3339)
	data.Msg = entry.Message
	level, _ := entry.Level.MarshalText()
	data.Level = string(level)

	encoder := json.NewEncoder(b)
	if mf.PrettyPrint{
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
