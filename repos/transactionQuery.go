package repos

import "time"

func NewTransactionQuery() TransactionQueryBuilder {
	return &transactionQueryBuilder{}
}

type transactionQueryBuilder struct {
	BaseQueryBuilder
}

func (b *transactionQueryBuilder) GUID(v string) TransactionQueryBuilder {
	return b
}

func (b *transactionQueryBuilder) Created(v time.Time) TransactionQueryBuilder {
	return b
}

func (b *transactionQueryBuilder) UserGUID(v string) TransactionQueryBuilder {
	return b
}

func (b *transactionQueryBuilder) Source(v string) TransactionQueryBuilder {
	return b
}
