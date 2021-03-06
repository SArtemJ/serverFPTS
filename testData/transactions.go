package testData

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"gopkg.in/guregu/null.v3"
	"time"
)

func GetTestTransactions() []repository.TransactionModel {
	items := []repository.TransactionModel{
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("f67cfdaf-1223-4424-9bf3-aada6992f440"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(101),
			Source: null.StringFrom("game"),
			User:   null.StringFrom("c99cec6c-7a34-4941-a988-33b52ca5c3ec"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("0c6e6bf2-623b-44ef-9982-2633b1c94911"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(101),
			Source: null.StringFrom("game"),
			User:   null.StringFrom("c99cec6c-7a34-4941-a988-33b52ca5c3ec"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("1f0630cd-b0cf-4205-ac0b-990101425e62"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(101),
			Source: null.StringFrom("game"),
			User:   null.StringFrom("5c1ad16e-89dd-4abf-96d3-7920b279bd02"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("c881a246-bb98-4fe6-8e68-673b7f3b0292"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(102),
			Source: null.StringFrom("server"),
			User:   null.StringFrom("5c1ad16e-89dd-4abf-96d3-7920b279bd02"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("3c540a79-c154-4663-954c-b7d7d867a083"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(102),
			Source: null.StringFrom("server"),
			User:   null.StringFrom("a52a078d-f69c-4b39-92c0-b1781e7d664c"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("61a9e65d-0567-4b12-9330-db2d1281a07c"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(102),
			Source: null.StringFrom("server"),
			User:   null.StringFrom("a52a078d-f69c-4b39-92c0-b1781e7d664c"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("34f61f9f-e0e5-43c2-bf1e-a23a5c035446"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(102),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("8fcd8eba-01df-45c5-872f-fb2576b7e9c1"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("67af7b08-16fe-4ee2-96b2-ab713f5adf5c"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("8fcd8eba-01df-45c5-872f-fb2576b7e9c1"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("58bd0f09-6407-47bb-be72-7134ab130a67"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("ede90d80-a6b3-4176-adc9-7dd21a8fb34b"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("296db792-c6d1-40fa-9b33-87a87dde31d6"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("game"),
			User:   null.StringFrom("ede90d80-a6b3-4176-adc9-7dd21a8fb34b"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("4567d580-7463-4fec-a90d-f88b940f6ebe"),
			State:  null.StringFrom("win"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("server"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("33805c34-64cf-495a-a15e-0f12a29bfaad"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("3fb5a6be-3fd9-450b-863a-0129936b14e2"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("7b83ae37-cc04-4efe-bd26-684e38c5484d"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("b94414c6-5faf-4e27-8cec-71052efb1434"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("d463fc51-bc1a-4d15-96ab-85c6e4187952"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("9ab0e6a8-5d41-4d92-ac39-d557dce74e17"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("2371ad80-0390-4448-ac1f-bdffed3fce52"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("b65c981d-c8f8-4c70-90c8-3369336f6f83"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("2956af6d-bc85-454e-852c-8a5f2d545cea"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("3b549970-12bf-4e6e-9cbd-86f414d05b8a"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("a7e5822f-434e-4a0f-b1cf-c70d23a68174"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("976c14d9-afaa-4915-bfa4-d61849a90314"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("c0da8cf9-397e-4226-93ee-7e94875e7366"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
		{BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:   null.StringFrom("bc8dcde5-b2a9-4dc9-bc49-8714de395ac1"),
			State:  null.StringFrom("lost"),
			Amount: null.IntFrom(103),
			Source: null.StringFrom("payment"),
			User:   null.StringFrom("4dbbdfcf-d3f1-4a29-81d4-0167fb3b8a24"),
		},
	}

	return items
}
