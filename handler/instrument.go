package handler

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type CreateInstrumentRequest struct {
	Name      string `json:"name"`
	SectionID int64  `json:"sectionId"`
}

type InstrumentHandler struct {
	instrumentService *service.InstrumentService
}

func NewInstrumentHandler(instrumentService *service.InstrumentService) *InstrumentHandler {
	return &InstrumentHandler{instrumentService: instrumentService}
}

func (h *InstrumentHandler) Create(ctx *router.Context) error {
	var body CreateInstrumentRequest

	if err := ctx.Parse(&body); err != nil {
		return err
	}

	instrument, err := h.instrumentService.Create(body.Name, body.SectionID)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, instrument)
}

func (h *InstrumentHandler) Register(r *router.Router) {}

func (h *InstrumentHandler) RegisterProtected(r *router.Group) {
	r.Post("/instruments", h.Create)
}
