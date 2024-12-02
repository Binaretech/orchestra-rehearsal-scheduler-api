package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type SectionService struct {
	db *gorm.DB
}

func NewSectionService(db *gorm.DB) *SectionService {
	return &SectionService{db: db}
}

func (s *SectionService) GetAll() []*model.Section {
	sections := []*model.Section{}

	s.db.Find(&sections)

	return sections
}

func (s *SectionService) GetPaginated(page int, limit int) []*model.Section {
	sections := []*model.Section{}

	s.db.Offset((page - 1) * limit).Limit(limit).Find(&sections)

	return sections
}

func (s *SectionService) GetByID(id uint) *model.Section {
	section := &model.Section{}

	if err := s.db.Preload("Instruments").First(section, id).Error; err != nil {
		return nil
	}

	return section
}

func (s *SectionService) GetByName(name string) *model.Section {
	section := &model.Section{}

	if err := s.db.Where("name = ?", name).First(section).Error; err != nil {
		return nil
	}

	return section
}

type GetSectionMusiciansParams struct {
	Page    int
	Limit   int
	Search  string
	Exclude []int64
}

func (s *SectionService) GetSectionMusicians(id int64, params *GetSectionMusiciansParams) ([]model.User, int64) {
	musicians := []model.User{}
	var totalCount int64

	page := 1
	limit := 20

	if params.Page > 0 {
		page = params.Page
	}

	if params.Limit > 0 {
		limit = params.Limit
	}

	query := s.db.Table("users").
		Joins("Profile").
		Joins("JOIN user_instruments ON user_instruments.user_id = users.id").
		Joins("JOIN sections ON sections.instrument_id = user_instruments.instrument_id").
		Where("sections.id = ?", id)

	if params.Search != "" {
		query = query.Where("users.email LIKE ? OR users.first_name LIKE ? OR users.last_name LIKE ?", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	if len(params.Exclude) > 0 {
		query = query.Where("users.id NOT IN (?)", params.Exclude)
	}

	query.Offset((page - 1) * limit).Limit(limit).Find(&musicians).Count(&totalCount)

	return musicians, totalCount
}

func (s *SectionService) Create(name string, intrumentId int64) *model.Section {
	section := &model.Section{Name: name, InstrumentID: intrumentId}

	s.db.Create(section)

	return section
}

func (s *SectionService) Update(section *model.Section) *model.Section {
	s.db.Save(section)

	return section
}

func (s *SectionService) Delete(section *model.Section) {
	s.db.Delete(section)
}
