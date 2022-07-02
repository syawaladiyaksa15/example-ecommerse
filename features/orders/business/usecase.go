package business

import (
	"e-Commerse/features/orders"
	"errors"
)

type orderUseCase struct {
	orderData orders.Data
}

func NewOrderBusiness(ordData orders.Data) orders.Business {
	return &orderUseCase{
		orderData: ordData,
	}
}

func (uc *orderUseCase) InsertCartBusiness(id, idUserLogin int) (response int, err error) {

	// checking idUserLogin != user_id (product)
	checkOwnerProduct, errCheckOwnerProduct := uc.orderData.CheckingOwnerProduct(id, idUserLogin)

	if errCheckOwnerProduct != nil {
		return 0, errors.New("failed insert product to cart")
	}

	if !checkOwnerProduct {
		return 0, errors.New("failed insert product to cart")
	}

	// checking produk in cart
	checkProduct, idOrder, qtyOrder, subTotal, errCheckProduct := uc.orderData.CheckingCartData(id, idUserLogin)

	if errCheckProduct != nil {
		return 0, errors.New("failed checking produk to cart")
	}

	if checkProduct != 1 {
		response, errInsertCart := uc.orderData.InsertCartData(id, idUserLogin)

		if errInsertCart != nil {
			return 0, errors.New("failed insert product to cart")
		}

		return response, err
	}

	_, errUpdateCart := uc.orderData.UpdateCartData(id, idOrder, qtyOrder, subTotal, idUserLogin)

	if errUpdateCart != nil {
		return 0, errors.New("failed insert product to cart")
	}

	return int(idOrder), err
}
