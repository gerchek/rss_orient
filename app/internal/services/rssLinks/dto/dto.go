package dto

type LinkDto struct {
	ID   uint64 `json:"id"`
	Link string `gorm:"unique" json:"link"`
}
