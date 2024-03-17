package models

import (
	"time"
)

type Source struct {
	ID        string     `json:"id" bson:"_id"`
	UserId    string     `json:"user_id" bson:"user_id"`
	LastRunAt *time.Time `json:"last_run_at,omitempty" bson:"last_run_at,omitempty"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
