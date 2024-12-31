package service

import (
	"gorm.io/gorm"
)

type Entry struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Date  string `json:"date"`
	Title string `json:"title"`
}

type CalendarService struct {
	db *gorm.DB
}

func NewCalendarService(db *gorm.DB) *CalendarService {
	return &CalendarService{db: db}
}
func (s *CalendarService) GetEntriesPerDate(month string, year string) (map[string][]Entry, error) {

	var results []Entry

	subQueryConcerts := s.db.Table("concerts").
		Select("id, 'concert' AS type, concert_date AS date, title").
		Where("EXTRACT(MONTH FROM concert_date) = ? AND EXTRACT(YEAR FROM concert_date) = ?", month, year)

	subQueryRehearsals := s.db.Table("rehearsals").
		Select("rehearsals.id, 'rehearsal' AS type, date, concerts.title").
		Joins("JOIN concerts ON concerts.id = rehearsals.concert_id").
		Where("EXTRACT(MONTH FROM date) = ? AND EXTRACT(YEAR FROM date) = ?", month, year)

	if err := s.db.Table("(?) as entries", s.db.Raw("? UNION ALL ?", subQueryConcerts, subQueryRehearsals)).
		Order("type ASC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	dateMap := make(map[string][]Entry)
	for _, entry := range results {
		date := entry.Date[:10]
		dateMap[date] = append(dateMap[date], entry)
	}

	return dateMap, nil
}
