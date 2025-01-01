package handler

import (
	"strconv"
	"time"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
)

type CalendarHandler struct {
	calendarService *service.CalendarService
}

func NewCalendarHandler(calendarService *service.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarService: calendarService}
}
func (h *CalendarHandler) GetEntries(ctx *router.Context) error {
	month := ctx.Query("month")
	year := ctx.Query("year")
	offset := ctx.Query("offset")

	if month == "" {
		month = strconv.Itoa(int(time.Now().Month()))
	}

	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}

	if offset == "" {
		offset = "0:00"
	}

	if _, err := strconv.Atoi(month); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if _, err := strconv.Atoi(year); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if !utils.ValidateTimeZoneOffset(offset) {
		offset = "0:00"
	}

	entries, err := h.calendarService.GetEntriesPerMonth(month, year, offset)
	if err != nil {
		return err
	}

	return ctx.JSON(200, NewResource(entries))
}

func (h *CalendarHandler) GetDateEntries(ctx *router.Context) error {
	month := ctx.Query("month")
	year := ctx.Query("year")
	day := ctx.Query("day")

	offset := ctx.Query("offset")

	if month == "" {
		month = strconv.Itoa(int(time.Now().Month()))
	}

	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}

	if day == "" {
		day = strconv.Itoa(time.Now().Day())
	}

	if offset == "" {
		offset = "0:00"
	}

	if _, err := strconv.Atoi(month); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if _, err := strconv.Atoi(year); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if _, err := strconv.Atoi(day); err != nil {
		return errors.NewBadRequestError(errors.INVALID_DATA_TYPE)
	}

	if !utils.ValidateTimeZoneOffset(offset) {
		offset = "0:00"
	}

	entries, err := h.calendarService.GetEntriesPerDate(day, month, year, offset)
	if err != nil {
		return err
	}

	return ctx.JSON(200, NewResource(entries))
}

func (h *CalendarHandler) Register(r *router.Router) {
}

func (h *CalendarHandler) RegisterProtected(r *router.Group) {
	r.Get("/calendar", h.GetEntries)
	r.Get("/calendar/date", h.GetDateEntries)

}
