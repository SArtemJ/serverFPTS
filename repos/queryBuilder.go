package repos

import (
	"fmt"
	"time"

	"go.uber.org/multierr"
)

type Operator string

const (
	// Compare operators
	Less         Operator = "$lt"
	LessOrEqual  Operator = "$lte"
	Great        Operator = "$gt"
	GreatOrEqual Operator = "$gte"
	Equal        Operator = "$eq"
	NotEqual     Operator = "$ne"
	Like         Operator = "$like"
	NotLike      Operator = "$notlike"
	Between      Operator = "$between"
	NotBetween   Operator = "$notbetween"
	In           Operator = "$in"
	Or           Operator = "$or"
	And          Operator = "$and"
	IsNull       Operator = "$null"
)

type E struct {
	Key   string
	Value interface{}
}

type QueryBuilder interface {
	//Join builder conditions with logical OR and add it as nested condition
	Or(b ...QueryBuilder) QueryBuilder
	//Join builder conditions with logical AND and add it as nested condition
	And(b ...QueryBuilder) QueryBuilder
	Add(name string, op Operator, v ...interface{}) QueryBuilder

	CreatedOp(op Operator, v ...interface{}) QueryBuilder
	UpdatedOp(op Operator, v ...interface{}) QueryBuilder

	OrderBy(name string, desc bool)

	Data() interface{}
	OrderData() []Order
	Reset()
	Error() error
}

type Order struct {
	Name string
	Desc bool
}

type BaseQueryBuilder struct {
	data  []E
	order []Order
	err   error
}

func (b *BaseQueryBuilder) addError(err error) {
	if err != nil {
		b.err = multierr.Append(b.err, err)
	}
}

func (b *BaseQueryBuilder) Error() error {
	return b.err
}

func (b *BaseQueryBuilder) Data() interface{} {
	return b.data
}

func (b *BaseQueryBuilder) OrderData() []Order {
	return b.order
}

func (b *BaseQueryBuilder) addGroup(isOr bool, builders ...QueryBuilder) QueryBuilder {
	for _, builder := range builders {
		if builder == nil || builder.Data() == nil {
			continue
		}
		op := And
		if isOr {
			op = Or
		}
		b.data = append(b.data, E{string(op), builder.Data()})
	}
	return b
}

func (b *BaseQueryBuilder) Or(builders ...QueryBuilder) QueryBuilder {
	return b.addGroup(true, builders...)
}

func (b *BaseQueryBuilder) And(builders ...QueryBuilder) QueryBuilder {
	return b.addGroup(false, builders...)
}

func (b *BaseQueryBuilder) CreatedOp(op Operator, v ...interface{}) QueryBuilder {
	return b.Add("created", op, v...)
}

func (b *BaseQueryBuilder) UpdatedOp(op Operator, v ...interface{}) QueryBuilder {
	return b.Add("updated", op, v...)
}

func subCondition(op Operator, v []interface{}) interface{} {
	if op == Equal {
		if len(v) == 1 {
			return v[0]
		}
		return v
	}
	var val interface{} = v
	if len(v) == 1 {
		val = v[0]
	}
	return E{string(op), val}
}

func (b *BaseQueryBuilder) Add(name string, op Operator, v ...interface{}) QueryBuilder {
	if len(v) == 0 {
		return b
	}
	// To prevent some errors when []interface{}, []string, []time.Time is passed without unpacking
	if len(v) == 1 {
		switch list := v[0].(type) {
		case []interface{}:
			v = list
		case []string:
			v = stringsToInterfaces(list)
		case []time.Time:
			v = timesToInterfaces(list)
		case []float64:
			v = floatToInterfaces(list)
		case []int64:
			v = intToInterfaces(list)
		}
	}

	if op == Equal && len(v) > 1 {
		op = In
	}
	if op == In && len(v) == 1 {
		op = Equal
	}

	switch op {
	case Between, NotBetween:
		if len(v) < 2 {
			b.addError(fmt.Errorf("\"Between(NotBetween)\" operation requare two values (field: %s)", name))
		}
	}
	b.data = append(b.data, E{name, subCondition(op, v)})
	return b
}

func (b *BaseQueryBuilder) OrderBy(name string, desc bool) {
	b.order = append(b.order, Order{Name: name, Desc: desc})
}

func (b *BaseQueryBuilder) Reset() {
	b.data = nil
	b.err = nil
}

func stringsToInterfaces(s []string) []interface{} {
	list := []interface{}{}
	for _, v := range s {
		list = append(list, v)
	}
	return list
}

func timesToInterfaces(t []time.Time) []interface{} {
	list := []interface{}{}
	for _, v := range t {
		list = append(list, v)
	}
	return list
}

func floatToInterfaces(f []float64) []interface{} {
	list := []interface{}{}
	for _, v := range f {
		list = append(list, v)
	}
	return list
}

func intToInterfaces(i []int64) []interface{} {
	list := []interface{}{}
	for _, v := range i {
		list = append(list, v)
	}
	return list
}
