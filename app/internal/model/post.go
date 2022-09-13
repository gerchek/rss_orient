package model

import "time"

type Post struct {
	ID          uint64    `json:"id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Link        string    `gorm:"unique" json:"link"`
	Date        time.Time `json:"date"`
	Summary     string    `json:"summary"`
	HistoryList []History `gorm:"foreignKey:PostID;constraint:onDelete:SET NULL,onUpdate:CASCADE"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type History struct {
	Updated string `json:"updated"`
	PostID  uint64
	// Post      Post      `gorm:"foreignKey:ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
