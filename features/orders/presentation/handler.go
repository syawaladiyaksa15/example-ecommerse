package presentation

import (
	"e-Commerse/features/orders"
	_middlewares "e-Commerse/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderBusiness orders.Business
}

func NewOrderHandler(business orders.Business) *OrderHandler {
	return &OrderHandler{
		OrderBusiness: business,
	}
}

func (h *OrderHandler) InsertCart(c echo.Context) error {
	id := c.Param("id")

	idProduct, _ := strconv.Atoi(id)

	idUserLogin, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	result, err := h.OrderBusiness.InsertCartBusiness(idProduct, idUserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed insert product to cart",
		})
	}

	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed insert product to cart",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": map[string]int{
			"IDOrder": result,
		},
	})

}
