package model

import "time"

type Post struct {
	ID           uint64     `json:"id"`
	Category     string     `json:"category"`
	Title        string     `json:"title"`
	Link         string     `gorm:"unique" json:"link"`
	Publish_date *time.Time `json:"publish_date"`
	Summary      string     `json:"summary"`
	HistoryList  []History  `gorm:"foreignKey:PostID;constraint:onDelete:SET NULL,onUpdate:CASCADE"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type History struct {
	Old_published_at *time.Time `json:"old_published_at"`
	New_published_at *time.Time `json:"new_published_at"`
	PostID           uint64
	// Post      Post      `gorm:"foreignKey:ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
