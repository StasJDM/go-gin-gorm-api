package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"size:512;not null;index" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserId  uint
	User    User
}
