package repos

import "time"

func NewUserQuery() UserQueryBuilder {
	return &userQueryBuilder{}
}

type userQueryBuilder struct {
	BaseQueryBuilder
}

func (b *userQueryBuilder) GUID(v string) UserQueryBuilder {
	return b
}

func (b *userQueryBuilder) Created(v time.Time) UserQueryBuilder {
	return b
}

func (b *userQueryBuilder) Updated(v time.Time) UserQueryBuilder {
	return b
}

func (b *userQueryBuilder) Email(v string) UserQueryBuilder {
	return b
}

func (b *userQueryBuilder) Gender(v string) UserQueryBuilder {
	return b
}
