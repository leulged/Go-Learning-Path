package usecases

import (
	"strings"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"
	"task_manager/Domain/interfaces"
	"task_manager/utils"
)

type UserUsecase interface {
	Register(user entities.User) (entities.User, error)
	Login(email, password string) (string, error)
	PromoteToAdmin(email string) error
	GetUserByEmail(email string) (entities.User, error)
}

type userUsecase struct {
	userRepo      interfaces.UserRepository
	tokenService  interfaces.TokenService
}

func NewUserUsecase(userRepo interfaces.UserRepository, tokenService interfaces.TokenService) UserUsecase {
	return &userUsecase{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (u *userUsecase) Register(user entities.User) (entities.User, error) {
	// Validate input
	if err := utils.ValidateEmail(user.Email); err != nil {
		return entities.User{}, err
	}
	if err := utils.ValidateName(user.Name); err != nil {
		return entities.User{}, err
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		return entities.User{}, err
	}

	user.Email = strings.ToLower(user.Email)

	count, err := u.userRepo.CountDocuments(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	if count > 0 {
		return entities.User{}, errors.EmailAlreadyExistsError{}
	}

	// First user becomes admin
	if count == 0 {
		user.SetRole("admin")
	} else {
		user.SetRole("user")
	}

	// Hash password using utils
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = hashedPassword

	createdUser, err := u.userRepo.InsertOne(user)
	if err != nil {
		return entities.User{}, err
	}

	// Never return password
	createdUser.Password = ""
	return createdUser, nil
}

func (u *userUsecase) Login(email, password string) (string, error) {
	// Validate email
	if err := utils.ValidateEmail(email); err != nil {
		return "", errors.InvalidCredentialsError{}
	}

	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.InvalidCredentialsError{}
	}

	// Check password using utils
	if !utils.CheckPassword(password, user.Password) {
		return "", errors.InvalidCredentialsError{}
	}

	token, err := u.tokenService.GenerateToken(user.Email, user.Role)
	if err != nil {
		return "", errors.InvalidCredentialsError{}
	}

	return token, nil
}

func (u *userUsecase) PromoteToAdmin(email string) error {
	return u.userRepo.UpdateRole(email, "admin")
}

func (u *userUsecase) GetUserByEmail(email string) (entities.User, error) {
	return u.userRepo.GetUserByEmail(email)
} 