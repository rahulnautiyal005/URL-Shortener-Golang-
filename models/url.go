package models

import "time"

type URL struct {
	ID          int64     `json:"id" bson:"_id"`
	LongURL     string    `json:"url" bson:"long_url"`
	ShortCode   string    `json:"short_code" bson:"short_code"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ExpiresAt   time.Time `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	ClickCount  int       `json:"click_count" bson:"click_count"`
}
