package repos

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type UserModel struct {
	BaseModel
	GUID   null.String
	Gender null.String
	Email  null.String
	Wallet null.Float
}

func NewUserModel() *UserModel {
	model := new(UserModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type UsersRepo interface {
	Count(b UserQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*UserModel, error)
	Create(model *UserModel) error
	Update(guid string, model *UserModel) (updatedModelsCount int64, err error)
	Find(b UserQueryBuilder, offset *int64, limit *int64) ([]*UserModel, error)
}

type UserQueryBuilder interface {
	QueryBuilder

	GUID(string) UserQueryBuilder
	Created(time.Time) UserQueryBuilder
	Updated(time.Time) UserQueryBuilder
	Email(string) UserQueryBuilder
	Gender(string) UserQueryBuilder
}
