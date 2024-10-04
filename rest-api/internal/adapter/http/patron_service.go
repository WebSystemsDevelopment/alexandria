package adapter

import (
	"net/http"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/in"
	"github.com/labstack/echo/v4"
)

type PatronHandler struct {
	service in.PatronService
}

func NewPatronHandler(s in.PatronService) PatronHandler {
	return PatronHandler{service: s}
}

// CreatePatron creates a new patron
// @Summary Create a new patron
// @Description Create a new patron and save it to the database
// @Tags patrons
// @Accept  json
// @Produce  json
// @Param patron body domain.Patron true "Patron data"
// @Success 201 {object} domain.Patron
// @Failure 400 {object} echo.HTTPError "Bad request"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /patrons [post]
func (handler *PatronHandler) CreatePatron(context echo.Context) error {
	patron := new(domain.Patron)
	if err := context.Bind(patron); err != nil {
		return err
	}
	err := handler.service.CreatePatron(patron)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusCreated, patron)
}

// GetAllBooks gets all patrons
// @Summary Get all patrons
// @Description Get a list of all patrons from the database
// @Tags patrons
// @Produce  json
// @Success 200 {array} domain.Patron
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /patrons [get]
func (handler *PatronHandler) GetAllPatrons(context echo.Context) error {
	Books, err := handler.service.GetAllPatrons()
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, Books)
}

