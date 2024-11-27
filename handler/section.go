package handler

import (
	"net/http"
	"strconv"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type CreateSection struct {
	Name         string `json:"name"`
	InstrumentID int64  `json:"instrumentId"`
}

type SectionHandler struct {
	sectionService *service.SectionService
}

func NewSectionHandler(sectionService *service.SectionService) *SectionHandler {
	return &SectionHandler{sectionService: sectionService}
}

func (h *SectionHandler) Get(ctx *router.Context) error {
	sections := h.sectionService.GetAll()

	return ctx.JSON(http.StatusOK, sections)
}

func (h *SectionHandler) GetById(ctx *router.Context) error {
	id := ctx.Param("id")

	sectionID, _ := strconv.ParseUint(id, 10, 64)

	section := h.sectionService.GetByID(uint(sectionID))

	if section == nil {
		return ctx.Error(http.StatusNotFound, errors.NOT_FOUND)
	}

	return ctx.JSON(http.StatusOK, section)

}

func (h *SectionHandler) Create(ctx *router.Context) error {
	body := CreateSection{}

	if err := ctx.Parse(&body); err != nil {
		return err
	}

	if exists := h.sectionService.GetByName(body.Name); exists != nil {
		return ctx.Error(http.StatusConflict, errors.CONFLICT)
	}

	section := h.sectionService.Create(body.Name, body.InstrumentID)

	if section == nil {
		return ctx.Error(http.StatusInternalServerError, errors.INTERNAL_ERROR)
	}

	return ctx.JSON(http.StatusOK, section)

}

func (h *SectionHandler) Update(ctx *router.Context) error {
	return nil
}

func (h *SectionHandler) Delete(ctx *router.Context) error {
	return nil
}

func (h *SectionHandler) Register(r *router.Router) {}

func (h *SectionHandler) RegisterProtected(r *router.Group) {
	r.Get("/sections", h.Get)
	r.Get("/sections/{id}", h.GetById)
	r.Post("/sections", h.Create)
	r.Put("/sections", h.Update)
	r.Delete("/sections", h.Delete)
}
