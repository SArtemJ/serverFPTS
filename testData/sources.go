package testData

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"gopkg.in/guregu/null.v3"
	"time"
)

func GetTestSources() []repository.SourcesModel {
	items := []repository.SourcesModel{
		{
			BaseModel:  repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:       null.StringFrom("08262e6c-2d70-45b9-9450-f1a36a25697c"),
			SourceType: null.StringFrom("game"),
		},
		{
			BaseModel:  repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:       null.StringFrom("09e38898-d758-4c30-bea9-c9f8cd28a37e"),
			SourceType: null.StringFrom("server"),
		},
		{
			BaseModel:  repository.BaseModel{Created: null.TimeFrom(time.Now())},
			GUID:       null.StringFrom("ab223aa6-5b78-4f27-9c26-5bda5701b66e"),
			SourceType: null.StringFrom("payment"),
		},
	}

	return items
}
