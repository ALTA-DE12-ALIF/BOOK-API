package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/features/pinjam/domain"
)

type pinjamHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := pinjamHandler{srv: srv}
	e.GET("/pinjams", handler.ShowAllPinjam())
	e.POST("/pinjams", handler.AddPinjam())
}

func (us *pinjamHandler) ShowAllPinjam() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllPinjam()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

func (us *pinjamHandler) AddPinjam() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input PinjamFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddPinjam(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}
