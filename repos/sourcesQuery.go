package repos

import "time"

func NewSourceQuery() SourceQueryBuilder {
	return &sourceQueryBuilder{}
}

type sourceQueryBuilder struct {
	BaseQueryBuilder
}

func (b *sourceQueryBuilder) GUID(v string) SourceQueryBuilder {
	return b
}

func (b *sourceQueryBuilder) Created(v time.Time) SourceQueryBuilder {
	return b
}

func (b *sourceQueryBuilder) Updated(v time.Time) SourceQueryBuilder {
	return b
}

func (b *sourceQueryBuilder) Type(v string) SourceQueryBuilder {
	return b
}
