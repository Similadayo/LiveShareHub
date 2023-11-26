package user

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Register(user User) (User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUser(user User) (User, error) {
	err := r.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserByID(userID string) (User, error) {
	var user User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
