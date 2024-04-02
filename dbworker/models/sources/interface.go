package models

import (
	"time"
)

type Source struct {
	ID        string     `json:"id" bson:"_id"`
	UserId    string     `json:"user_id" bson:"user_id"`
	FileName  string     `json:"file_name" bson:"file_name"`
	Data      string     `json:"data" bson:"data"`
	LastRunAt *time.Time `json:"last_run_at,omitempty" bson:"last_run_at,omitempty"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type SourceDTO struct {
	UserId   string `json:"user_id" bson:"user_id"`
	FileName string `json:"filename" bson:"filename"`
	Data     string `json:"data" bson:"data"`
}

type SourceRunDTO struct {
	SourceID string `json:"source_id"`
	Command  string `json:"command"`
	Files    []struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"files"`
}

type SourceRunResponse struct {
	ExitCode  int     `json:"exit_code"`
	Stdout    string  `json:"stdout"`
	Stderr    string  `json:"stderr"`
	OomKilled bool    `json:"oom_killed"`
	Timeout   bool    `json:"timeout"`
	Duration  float64 `json:"duration"`
}
