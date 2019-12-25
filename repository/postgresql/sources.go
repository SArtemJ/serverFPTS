package postgresql

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/didi/gendry/scanner"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
)

type SourcesRepository struct {
	BaseRepository
}

func (repo *SourcesRepository) initSelect(b repository.QueryBuilder, offset *int64, limit *int64, count bool) (*goqu.SelectDataset, error) {
	expression, err := builderToExpression(b, SourcesTable)
	if err != nil {
		return nil, err
	}
	var q *goqu.SelectDataset
	q = newQuery().From(SourcesTable).Select(
		goqu.I(SourcesTable+".guid").As("guid"),
		goqu.I(SourcesTable+".created").As("created"),
		goqu.I(SourcesTable+".type").As("type"),
	).Where(expression...).Prepared(true)

	if !count {
		order := builderToOrderExpression(b)
		if len(order) != 0 {
			q = q.Order(order...)
		}
	}

	if limit != nil {
		q = q.Limit(uint(*limit))
	}
	if offset != nil {
		q = q.Offset(uint(*offset))
	}

	return q, nil
}

func (repo *SourcesRepository) Count(b repository.SourcesQueryBuilder) (int64, error) {
	q, err := repo.initSelect(b, nil, nil, true)
	if err != nil {
		return 0, err
	}
	return countFromQuery(repo.conn, SourcesTable, q)
}

func (repo *SourcesRepository) Collection(offset *int64, limit *int64) ([]*repository.SourcesModel, error) {
	return repo.Find(nil, offset, limit)
}

func (repo *SourcesRepository) Create(model *repository.SourcesModel) error {
	modelMap, err := sourceMapFromModel(repo.conn, model)
	if err != nil {
		return err
	}
	_, err = insertToTable(repo.conn, SourcesTable, modelMap, true)
	return err
}

func (repo *SourcesRepository) Update(guid string, model *repository.SourcesModel) (int64, error) {
	modelMap, err := sourceMapFromModel(repo.conn, model)
	if err != nil {
		return 0, err
	}
	return updateInTableByGUID(repo.conn, SourcesTable, guid, modelMap)
}

func (repo *SourcesRepository) Find(b repository.SourcesQueryBuilder, offset *int64, limit *int64) ([]*repository.SourcesModel, error) {
	q, err := repo.initSelect(b, offset, limit, false)
	if err != nil {
		return nil, err
	}
	sql, args, _ := q.ToSQL()

	rows, err := Query(repo.conn, sql, args...)
	if err != nil {
		return nil, err
	}

	var result []*repository.SourcesModel

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return nil, err
	}

	for _, dataMap := range dbResult {
		result = append(result, sourceModelFromMap(dataMap))
	}

	return result, nil
}

func NewSourcesRepo(conn Conn) *SourcesRepository {
	return &SourcesRepository{
		newBaseRepository(conn),
	}
}

func sourceMapFromModel(conn Conn, model *repository.SourcesModel) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	if !model.GUID.IsZero() {
		dataMap["guid"] = model.GUID.String
	}
	if !model.Created.IsZero() {
		dataMap["created"] = model.Created.Time
	}
	if !model.SourceType.IsZero() {
		dataMap["type"] = model.SourceType.String
	}
	return dataMap, nil
}

func sourceModelFromMap(data map[string]interface{}) *repository.SourcesModel {
	return &repository.SourcesModel{
		GUID:       null.StringFrom(cast.ToString(data["guid"])),
		BaseModel:  repository.BaseModel{Created: dbValueToNullTime(data["created"])},
		SourceType: null.StringFrom(cast.ToString(data["type"])),
	}
}
