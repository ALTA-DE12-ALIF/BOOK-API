package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main.go/features/book/domain"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	o := e.Group("/books")
	o.Use(middleware.JWT([]byte("Anakmama!!12")))
	o.GET("", handler.ShowAllBook())
	o.POST("", handler.AddBook())
	o.POST("/update", handler.UpdateBook())
	o.POST("/delete", handler.DeleteBook())
}

func (bs *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}
}

func (bs *bookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.UpdateBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "reg")))
	}
}

func (bs *bookHandler) ShowAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		res, err := bs.srv.ShowAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success get book", ToResponse(res, "all")))
	}
}

func (bs *bookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		err := bs.srv.Delete(cnv)
		if err != nil {
			return c.JSON(http.StatusOK, SuccessResponse("Berhasil Delete", err))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("Gagal Delete"))
	}
}
