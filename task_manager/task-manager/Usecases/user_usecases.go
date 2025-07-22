package Usecases

import (
	"errors"
	"strings"

	"task_manager/domain"
	"task_manager/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(user domain.User) (domain.User, error)
	Login(email, password string) (string, error)
	PromoteToAdmin(email string) error
	GetUserByEmail(email string) (domain.User, error)
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(user domain.User) (domain.User, error) {
	user.Email = strings.ToLower(user.Email)

	count, err := u.userRepo.CountDocuments(user.Email)
	if err != nil {
		return domain.User{}, err
	}

	role := "user"
	if count == 0 {
		role = "admin" // First user becomes admin
	}
	user.Role = role

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = string(hashedPassword)

	createdUser, err := u.userRepo.InsertOne(user)
	if err != nil {
		return domain.User{}, err
	}

	createdUser.Password = "" // never return password
	return createdUser, nil
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (u *userUsecase) PromoteToAdmin(email string) error {
	return u.userRepo.UpdateRole(email, "admin")
}

func (u *userUsecase) GetUserByEmail(email string) (domain.User, error) {
	return u.userRepo.GetUserByEmail(email)
}
