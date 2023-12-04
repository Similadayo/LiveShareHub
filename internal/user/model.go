package user

import (
	"time"
)

type User struct {
	ID             string          `json:"id" gorm:"primary_key;type:varchar(36)"`
	UserName       string          `json:"userName"`
	Password       string          `json:"password"`
	Email          string          `json:"email"`
	FirstName      string          `json:"firstName"`
	LastName       string          `json:"lastName"`
	AvatarURL      string          `json:"avatarURL"`
	Created        time.Time       `json:"created"`
	Updated        time.Time       `json:"updated"`
	Collaborations []Collaboration `json:"collaborations" gorm:"many2many:user_collaborations;"`
}

type Collaboration struct {
	ID        string     `json:"id"`
	ProjectID uint64     `json:"projectId"`
	Name      string     `json:"name"`
	Created   time.Time  `json:"created"`
	Updated   time.Time  `json:"updated"`
	Users     []User     `json:"user" gorm:"many2many:user_collaborations;"`
	Documents []Document `json:"documents" gorm:"many2many:collaboration_documents;"`
}

type Document struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Users   []User    `json:"user" gorm:"many2many:collaboration_documents;"`
}
