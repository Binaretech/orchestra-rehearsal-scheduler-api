package handler

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type FamilyHandler struct {
	familyService *service.FamilyService
}

func NewFamilyHandler(familyService *service.FamilyService) *FamilyHandler {
	return &FamilyHandler{familyService: familyService}
}

func (h *FamilyHandler) GetAllData(ctx *router.Context) error {
	families, error := h.familyService.GetAllData()

	if error != nil {
		return error
	}

	return ctx.JSON(200, families)
}
func (h *FamilyHandler) Register(r *router.Router) {
}

func (h *FamilyHandler) RegisterProtected(r *router.Group) {
	r.Get("/families", h.GetAllData)
}
