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

func (r *Repository) Login(user User) (User, error) {
	err := r.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
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

func (r *Repository) GetUserByUserName(userName string) (User, error) {
	var user User
	err := r.DB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserProfile(userID string) (User, error) {
	var user User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(user User) (User, error) {
	err := r.DB.Model(&user).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(userID string) error {
	err := r.DB.Where("id = ?", userID).Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FilterUserByName(userName string) ([]User, error) {
	var users []User
	err := r.DB.Where("user_name LIKE ?", "%"+userName+"%").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *Repository) PaginationUser(page int, limit int) ([]User, error) {
	var users []User
	err := r.DB.Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
