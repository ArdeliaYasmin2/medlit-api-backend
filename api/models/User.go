package models

import (
	"errors"
	"time"

	"medlit-api-backend/api/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// helper
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ConfirmPassword(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("password does not match")
	}
	return nil
}

// repo
type UserInterface interface {
	GetNameByEmail(email string) string
	GetUserByID(uid uint) (User, error)
	LoginCheck(email string, password string) (string, error)
	RegisterCheck(password string, confirmPassword string) error
	SaveUser(data User) (*User, error)
}

func (u *Repo) GetNameByEmail(email string) string {
	var user User
	DB.Where("email = ?", email).Take(&user)
	return user.Name
}

func (u *Repo) GetUserByID(uid uint) (User, error) {
	var user User

	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found")
	}

	user.Password = ""

	return user, nil
}

func (u *Repo) LoginCheck(email string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *Repo) RegisterCheck(password string, confirmPassword string) error {
	if err := ConfirmPassword(password, confirmPassword); err != nil {
		return err
	}
	return nil
}

func (u *Repo) SaveUser(data User) (*User, error) {
	err := DB.Create(&data).Error
	if err != nil {
		return &User{}, err
	}
	return &data, nil
}
