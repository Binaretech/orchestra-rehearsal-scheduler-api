package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type CreateSection struct {
	Name         string `json:"name"`
	InstrumentID uint   `json:"instrumentId"`
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

func (h *SectionHandler) GetMusicians(ctx *router.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	search := ctx.Query("search")

	queryExclude := ctx.Query("exclude")
	excludeIDs := []int64{}

	if queryExclude != "" {
		for _, idStr := range strings.Split(queryExclude, ",") {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err == nil {
				excludeIDs = append(excludeIDs, id)
			}
		}
	}

	musicians, count := h.sectionService.GetSectionMusicians(id, &service.GetSectionMusiciansParams{
		Page:    page,
		Limit:   limit,
		Search:  search,
		Exclude: excludeIDs,
	})

	return ctx.JSON(http.StatusOK, Resource[model.User]{Data: musicians, Total: count})
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

func (h *SectionHandler) Register(r *router.Router) {}

func (h *SectionHandler) RegisterProtected(r *router.Group) {
	r.Get("/sections", h.Get)
	r.Get("/sections/{id}/musicians", h.GetMusicians)
	r.Get("/sections/{id}", h.GetById)
	r.Post("/sections", h.Create)
}
