package service

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func serviceCreate(featuresName string) {
	name := "./features/" + featuresName + "/service/service.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, serviceContent(featuresName))
		fmt.Println("Service File Created")
	}
}

func serviceContent(featuresName string) string {
	var text = `package service

import (
	"` + generates.ModuleName + `/features/` + featuresName + `"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	Data     ` + featuresName + `.DataInterface
	validate *validator.Validate
}

func New(data ` + featuresName + `.DataInterface) ` + featuresName + `.ServiceInterface {
	return &Service{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *Service) GetAll() ([]` + featuresName + `.Core, error) {
	return s.Data.SelectAll()
}

func (s *Service) GetById(id uint) (` + featuresName + `.Core, error) {
	return s.Data.SelectById(id)
}

func (s *Service) Create(dataCore ` + featuresName + `.Core) (` + featuresName + `.Core, error) {
	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return ` + featuresName + `.Core{}, err
	}

	id, err := s.Data.Store(dataCore)
	if err != nil {
		return ` + featuresName + `.Core{}, err
	}

	return s.Data.SelectById(id)
}

func (s *Service) Update(dataCore ` + featuresName + `.Core, id uint) (` + featuresName + `.Core, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return ` + featuresName + `.Core{}, err
	}

	err := s.Data.Edit(dataCore, id)
	if err != nil {
		return ` + featuresName + `.Core{}, err
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
