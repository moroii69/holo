package server

import (
	"encoding/json"
	"log"
	"time"
)

type logFields map[string]any

type logEvent struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
	Message string    `json:"msg"`
	Fields  logFields `json:"fields,omitempty"`
}

func logJSON(level, msg string, fields logFields) {
	e := logEvent{
		Level:   level,
		Time:    time.Now().UTC(),
		Message: msg,
		Fields:  fields,
	}
	b, err := json.Marshal(e)
	if err != nil {
		log.Printf("log-marshal-error: %v", err)
		return
	}
	log.Print(string(b))
}

func logInfo(msg string, fields logFields) {
	logJSON("info", msg, fields)
}

func logError(msg string, fields logFields) {
	logJSON("error", msg, fields)
}

