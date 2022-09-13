package dto

type LinkDto struct {
	Name   string `gorm:"not null" json:"name"`
	Source string `gorm:"unique" json:"source"`
}
