package presentation

import (
	"e-Commerse/features/users"
	_requestUser "e-Commerse/features/users/presentation/request"
	_responseUser "e-Commerse/features/users/presentation/response"
	_middlewares "e-Commerse/middlewares"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		UserBusiness: business,
	}
}

func (h *UserHandler) AuthLogin(c echo.Context) error {

	var loginUser _requestUser.User

	errBind := c.Bind(&loginUser)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtUser := _requestUser.ToCore(loginUser)
	result, data, err := h.UserBusiness.AuthLoginBusiness(dtUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to login",
		})
	}

	if result == -1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "email not found",
		})
	}

	token, errToken := _middlewares.CreateToken(data.ID, data.Name)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to login data",
		})
	}

	dataResponse := map[string]interface{}{
		"token":  token,
		"ID":     data.ID,
		"name":   data.Name,
		"email":  data.Email,
		"avatar": data.Avatar,
		"phone":  data.Phone,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    dataResponse,
	})

}

func (h *UserHandler) Register(c echo.Context) error {
	var newUser _requestUser.User

	errBind := c.Bind(&newUser)

	validate := validator.New()

	if errValidate := validate.Struct(newUser); errValidate != nil {
		return errValidate
	}

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtUser := _requestUser.ToCore(newUser)
	result, err := h.UserBusiness.RegisterBusiness(dtUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})

}

func (h *UserHandler) Profile(c echo.Context) error {
	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	result, err := h.UserBusiness.ProfileBusiness(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to show detail data",
		})
	}

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	var editUser _requestUser.User

	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	errBind := c.Bind(&editUser)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	dtUser := _requestUser.ToCore(editUser)

	result, err := h.UserBusiness.UpdateProfileBusiness(dtUser, idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})

}

func (h *UserHandler) Destroy(c echo.Context) error {

	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	_, err := h.UserBusiness.DeleteBusiness(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete detail data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})

}
