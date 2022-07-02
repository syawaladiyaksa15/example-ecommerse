package factory

import (
	// user
	_userBusiness "e-Commerse/features/users/business"
	_userData "e-Commerse/features/users/data"
	_userPresentation "e-Commerse/features/users/presentation"

	// product
	_productBusiness "e-Commerse/features/products/business"
	_productData "e-Commerse/features/products/data"
	_productPresentation "e-Commerse/features/products/presentation"

	// order
	_orderBusiness "e-Commerse/features/orders/business"
	_orderData "e-Commerse/features/orders/data"
	_orderPresentation "e-Commerse/features/orders/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	ProductPresenter *_productPresentation.ProductHandler
	OrderPresenter   *_orderPresentation.OrderHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// user
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	// product
	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)

	// order
	orderData := _orderData.NewOrderRepository(dbConn)
	orderBusiness := _orderBusiness.NewOrderBusiness(orderData)
	orderPresentation := _orderPresentation.NewOrderHandler(orderBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		ProductPresenter: productPresentation,
		OrderPresenter:   orderPresentation,
	}
}
