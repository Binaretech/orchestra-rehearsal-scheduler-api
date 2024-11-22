package handler

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type CreateSection struct {
	Name string `json:"name"`
}

type SectionHandler struct {
	sectionService *service.SectionService
}

func NewSectionHandler(sectionService *service.SectionService) *SectionHandler {
	return &SectionHandler{sectionService: sectionService}
}

func (h *SectionHandler) Create(ctx *router.Context) error {
	body := CreateSection{}

	if err := ctx.Parse(&body); err != nil {
		return err
	}

	if exists := h.sectionService.GetByName(body.Name); exists != nil {
		return ctx.Error(http.StatusConflict, errors.CONFLICT)
	}

	section := h.sectionService.Create(body.Name)

	if section == nil {
		return ctx.Error(http.StatusInternalServerError, errors.INTERNAL_ERROR)
	}

	return ctx.JSON(http.StatusOK, section)

}

func (h *SectionHandler) Register(r *router.Router) {
}

func (h *SectionHandler) RegisterProtected(r *router.Group) {
	r.Post("/sections", h.Create)
}
