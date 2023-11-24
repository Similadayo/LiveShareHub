package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	FullName      string
	Avatar        string
	Bio           string
	SocialLinks   string // JSON-encoded array of social media links
	IsAdmin       bool
	LastLogin     *gorm.DeletedAt
	Collaborators []User `gorm:"many2many:user_collaborators;"` // many to many relationship
}
