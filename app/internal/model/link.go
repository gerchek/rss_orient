package model

import "time"

type Link struct {
	ID        uint64    `json:"id"`
	Link      string    `gorm:"unique" json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
