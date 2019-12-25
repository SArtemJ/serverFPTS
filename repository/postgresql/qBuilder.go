package postgresql

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"strings"
)

func builderToOrderExpression(builder repository.QueryBuilder) (exps []exp.OrderedExpression) {
	exps = []exp.OrderedExpression{}
	if builder == nil {
		return
	}
	for _, o := range builder.OrderData() {
		if o.Desc {
			exps = append(exps, goqu.I(o.Name).Desc())
		} else {
			exps = append(exps, goqu.I(o.Name).Asc())
		}
	}
	return
}

func builderToExpression(builder repository.QueryBuilder, table ...string) ([]goqu.Expression, error) {
	if builder == nil {
		return []goqu.Expression{}, nil
	}
	if err := builder.Error(); err != nil {
		return nil, err
	}
	t := ""
	if len(table) > 0 {
		t = table[0]
	}
	return queryDataToGoquExpression(builder.Data(), t)
}

func queryDataToGoquExpression(data interface{}, table string) ([]goqu.Expression, error) {
	conditions := []goqu.Expression{}
	onCondition := func(name string, op repository.Operator, v []interface{}) error {
		castType := ""
		if parts := strings.Split(name, "::"); len(parts) >= 2 {
			name, castType = parts[0], parts[1]
		}

		switch name {
		case "users_guid":
			name = UsersTable + ".guid"
		case "transactions_guid":
			name = TransactionTable + ".guid"
		case "sources_guid":
			name = SourcesTable + ".guid"
		case "guid":
			if table != "" {
				name = table + ".guid"
			}
		}
		var e goqu.Expression
		ident := goqu.I(name)
		if castType == "" {
			switch op {
			case repository.Between:
				e = ident.Between(goqu.Range(v[0], v[1]))
			case repository.NotBetween:
				e = ident.NotBetween(goqu.Range(v[0], v[1]))
			case repository.In:
				e = ident.In(v...)
			default:
				for _, v := range v {
					switch op {
					case repository.Less:
						e = ident.Lt(v)
					case repository.LessOrEqual:
						e = ident.Lte(v)
					case repository.Great:
						e = ident.Gt(v)
					case repository.GreatOrEqual:
						e = ident.Gte(v)
					case repository.Equal:
						e = ident.Eq(v)
					case repository.NotEqual:
						e = ident.Neq(v)
					case repository.Like:
						e = ident.Like(v)
					case repository.NotLike:
						e = ident.NotLike(v)
					case repository.IsNull:
						e = ident.Eq(nil)
					}
					conditions = append(conditions, e)
				}
				return nil
			}
		} else {
			switch op {
			case repository.Between:
				e = ident.Cast(castType).Between(goqu.Range(goqu.L("?::"+castType, v[0]), goqu.L("?::"+castType, v[1])))
			case repository.NotBetween:
				e = ident.Cast(castType).NotBetween(goqu.Range(goqu.L("?::"+castType, v[0]), goqu.L("?::"+castType, v[1])))
			case repository.In:
				e = ident.In(v...)
			default:
				for _, v := range v {
					switch op {
					case repository.Less:
						e = ident.Cast(castType).Lt(goqu.L("?::"+castType, v))
					case repository.LessOrEqual:
						e = ident.Cast(castType).Lte(goqu.L("?::"+castType, v))
					case repository.Great:
						e = ident.Cast(castType).Gt(goqu.L("?::"+castType, v))
					case repository.GreatOrEqual:
						e = ident.Cast(castType).Gte(goqu.L("?::"+castType, v))
					case repository.Equal:
						e = ident.Cast(castType).Eq(goqu.L("?::"+castType, v))
					case repository.NotEqual:
						e = ident.Cast(castType).Neq(goqu.L("?::"+castType, v))
					case repository.Like:
						e = ident.Cast(castType).Like(goqu.L("?::"+castType, v))
					case repository.NotLike:
						e = ident.Cast(castType).NotLike(goqu.L("?::"+castType, v))
					case repository.IsNull:
						e = ident.Cast(castType).Eq(nil)
					}
					conditions = append(conditions, e)
				}
				return nil
			}
		}
		conditions = append(conditions, e)
		return nil
	}
	onGroup := func(groupType repository.Operator, data interface{}) error {
		c, err := queryDataToGoquExpression(data, table)
		if err != nil {
			return err
		}
		switch groupType {
		case repository.And:
			conditions = append(conditions, goqu.And(c...))
		case repository.Or:
			conditions = append(conditions, goqu.Or(c...))
		}
		return nil
	}
	err := repository.TraverseQueryData(data, onCondition, onGroup)
	if err != nil {
		return nil, err
	}

	return conditions, nil
}
