package repository

import "time"

func NewSourceQuery() SourcesQueryBuilder {
	return &sourcesQueryBuilder{}
}

type sourcesQueryBuilder struct {
	BaseQueryBuilder
}

func (b *sourcesQueryBuilder) GUID(v string) SourcesQueryBuilder {
	b.Add("sources_guid", Equal, v)
	return b
}

func (b *sourcesQueryBuilder) Created(v time.Time) SourcesQueryBuilder {
	b.Add("sources.created", Equal, v)
	return b
}

func (b *sourcesQueryBuilder) Type(v string) SourcesQueryBuilder {
	b.Add("type", Equal, v)
	return b
}
