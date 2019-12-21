package repos

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type TransactionModel struct {
	BaseModel
	GUID   null.String
	State  null.String
	Amount null.Float
	Source null.String
}

func NewTransactionModel() *TransactionModel {
	model := new(TransactionModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type TransactionsRepo interface {
	Count(b UserQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*TransactionModel, error)
	Create(model *TransactionModel) error
	Update(guid string, model *TransactionModel) (updatedModelsCount int64, err error)
	Find(b UserQueryBuilder, offset *int64, limit *int64) ([]*TransactionModel, error)
}

type TransactionQueryBuilder interface {
	QueryBuilder

	GUID(string) TransactionQueryBuilder
	Created(time.Time) TransactionQueryBuilder
	UserGUID(string) TransactionQueryBuilder
	Source(string) TransactionQueryBuilder
}
