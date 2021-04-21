package model

import (
	"time"

	"gorm.io/gorm"
)

type Burger struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index:unique" json:"name"`
	ImageURL    string `json:"imageUrl"`
	Ingredients string `json:"ingredients"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"default:null" json:"-"`
}
