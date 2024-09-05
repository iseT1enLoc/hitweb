package repository

import (
	"log"
	"time"

	"go_practice.com/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db      *gorm.DB
	timeout time.Duration
}

// GetAllUsers implements domain.IUserRepository.
func (u userRepository) GetAllUsers() (users []domain.User, err error) {
	var listusers []domain.User
	result := u.db.Find(&listusers)
	if result.Error != nil {
		log.Fatalf("Error happened while retrieving all records from database, [error]-%v", result.Error)
		return nil, result.Error
	}
	return listusers, nil
}

// GetUserByEmail implements domain.IUserRepository.
func (u userRepository) GetUserByEmail(user_email string) (userId string, err error) {
	var desireduser domain.User
	result := u.db.Where("user_email = ?", user_email).First(&desireduser)
	if result.Error != nil {
		//log.Fatalf("There is no user with that email %v, [error]- %v", user_email, result.Error)
		return "", result.Error
	}
	return desireduser.Id, nil
}

// InsertUserToDatabase implements domain.IUserRepository.
func (u userRepository) InsertUserToDatabase(user domain.User) (iuser domain.User, err error) {
	inserted_user := u.db.Create(&user)
	if inserted_user.Error != nil {
		log.Fatalf("Error at user repository [error]: %v", inserted_user.Error)
		return
	}
	return iuser, nil
}

func NewUserRepository(db *gorm.DB, timeout time.Duration) domain.IUserRepository {
	return userRepository{
		db:      db,
		timeout: timeout,
	}
}
