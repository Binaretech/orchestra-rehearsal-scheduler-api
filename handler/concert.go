package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
)

// CreateConcertRequest holds the request payload to create a new concert
type CreateConcertRequest struct {
	Repertoire    []string                      `json:"repertoire" validate:"required,dive,required"`
	RehearsalDays []string                      `json:"rehearsalDays" validate:"required,dive,required,datetime=2006-01-02T15:04:05"`
	Distribution  []service.ConcertDistribution `json:"distribution" validate:"required,dive"`
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

func (h *ConcertHandler) Find(ctx *router.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	concert, err := h.concertService.Find(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, NewResource(concert))
}

func (h *ConcertHandler) Create(ctx *router.Context) error {
	var body CreateConcertRequest

	if err := ctx.Parse(&body); err != nil {
		return err
	}
	if err := utils.ValidateStruct(&body); err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		return err
	}
	now := time.Now().Truncate(24 * time.Hour)
	if parsedDate.Before(now) {
		return errors.NewBadRequestError(errors.CONCERT_PAST_DATE)
	}

	for _, dayStr := range body.RehearsalDays {
		parsedDay, err := time.Parse("2006-01-02T15:04:05", dayStr)
		if err != nil {
			return err
		}
		if parsedDay.Before(time.Now()) {
			return errors.NewBadRequestError(errors.CONCERT_PAST_REHEARSAL_DATE)
		}
	}

	concert, err := h.concertService.Create(
		body.Title,
		body.Date,
		body.Location,
		body.IsDefinitive,
		body.RehearsalDays,
		body.Distribution,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, NewResource(concert))
}

func (h *ConcertHandler) Register(r *router.Router) {
}

func (h *ConcertHandler) RegisterProtected(r *router.Group) {
	r.Get("/concerts/{id}", h.Find)
	r.Post("/concerts", h.Create)
}
