package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"go_practice.com/domain"
)

type userRepository struct {
	db      *sql.DB
	timeout time.Duration
}

// GetAllUsers implements domain.IUserRepository.
func (u userRepository) GetAllUsers() (users []domain.User, err error) {
	var listusers []domain.User
	query := `SELECT * FROM USERS;`
	result, err := u.db.Query(query)
	if err != nil {
		log.Fatalf("Error happened while retrieving all records from database, [error]-%v", err)
		return nil, err
	}
	for result.Next() {
		var user domain.User
		if err := result.Scan(&user.Id, &user.UserEmail, &user.UserName, &user.Password); err != nil {
			return users, err
		}
		listusers = append(listusers, user)
	}
	return listusers, nil
}

// GetUserByEmail implements domain.IUserRepository.
func (u userRepository) GetUserByEmail(user_email string) (userId domain.User, err error) {
	var desireduser domain.User
	query := `SELECT * FROM USERS WHERE user_email = $1 LIMIT 1;`
	errs := u.db.QueryRow(query, user_email).Scan(&desireduser.Id, &desireduser.UserName, &desireduser.UserEmail, &desireduser.Password)
	fmt.Printf("id = %s, name %s, email = %s, password=%s", desireduser.Id, desireduser.UserName, desireduser.UserEmail, desireduser.Password)
	if errs != nil {
		fmt.Printf("There is no user with that email %v, [error]- %v", user_email, errs)
		return desireduser, errs
	}

	return desireduser, nil
}

// InsertUserToDatabase implements domain.IUserRepository.
func (u userRepository) InsertUserToDatabase(user domain.User) (iuser domain.User, err error) {
	query := `INSERT INTO USERS (id,user_name,user_email,pass_word) VALUES ($1,$2,$3,$4);`
	_, errs := u.db.Exec(query, user.Id, user.UserName, user.UserEmail, user.Password)
	if errs != nil {
		log.Fatalf("Error at user repository [error]: %v", err)
		return
	}
	//var structuser domain.User
	structuser := domain.User{
		Id:        user.Id,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Password:  user.Password,
	}
	return structuser, nil
}

func NewUserRepository(db *sql.DB, timeout time.Duration) domain.IUserRepository {
	return userRepository{
		db:      db,
		timeout: timeout,
	}
}
