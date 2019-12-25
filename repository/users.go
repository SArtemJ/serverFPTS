package repository

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type UsersModel struct {
	BaseModel
	GUID   null.String
	Email  null.String
	Wallet null.Int
}

func NewUserModel() *UsersModel {
	model := new(UsersModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type UsersRepository interface {
	Count(b UsersQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*UsersModel, error)
	Create(model *UsersModel) error
	Update(guid string, model *UsersModel) (updatedModelsCount int64, err error)
	Find(b UsersQueryBuilder, offset *int64, limit *int64) ([]*UsersModel, error)
}

type UsersQueryBuilder interface {
	QueryBuilder

	GUID(string) UsersQueryBuilder
	Created(time.Time) UsersQueryBuilder
	Email(string) UsersQueryBuilder
}
