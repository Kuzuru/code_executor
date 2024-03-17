package models

type Role int8

const (
	UserRole Role = iota
	AdminRole
)

type User struct {
	ID           int64  `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	Role         Role   `json:"role" bson:"role"`
}