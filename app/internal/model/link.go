package model

import "time"

type Link struct {
	ID        uint64    `json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Source    string    `gorm:"unique" json:"source"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
