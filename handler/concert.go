package handler

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
)

type CreateConcertRequest struct {
	Repertoire    []string                      `json:"repertoire" validate:"required,dive,required"`
	RehearsalDays []string                      `json:"rehearsalDays" validate:"required,dive,required"`
	Distribution  []service.ConcertDistribution `validate:"required,dive"`
	Title         string                        `json:"title" validate:"required"`
	Location      string                        `json:"location" validate:"required"`
	Date          string                        `json:"date" validate:"required,datetime=2006-01-02"`
	IsDefinitive  bool                          `json:"isDefinitive"`
}

type ConcertHandler struct {
	concertService *service.ConcertService
}

func NewConcertHandler(concertService *service.ConcertService) *ConcertHandler {
	return &ConcertHandler{concertService: concertService}
}

func (h *ConcertHandler) Create(ctx *router.Context) error {
	var body CreateConcertRequest

	if err := ctx.Parse(&body); err != nil {
		return err
	}

	if err := utils.ValidateStruct(&body); err != nil {
		return err
	}

	concert, err := h.concertService.Create(body.Title, body.Date, body.Location, body.IsDefinitive, body.RehearsalDays, body.Distribution)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, concert)
}

func (h *ConcertHandler) Register(r *router.Router) {}

func (h *ConcertHandler) RegisterProtected(r *router.Group) {
	r.Post("/concerts", h.Create)
}
