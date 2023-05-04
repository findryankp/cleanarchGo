package delivery

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func handlerCreate(featuresName string) {
	name := "./features/" + featuresName + "/delivery/handler.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, handlerContent(featuresName))
		fmt.Println("Handler File Created")
	}
}

func handlerContent(featuresName string) string {
	var text = `package delivery

import (
	"` + generates.ModuleName + `/features/` + featuresName + `"
	"` + generates.ModuleName + `/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service ` + featuresName + `.ServiceInterface
}

func New(s ` + featuresName + `.ServiceInterface) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetAll(c echo.Context) error {
	dataCore, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	dataResponse := ListCoreToResponse(dataCore)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataResponse))
}

func (t *Handler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	dataCore, err := t.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}

	dataResponse := CoreToResponse(dataCore)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataResponse))
}

func (t *Handler) Create(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	dataResponse, err := t.Service.Create(RequestToCore(&request))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", CoreToResponse(dataResponse)))
}

func (t *Handler) Update(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	dataCore, err := t.Service.Update(RequestToCore(&request), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", CoreToResponse(dataCore)))
}

func (t *Handler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := t.Service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
	
	`

	return text
}
