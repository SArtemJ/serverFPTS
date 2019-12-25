package repository

import "time"

func NewTransactionQuery() TransactionsQueryBuilder {
	return &transactionsQueryBuilder{}
}

type transactionsQueryBuilder struct {
	BaseQueryBuilder
}

func (b *transactionsQueryBuilder) GUID(v string) TransactionsQueryBuilder {
	b.Add("transactions_guid", Equal, v)
	return b
}

func (b *transactionsQueryBuilder) Created(v time.Time) TransactionsQueryBuilder {
	b.Add("transactions.created", Equal, v)
	return b
}

func (b *transactionsQueryBuilder) UserGUID(v string) TransactionsQueryBuilder {
	b.Add("user_guid", Equal, v)
	return b
}

func (b *transactionsQueryBuilder) SourceGUID(v string) TransactionsQueryBuilder {
	b.Add("source_guid", Equal, v)
	return b
}

func (b *transactionsQueryBuilder) Done() TransactionsQueryBuilder {
	b.Add("done", Equal, true)
	return b
}

func (b *transactionsQueryBuilder) OrderById(desc bool) TransactionsQueryBuilder {
	b.OrderBy("id", desc)
	return b
}
