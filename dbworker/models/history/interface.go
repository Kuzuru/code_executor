package models

import (
	"time"
)

type History struct {
	ID        int64      `json:"id" bson:"_id"`
	SourceId  int64      `json:"source_id" bson:"source_id"`
	ExitCode  int64      `json:"exit_code" bson:"exit_code"`
	Stdout    string     `json:"stdout" bson:"stdout"`
	Stderr    string     `json:"stderr" bson:"stderr"`
	OOMKilled bool       `json:"oom_killed" bson:"oom_killed"`
	Timeout   bool       `json:"timeout" bson:"timeout"`
	Duration  int64      `json:"duration" bson:"duration"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
}
