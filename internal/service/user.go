package userservice

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	auth "SchoolMarket-run-with-go-/internal/middleware"
	"SchoolMarket-run-with-go-/internal/model"
	"SchoolMarket-run-with-go-/internal/repository"
)

type CreateUserRequest struct {
	Name string
	Psw  string
}

type LoginUserRequest struct {
	Name string
	Password string
}

type DeleteUserRequest struct {
	Id uint
}

type GetAimByNameRequest struct {
	Name string
}

type GetAimUserRequest struct {
	Id uint
}

type UpdateUserRequest struct {
	Id      uint
	NewName *string
	NewPsw  *string
}

type UserService interface {
	CreateUser(req CreateUserRequest) (*model.User, error)
	Login(req LoginUserRequest) (*model.User, string, error)
	Delete(req DeleteUserRequest) error
	GetUserById(req GetAimUserRequest) (*model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(req UpdateUserRequest) (*model.User, error)
}

type userService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req CreateUserRequest) (*model.User, error) {
	if req.Name == "" {
		return nil, errors.New("user name is required")
	}
	if req.Psw == "" {
		return nil, errors.New("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Psw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name: req.Name,
		Psd:  string(hash),
	}

	return user, s.repo.CreateUser(user)
}

func (s *userService) Login(req LoginUserRequest) (*model.User, string, error) {
	user, err := s.GetUserByName(req.Name)
	if err != nil {
		return nil, "", errors.New("用户名或密码错误")
	}
	if user == nil {
		return nil, "", errors.New("用户名或密码错误")
	}
	erro := bcrypt.CompareHashAndPassword([]byte(user.Psd), []byte(req.Password))
	if erro != nil {
		return nil, "", erro
	}
	token,err := auth.GenerateToken(user.ID, user.Name)
	if err != nil {
		return nil, "", errors.New("服务器错误")
	}
	return user, token, nil
}

func (s *userService) Delete(req DeleteUserRequest) error {
	if req.Id == 0 {
		return errors.New("invalid user id")
	}

	rows, err := s.repo.DeleteUser(req.Id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (s *userService) GetUserByName(name string) (*model.User, error) {
	user, err := s.repo.FindByName(name)
	return user, err
}

func (s *userService) GetUserById(req GetAimUserRequest) (*model.User, error) {
	if req.Id == 0 {
		return nil, errors.New("invalid user id")
	}

	user, err := s.repo.FindAim(req.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) GetAllUser() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) UpdateUser(req UpdateUserRequest) (*model.User, error) {
	if req.Id == 0 {
		return nil, errors.New("invalid user id")
	}
	if req.NewName == nil && req.NewPsw == nil {
		return nil, errors.New("no fields to update")
	}

	user, err := s.repo.FindAim(req.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if req.NewName != nil {
		if *req.NewName == "" {
			return nil, errors.New("user name is required")
		}
		user.Name = *req.NewName
	}

	if req.NewPsw != nil {
		if *req.NewPsw == "" {
			return nil, errors.New("password is required")
		}
		user.Psd = *req.NewPsw
	}

	return s.repo.UpdateUser(user)
}
