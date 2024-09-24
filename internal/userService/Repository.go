package userService

import "gorm.io/gorm"

type UsersRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id int, user User) (User, error)
	DeleteUserByID(id int) error
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user User) (User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) UpdateUserByID(id int, user User) (User, error) {
	return user, r.DB.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) DeleteUserByID(id int) error {
	return r.DB.Delete(&User{}, id).Error
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &UserRepository{DB: db}
}
