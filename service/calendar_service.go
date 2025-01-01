package service

import (
	"fmt"
	"strings"

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

func (s *CalendarService) GetEntriesPerMonth(month string, year string, offset string) (map[string][]Entry, error) {
	var results []Entry

	interval := fmt.Sprintf("%s hour", offset)
	monthYear := fmt.Sprintf("%s-%s", year, month)

	subQueryConcerts := s.db.Table("concerts").
		Select("id, 'concert' AS type, concert_date + ?::interval AS date, title", interval).
		Where("to_char(concert_date + ?::interval, 'YYYY-FMMM') = ?", interval, monthYear)

	subQueryRehearsals := s.db.Table("rehearsals").
		Select("rehearsals.id, 'rehearsal' AS type, date + ?::interval AS date, concerts.title", interval).
		Joins("JOIN concerts ON concerts.id = rehearsals.concert_id").
		Where("to_char(date + ?::interval, 'YYYY-FMMM') = ?", interval, monthYear)

	if err := s.db.Table("(?) as entries", s.db.Raw("? UNION ALL ?", subQueryConcerts, subQueryRehearsals)).
		Order("type ASC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	dateMap := make(map[string][]Entry)
	for _, entry := range results {
		dateParts := strings.Split(entry.Date[:10], "-")
		month := strings.TrimLeft(dateParts[1], "0")
		day := strings.TrimLeft(dateParts[2], "0")

		date := dateParts[0] + "-" + month + "-" + day

		dateMap[date] = append(dateMap[date], entry)
	}

	return dateMap, nil
}

func (s *CalendarService) GetEntriesPerDate(day string, month string, year string, offset string) ([]Entry, error) {
	var results []Entry

	interval := fmt.Sprintf("%s hour", offset)
	date := fmt.Sprintf("%s-%s-%s", year, month, day)

	subQueryConcerts := s.db.Table("concerts").
		Select("id, 'concert' AS type, concert_date + ?::interval AS date, title", interval).
		Where("(concert_date + ?::interval)::date = ?::date", interval, date)

	subQueryRehearsals := s.db.Table("rehearsals").
		Select("rehearsals.id, 'rehearsal' AS type, date + ?::interval AS date, concerts.title", interval).
		Joins("JOIN concerts ON concerts.id = rehearsals.concert_id").
		Where("(date + ?::interval)::date = ?::date", interval, date)

	err := s.db.Table("(?) as entries", s.db.Raw("? UNION ALL ?", subQueryConcerts, subQueryRehearsals)).
		Order("type ASC").
		Scan(&results).Error

	return results, err
}
