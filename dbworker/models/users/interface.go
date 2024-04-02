package models

type Role int8

const (
	UserRole Role = iota
	AdminRole
)

type User struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	Role         Role   `json:"role" bson:"role"`
}

type UserJWT struct {
	Exp  int    `json:"exp"`
	Id   string `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
