package entity

import "time"

type Blog struct {
	ID        int       `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Author    string    `json:"author" bson:"author"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
