package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/features/book/domain"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	e.GET("/books", handler.ShowAllBook())
	e.POST("/books", handler.AddBook())
}

func (us *bookHandler) ShowAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

func (us *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}
