package repository

import "time"

func NewUserQuery() UsersQueryBuilder {
	return &usersQueryBuilder{}
}

type usersQueryBuilder struct {
	BaseQueryBuilder
}

func (b *usersQueryBuilder) GUID(v string) UsersQueryBuilder {
	b.Add("users_guid", Equal, v)
	return b
}

func (b *usersQueryBuilder) Created(v time.Time) UsersQueryBuilder {
	b.Add("users.created", Equal, v)
	return b
}

func (b *usersQueryBuilder) Email(v string) UsersQueryBuilder {
	b.Add("email", Equal, v)
	return b
}
