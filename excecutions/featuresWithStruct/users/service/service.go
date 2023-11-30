package service

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func serviceCreate() {
	name := "./features/users/service/service.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, serviceContent())
		fmt.Println("Service File Created")
	}
}

func serviceContent() string {
	var text = `package service

import (
	"errors"
	"` + generates.ModuleName + `/features/users"
	"` + generates.ModuleName + `/utils/helpers"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	Data     users.DataInterface
	validate *validator.Validate
}

func New(data users.DataInterface) users.ServiceInterface {
	return &Service{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *Service) Login(email, password string) (users.Core, error) {
	if email == "" || password == "" {
		return users.Core{}, errors.New("email and password must be fill")
	}

	user, err := s.Data.SelectByEmail(email)
	if err != nil || !helpers.CheckPasswordHash(password, user.Password) {
		return users.Core{}, errors.New("user and password not found")
	}

	return user, nil
}

func (s *Service) ChangePassword(id uint, oldPassword, newPassword, confirmPssword, hash string) error {
	if oldPassword == "" || newPassword == "" || confirmPssword == "" {
		return errors.New("old password,new password, and confirm password cannot be empty")
	}

	if newPassword != confirmPssword {
		return errors.New("new password and confirm password must be similarity")
	}

	user, _ := s.Data.SelectById(id)
	if !helpers.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("old password not match with exist password")
	}

	user.Password = newPassword
	return s.Data.Edit(user, id)
}

func (s *Service) GetAll() ([]users.Core, error) {
	return s.Data.SelectAll()
}

func (s *Service) GetById(id uint) (users.Core, error) {
	return s.Data.SelectById(id)
}

func (s *Service) Create(dataCore users.Core) (users.Core, error) {
	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return users.Core{}, err
	}

	if dataCore.Role == "" {
		dataCore.Role = "user"
	}
	password, _ := helpers.HashPassword(dataCore.Password)
	dataCore.Password = password

	id, err := s.Data.Store(dataCore)
	if err != nil {
		return users.Core{}, err
	}

	return s.Data.SelectById(id)
}

func (s *Service) Update(dataCore users.Core, id uint) (users.Core, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return users.Core{}, err
	}

	err := s.Data.Edit(dataCore, id)
	if err != nil {
		return users.Core{}, err
	}
	return s.Data.SelectById(id)
}

func (s *Service) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
`
	return text
}
