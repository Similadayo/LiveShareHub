package collaboration

import (
	"github.com/google/uuid"
	"github.com/similadayo/internal/user"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateCollaboration(collaboration *user.Collaboration) error {
	return r.DB.Create(collaboration).Error
}

func (r *Repository) GetCollaborationByID(collaborationID string) (*user.Collaboration, error) {
	var collaboration *user.Collaboration
	err := r.DB.Preload("Users").First(&collaboration, "id = ?", collaborationID).Error
	return &user.Collaboration{}, err
}

func (r *Repository) GetCollaborationsByUserID(userID string) ([]*user.Collaboration, error) {
	var collaborations []*user.Collaboration
	err := r.DB.Preload("Users").Find(&collaborations, "id = ?", userID).Error
	return collaborations, err
}

func (r *Repository) GetCollaborationsByProjectID(projectID string) ([]*user.Collaboration, error) {
	var collaborations []*user.Collaboration
	err := r.DB.Preload("Users").Find(&collaborations, "id = ?", projectID).Error
	return collaborations, err
}

func (r *Repository) GetCollaborationsByUsers(users []string) ([]*user.Collaboration, error) {
	var collaborations []*user.Collaboration
	err := r.DB.Preload("Users").Find(&collaborations, "id = ?", users).Error
	return collaborations, err
}

func (r *Repository) AddUserToCollaboration(collaborationID string, UserID string) error {
	var collaboration *user.Collaboration
	err := r.DB.Preload("Users").First(&collaboration, "id = ?", collaborationID).Error
	if err != nil {
		return err
	}

	var user user.User
	err = r.DB.First(&user, UserID).Error
	if err != nil {
		return err
	}

	return r.DB.Model(&collaboration).Association("Users").Append(&user)
}

func (r *Repository) RemoveUserFromCollaboration(collaborationID string, UserID string) error {
	var collaboration *user.Collaboration
	err := r.DB.Preload("Users").First(&collaboration, "id = ?", collaborationID).Error
	if err != nil {
		return err
	}

	var user user.User
	err = r.DB.First(&user, UserID).Error
	if err != nil {
		return err
	}

	return r.DB.Model(&collaboration).Association("Users").Delete(&user)
}

func (r *Repository) AddDocumentToCollaboration(collaborationID string, document *user.Document) error {
	var collaboration *user.Collaboration
	err := r.DB.Preload("Documents").First(&collaboration, "id = ?", collaborationID).Error
	if err != nil {
		return err
	}

	document.ID = uuid.New().String()
	return r.DB.Create(document).Error
}
