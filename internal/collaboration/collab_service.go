package collaboration

import (
	"time"

	"github.com/google/uuid"
	"github.com/similadayo/internal/user"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateCollaboration(projectId uint64, name string, userIDs []string) (*user.Collaboration, error) {
	collaboration := &user.Collaboration{
		ID:        uuid.New().String(),
		ProjectID: projectId,
		Name:      name,
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	err := s.Repo.CreateCollaboration(collaboration)
	if err != nil {
		return nil, err
	}

	for _, userID := range userIDs {
		err = s.Repo.AddUserToCollaboration(collaboration.ID, userID)
		if err != nil {
			return nil, err
		}
	}

	return collaboration, nil
}

func (s *Service) GetCollaborationByID(collaborationID string) (*user.Collaboration, error) {
	return s.Repo.GetCollaborationByID(collaborationID)
}

func (s *Service) InviteUserToCollaboration(collaborationID string, userID string) error {
	return s.Repo.AddUserToCollaboration(collaborationID, userID)
}

func (s *Service) RemoveUserFromCollaboration(collaborationID string, userID string) error {
	return s.Repo.RemoveUserFromCollaboration(collaborationID, userID)
}

func (s *Service) GetCollaborationsByUsers(users []string) ([]*user.Collaboration, error) {
	return s.Repo.GetCollaborationsByUsers(users)
}

func (s *Service) GetCollaborationsByUserID(userID string) ([]*user.Collaboration, error) {
	return s.Repo.GetCollaborationsByUserID(userID)
}

func (s *Service) GetCollaborationsByProjectID(projectID string) ([]*user.Collaboration, error) {
	return s.Repo.GetCollaborationsByProjectID(projectID)
}

func (s *Service) CreateDocumentInCollaboration(collaborationID string, name string, title string, content string) (*user.Document, error) {
	document := &user.Document{
		ID:      uuid.New().String(),
		Name:    name,
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	err := s.Repo.AddDocumentToCollaboration(collaborationID, document)
	if err != nil {
		return nil, err
	}

	return document, nil
}
