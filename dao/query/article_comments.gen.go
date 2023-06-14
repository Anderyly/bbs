// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bbs/dao/model"
)

func newArticleComment(db *gorm.DB, opts ...gen.DOOption) articleComment {
	_articleComment := articleComment{}

	_articleComment.articleCommentDo.UseDB(db, opts...)
	_articleComment.articleCommentDo.UseModel(&model.ArticleComment{})

	tableName := _articleComment.articleCommentDo.TableName()
	_articleComment.ALL = field.NewAsterisk(tableName)
	_articleComment.Id = field.NewInt(tableName, "id")
	_articleComment.CreatedAt = field.NewInt64(tableName, "created_at")
	_articleComment.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_articleComment.DeletedAt = field.NewInt64(tableName, "deleted_at")
	_articleComment.UserId = field.NewInt(tableName, "user_id")
	_articleComment.ArticleId = field.NewInt(tableName, "article_id")
	_articleComment.ReplyId = field.NewInt(tableName, "reply_id")
	_articleComment.Content = field.NewString(tableName, "content")

	_articleComment.fillFieldMap()

	return _articleComment
}

type articleComment struct {
	articleCommentDo articleCommentDo

	ALL       field.Asterisk
	Id        field.Int
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Int64
	UserId    field.Int
	ArticleId field.Int
	ReplyId   field.Int
	Content   field.String

	fieldMap map[string]field.Expr
}

func (a articleComment) Table(newTableName string) *articleComment {
	a.articleCommentDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleComment) As(alias string) *articleComment {
	a.articleCommentDo.DO = *(a.articleCommentDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleComment) updateTableName(table string) *articleComment {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewInt(table, "id")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.UpdatedAt = field.NewInt64(table, "updated_at")
	a.DeletedAt = field.NewInt64(table, "deleted_at")
	a.UserId = field.NewInt(table, "user_id")
	a.ArticleId = field.NewInt(table, "article_id")
	a.ReplyId = field.NewInt(table, "reply_id")
	a.Content = field.NewString(table, "content")

	a.fillFieldMap()

	return a
}

func (a *articleComment) WithContext(ctx context.Context) *articleCommentDo {
	return a.articleCommentDo.WithContext(ctx)
}

func (a articleComment) TableName() string { return a.articleCommentDo.TableName() }

func (a articleComment) Alias() string { return a.articleCommentDo.Alias() }

func (a *articleComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleComment) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.Id
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["user_id"] = a.UserId
	a.fieldMap["article_id"] = a.ArticleId
	a.fieldMap["reply_id"] = a.ReplyId
	a.fieldMap["content"] = a.Content
}

func (a articleComment) clone(db *gorm.DB) articleComment {
	a.articleCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleComment) replaceDB(db *gorm.DB) articleComment {
	a.articleCommentDo.ReplaceDB(db)
	return a
}

type articleCommentDo struct{ gen.DO }

func (a articleCommentDo) Debug() *articleCommentDo {
	return a.withDO(a.DO.Debug())
}

func (a articleCommentDo) WithContext(ctx context.Context) *articleCommentDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleCommentDo) ReadDB() *articleCommentDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleCommentDo) WriteDB() *articleCommentDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleCommentDo) Session(config *gorm.Session) *articleCommentDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleCommentDo) Clauses(conds ...clause.Expression) *articleCommentDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleCommentDo) Returning(value interface{}, columns ...string) *articleCommentDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleCommentDo) Not(conds ...gen.Condition) *articleCommentDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleCommentDo) Or(conds ...gen.Condition) *articleCommentDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleCommentDo) Select(conds ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleCommentDo) Where(conds ...gen.Condition) *articleCommentDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *articleCommentDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a articleCommentDo) Order(conds ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleCommentDo) Distinct(cols ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleCommentDo) Omit(cols ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleCommentDo) Join(table schema.Tabler, on ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleCommentDo) Group(cols ...field.Expr) *articleCommentDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleCommentDo) Having(conds ...gen.Condition) *articleCommentDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleCommentDo) Limit(limit int) *articleCommentDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleCommentDo) Offset(offset int) *articleCommentDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleCommentDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleCommentDo) Unscoped() *articleCommentDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleCommentDo) Create(values ...*model.ArticleComment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleCommentDo) CreateInBatches(values []*model.ArticleComment, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleCommentDo) Save(values ...*model.ArticleComment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleCommentDo) First() (*model.ArticleComment, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleComment), nil
	}
}

func (a articleCommentDo) Take() (*model.ArticleComment, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleComment), nil
	}
}

func (a articleCommentDo) Last() (*model.ArticleComment, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleComment), nil
	}
}

func (a articleCommentDo) Find() ([]*model.ArticleComment, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleComment), err
}

func (a articleCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleComment, err error) {
	buf := make([]*model.ArticleComment, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleCommentDo) FindInBatches(result *[]*model.ArticleComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleCommentDo) Attrs(attrs ...field.AssignExpr) *articleCommentDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleCommentDo) Assign(attrs ...field.AssignExpr) *articleCommentDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleCommentDo) Joins(fields ...field.RelationField) *articleCommentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleCommentDo) Preload(fields ...field.RelationField) *articleCommentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleCommentDo) FirstOrInit() (*model.ArticleComment, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleComment), nil
	}
}

func (a articleCommentDo) FirstOrCreate() (*model.ArticleComment, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleComment), nil
	}
}

func (a articleCommentDo) FindByPage(offset int, limit int) (result []*model.ArticleComment, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleCommentDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleCommentDo) Delete(models ...*model.ArticleComment) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleCommentDo) withDO(do gen.Dao) *articleCommentDo {
	a.DO = *do.(*gen.DO)
	return a
}