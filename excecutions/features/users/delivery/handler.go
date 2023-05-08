package delivery

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func handlerCreate() {
	name := "./features/users/delivery/handler.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, handlerContent())
		fmt.Println("Handler File Created")
	}
}

func handlerContent() string {
	var text = `package delivery

import (
	"net/http"
	"strconv"
	"` + generates.ModuleName + `/apps/middlewares"
	"` + generates.ModuleName + `/features/users"
	"` + generates.ModuleName + `/utils/helpers"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service users.ServiceInterface
}

func New(s users.ServiceInterface) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Login(c echo.Context) error {
	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, helpers.ResponseFail("error bind data"))
	}

	user, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		return c.JSON(401, helpers.ResponseFail(err.Error()))
	}

	token, _ := middlewares.CreateToken(int(user.Id), user.Role)
	response := map[string]any{
		"token": token,
		"user":  CoreToResponse(user),
	}

	return c.JSON(200, helpers.ResponseSuccess("login success", response))
}

func (h *Handler) Register(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, helpers.ResponseFail("error bind data"))
	}

	dataResponse, err := h.Service.Create(RequestToCore(&request))
	if err != nil {
		return c.JSON(400, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(201, helpers.ResponseSuccess("create data success", CoreToResponse(dataResponse)))
}

func (h *Handler) ChangePassword(c echo.Context) error {
	user_id := middlewares.ClaimsToken(c).Id

	r := EditPassword{}
	if err := c.Bind(&r); err != nil {
		return c.JSON(400, helpers.ResponseFail("error bind data"))
	}

	hash, _ := helpers.HashPassword(r.NewPassword)
	if err := h.Service.ChangePassword(uint(user_id), r.OldPassword, r.NewPassword, r.ConfirmPassword, hash); err != nil {
		return c.JSON(400, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("change password success", nil))
}

func (h *Handler) ClaimsToken(c echo.Context) error {
	response := map[string]any{
		"user": middlewares.ClaimsToken(c),
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", response))
}

func (h *Handler) GetAll(c echo.Context) error {
	dataCore, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(500, helpers.ResponseFail("error read data"))
	}
	dataResponse := ListCoreToResponse(dataCore)
	return c.JSON(200, helpers.ResponseSuccess("-", dataResponse))
}

func (h *Handler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	dataCore, err := h.Service.GetById(id)
	if err != nil {
		return c.JSON(404, helpers.ResponseFail("data not found"))
	}

	dataResponse := CoreToResponse(dataCore)
	return c.JSON(200, helpers.ResponseSuccess("-", dataResponse))
}

func (h *Handler) Update(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, helpers.ResponseFail("error bind data"))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	dataCore, err := h.Service.Update(RequestToCore(&request), uint(id))
	if err != nil {
		return c.JSON(404, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(200, helpers.ResponseSuccess("update data success", CoreToResponse(dataCore)))
}

func (h *Handler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Service.Delete(uint(id)); err != nil {
		return c.JSON(404, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(200, helpers.ResponseSuccess("delete data success", nil))
}
	
	`

	return text
}
