package repository

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type SourcesModel struct {
	BaseModel
	GUID       null.String
	SourceType null.String
}

func NewSourceModel() *SourcesModel {
	model := new(SourcesModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type SourcesRepository interface {
	Count(b SourcesQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*SourcesModel, error)
	Create(model *SourcesModel) error
	Update(guid string, model *SourcesModel) (updatedModelsCount int64, err error)
	Find(b SourcesQueryBuilder, offset *int64, limit *int64) ([]*SourcesModel, error)
}

type SourcesQueryBuilder interface {
	QueryBuilder

	GUID(string) SourcesQueryBuilder
	Created(time.Time) SourcesQueryBuilder
	Type(string) SourcesQueryBuilder
}
