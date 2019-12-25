package repository

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
)

type TransactionModel struct {
	BaseModel
	GUID   null.String
	State  null.String
	Amount null.Int
	Source null.String
	User   null.String
	Done   null.Bool
}

func NewTransactionModel() *TransactionModel {
	model := new(TransactionModel)
	model.GUID = null.StringFrom(uuid.New().String())

	return model
}

type TransactionsRepository interface {
	Count(b TransactionsQueryBuilder) (int64, error)
	Collection(offset *int64, limit *int64) ([]*TransactionModel, error)
	Create(model *TransactionModel) error
	Update(guid string, model *TransactionModel) (updatedModelsCount int64, err error)
	Find(b TransactionsQueryBuilder, offset *int64, limit *int64) ([]*TransactionModel, error)
	FindOdd(b TransactionsQueryBuilder, limit *int64) ([]*TransactionModel, error)
}

type TransactionsQueryBuilder interface {
	QueryBuilder

	GUID(string) TransactionsQueryBuilder
	Created(time.Time) TransactionsQueryBuilder
	UserGUID(string) TransactionsQueryBuilder
	SourceGUID(string) TransactionsQueryBuilder
	OrderById(desc bool) TransactionsQueryBuilder
	Done() TransactionsQueryBuilder
}
