package testData

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"gopkg.in/guregu/null.v3"
	"time"
)

func GetTestUsers() []repository.UsersModel {
	items := []repository.UsersModel{
		{
			BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:      null.StringFrom("c99cec6c-7a34-4941-a988-33b52ca5c3ec"),
			Email:     null.StringFrom("u1email@test.com"),
			Wallet:    null.IntFrom(101),
		},
		{
			BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:      null.StringFrom("5c1ad16e-89dd-4abf-96d3-7920b279bd02"),
			Email:     null.StringFrom("u2email@test.com"),
			Wallet:    null.IntFrom(102),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("a52a078d-f69c-4b39-92c0-b1781e7d664c"),
			Email:  null.StringFrom("u3email@test.com"),
			Wallet: null.IntFrom(103),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("8fcd8eba-01df-45c5-872f-fb2576b7e9c1"),

			Email:  null.StringFrom("u4email@test.com"),
			Wallet: null.IntFrom(201),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("ede90d80-a6b3-4176-adc9-7dd21a8fb34b"),

			Email:  null.StringFrom("u5email@test.com"),
			Wallet: null.IntFrom(202),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),

			Email:  null.StringFrom("u6email@test.com"),
			Wallet: null.IntFrom(203),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("aa4d02cb-2ca9-4e34-8b39-5e49559c1136"),

			Email:  null.StringFrom("u10email@test.com"),
			Wallet: null.IntFrom(204),
		},
		//users for repeated transactions
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("6fd884b2-e612-447b-8417-4cb69fc278b3"),
			Email:  null.StringFrom("u7email@test.com"),
			Wallet: null.IntFrom(205),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("b5b92823-d95c-42cd-a218-2ce56c77dd17"),

			Email:  null.StringFrom("u8email@test.com"),
			Wallet: null.IntFrom(301),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID: null.StringFrom("996b4fa1-13d0-40fe-95fa-d1fe3bdc32d1"),

			Email:  null.StringFrom("u9email@test.com"),
			Wallet: null.IntFrom(302),
		},
	}

	return items
}
