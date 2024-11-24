package seeders

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/db"
	model "github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	faker "github.com/brianvoe/gofakeit/v7"
)

// InitialSeeder seeds the database with initial data
func InitialSeeder() error {
	query, err := db.Connect()
	tx := query.Begin()

	// Create family
	family := model.Family{Name: "family " + faker.Word()}
	err = tx.Create(&family).Error

	// Create a new section
	section := model.Section{Name: "section " + faker.Word(), FamilyID: uint(family.ID)}
	err = tx.Create(&section).Error

	instrument := model.Instrument{Name: "instrument " + faker.Word(), SectionID: section.ID}
	err = tx.Create(&instrument).Error

	// Create concert
	concert := model.Concert{Title: "concert " + faker.LoremIpsumSentence(6), Location: faker.City(), ConcertDate: faker.Date(), Description: faker.LoremIpsumSentence(12), ConcertDateStatus: "definitive"}
	err = tx.Create(&concert).Error

	// Create rehearsal
	rehearsal := model.Rehearsal{RehearsalDate: faker.FutureDate(), RehearsalTime: faker.FutureDate(), Location: faker.Address().Address, IsGeneral: true, ConcertID: uint(concert.ID)}
	err = tx.Create(&rehearsal).Error

	// Create users
	var users []model.User
	for i := 0; i < 10; i++ {
		user := model.User{
			Email:    faker.Email(),
			Password: faker.Password(true, true, true, true, false, 10),
			Role:     "user",
			Profile: model.Profile{
				FirstName: faker.Name(),
				LastName:  faker.LastName(),
			},
		}
		tx.Model(&user).Association("Instruments").Append(instrument)

		users = append(users, user)
	}

	err = tx.Create(&users).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Model(&rehearsal).Association("Users").Append(users); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
