package repos

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type SourceModel struct {
	BaseModel
	GUID       null.String
	SourceType null.String
}

func NewSourceModel() *SourceModel {
	model := new(SourceModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type SourcesRepo interface {
	Count(b UserQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*SourceModel, error)
	Create(model *SourceModel) error
	Update(guid string, model *SourceModel) (updatedModelsCount int64, err error)
	Find(b UserQueryBuilder, offset *int64, limit *int64) ([]*SourceModel, error)
}

type SourceQueryBuilder interface {
	QueryBuilder

	GUID(string) SourceQueryBuilder
	Created(time.Time) SourceQueryBuilder
	Updated(time.Time) SourceQueryBuilder
	Type(string) SourceQueryBuilder
}
