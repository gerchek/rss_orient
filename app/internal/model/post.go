package model

import "time"

type Post struct {
	ID        uint64    `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Link      string    `gorm:"unique" json:"link"`
	Date      time.Time `json:"date"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
