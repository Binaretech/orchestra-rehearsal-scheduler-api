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

	violin := model.Instrument{Name: "Violin"}
	viola := model.Instrument{Name: "Viola"}
	cello := model.Instrument{Name: "Cello"}
	doubleBass := model.Instrument{Name: "Double Bass"}
	flute := model.Instrument{Name: "Flute"}
	oboe := model.Instrument{Name: "Oboe"}
	clarinet := model.Instrument{Name: "Clarinet"}
	bassoon := model.Instrument{Name: "Bassoon"}
	frenchHorn := model.Instrument{Name: "French Horn"}
	trumpet := model.Instrument{Name: "Trumpet"}
	trombone := model.Instrument{Name: "Trombone"}
	tuba := model.Instrument{Name: "Tuba"}
	percussion := model.Instrument{Name: "Percussion"}
	piano := model.Instrument{Name: "Piano"}

	instruments := []*model.Instrument{
		&violin,
		&viola,
		&cello,
		&doubleBass,
		&flute,
		&oboe,
		&clarinet,
		&bassoon,
		&frenchHorn,
		&trumpet,
		&trombone,
		&tuba,
		&percussion,
		&piano,
	}

	for _, instrument := range instruments {
		tx.Create(instrument)

	}

	families := []model.Family{
		{
			Name: "Strings",
			Sections: []model.Section{
				{
					Name:         "1st Violins",
					InstrumentID: violin.ID,
				},
				{
					Name:         "2nd Violins",
					InstrumentID: violin.ID,
				},
				{
					Name:         "Violas",
					InstrumentID: viola.ID,
				},
				{
					Name:         "Cellos",
					InstrumentID: cello.ID,
				},
				{
					Name:         "Double Basses",
					InstrumentID: doubleBass.ID,
				},
			},
		},
		{
			Name: "Woodwinds",
			Sections: []model.Section{
				{
					Name:         "1st Flute",
					InstrumentID: flute.ID,
				},
				{
					Name:         "2nd Flute",
					InstrumentID: flute.ID,
				},
				{
					Name:         "1st Oboe",
					InstrumentID: oboe.ID,
				},
				{
					Name:         "2nd Oboe",
					InstrumentID: oboe.ID,
				},
				{
					Name:         "1st Clarinet",
					InstrumentID: clarinet.ID,
				},
				{
					Name:         "2nd Clarinet",
					InstrumentID: clarinet.ID,
				},
				{
					Name:         "1st Bassoon",
					InstrumentID: bassoon.ID,
				},
				{
					Name:         "2nd Bassoon",
					InstrumentID: bassoon.ID,
				},
			},
		},
		{
			Name: "Brass",
			Sections: []model.Section{
				{
					Name:         "1st French Horn",
					InstrumentID: frenchHorn.ID,
				},
				{
					Name:         "2nd French Horn",
					InstrumentID: frenchHorn.ID,
				},
				{
					Name:         "1st Trumpet",
					InstrumentID: trumpet.ID,
				},
				{
					Name:         "2nd Trumpet",
					InstrumentID: trumpet.ID,
				},
				{
					Name:         "1st Trombone",
					InstrumentID: trombone.ID,
				},
				{
					Name:         "2nd Trombone",
					InstrumentID: trombone.ID,
				},
				{
					Name:         "Tubas",
					InstrumentID: tuba.ID,
				},
			},
		},
		{
			Name: "Percussion",
			Sections: []model.Section{
				{
					Name:         "Percussion",
					InstrumentID: percussion.ID,
				},
			},
		},
		{
			Name: "Others",
			Sections: []model.Section{
				{
					Name:         "Piano",
					InstrumentID: piano.ID,
				},
			},
		},
	}

	tx.Create(families)

	// // Create family
	// family := model.Family{Name: "family " + faker.Word()}
	// err = tx.Create(&family).Error

	// // Create a new section
	// section := model.Section{Name: "section " + faker.Word(), FamilyID: uint(family.ID)}
	// err = tx.Create(&section).Error

	// instrument := model.Instrument{Name: "instrument " + faker.Word(), Sections: []model.Section{{}}}

	// err = tx.Create(&instrument).Error

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
			Role:     model.USER_ADMIN_ROLE,
			Profile: model.Profile{
				FirstName: faker.Name(),
				LastName:  faker.LastName(),
			},
		}
		// tx.Model(&user).Association("Instruments").Append(instrument)

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
