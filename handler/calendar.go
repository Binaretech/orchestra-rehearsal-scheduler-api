package handler

import (
	"strconv"
	"time"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type CalendarHandler struct {
	calendarService *service.CalendarService
}

func NewCalendarHandler(calendarService *service.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarService: calendarService}
}
func (h *CalendarHandler) GetEntries(ctx *router.Context) error {
	month := ctx.Param("month")
	year := ctx.Param("year")

	if month == "" {
		month = strconv.Itoa(int(time.Now().Month()))
	}

	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}

	if _, err := strconv.Atoi(month); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if _, err := strconv.Atoi(year); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	entries, err := h.calendarService.GetEntriesPerDate(month, year)

	if err != nil {
		return err
	}

	return ctx.JSON(200, NewResource(entries))
}

func (h *CalendarHandler) Register(r *router.Router) {
}

func (h *CalendarHandler) RegisterProtected(r *router.Group) {
	r.Get("/calendar", h.GetEntries)
}
