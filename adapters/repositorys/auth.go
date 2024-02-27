package repositorys

import (
	"auth/hexagonal/core/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *DB) SignUp(user *domain.User) (*domain.User, error) {
	req := a.db.First(&user, "username = ? ", user.Username)
	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	// สร้าง ID ใหม่
	user.ID = uuid.New().String()
	user.Password = string(hash)

	// เพิ่มข้อมูลในฐานข้อมูล
	req = a.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	fmt.Println(user)
	return user, nil
}

func (a *DB) SignIn(username, password string) (*domain.User, error) {
	// Implement the SignIn logic here
	user, err := a.findUsername(username)
	if err != nil {
		return nil, err
	}

	err = a.verifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *DB) findUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	req := a.db.First(&user, "username = ? ", username)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (a *DB) verifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password not matched")
	}
	return nil
}
