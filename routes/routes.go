package routes

import (
	"e-Commerse/factory"
	"e-Commerse/middlewares"
	_validationUsers "e-Commerse/validation/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()
	// validate := validator.New()

	e.HTTPErrorHandler = _validationUsers.ErrohandlerUser

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	//  login
	e.POST("/login", presenter.UserPresenter.AuthLogin)
	e.POST("/register", presenter.UserPresenter.Register)
	e.GET("/users/:id", presenter.UserPresenter.Profile, middlewares.JWTMiddleware())
	e.PUT("/users/:id", presenter.UserPresenter.UpdateProfile, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", presenter.UserPresenter.Destroy, middlewares.JWTMiddleware())

	// product
	e.POST("/products", presenter.ProductPresenter.CreateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DestroyProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", presenter.ProductPresenter.DetailProduct)
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateProduct, middlewares.JWTMiddleware())
	e.GET("/products", presenter.ProductPresenter.AllProduct)
	e.GET("/my-products", presenter.ProductPresenter.MyProduct, middlewares.JWTMiddleware())

	// cart
	e.POST("/cart/:id", presenter.OrderPresenter.InsertCart, middlewares.JWTMiddleware())

	return e
}
