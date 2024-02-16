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

	"MapReduce/dal/model"
)

func newProduct(db *gorm.DB, opts ...gen.DOOption) product {
	_product := product{}

	_product.productDo.UseDB(db, opts...)
	_product.productDo.UseModel(&model.Product{})

	tableName := _product.productDo.TableName()
	_product.ALL = field.NewAsterisk(tableName)
	_product.ID = field.NewInt32(tableName, "ID")
	_product.Name = field.NewString(tableName, "Name")
	_product.Description = field.NewString(tableName, "Description")
	_product.Category = field.NewString(tableName, "Category")
	_product.Price = field.NewFloat64(tableName, "Price")
	_product.StockQuantity = field.NewInt32(tableName, "StockQuantity")
	_product.CountryOfManufacture = field.NewString(tableName, "CountryOfManufacture")
	_product.DateAdded = field.NewTime(tableName, "DateAdded")
	_product.LastUpdated = field.NewTime(tableName, "LastUpdated")
	_product.UnitsSold = field.NewInt32(tableName, "UnitsSold")
	_product.NumberOfReviews = field.NewInt32(tableName, "NumberOfReviews")
	_product.AverageRating = field.NewFloat64(tableName, "AverageRating")

	_product.fillFieldMap()

	return _product
}

type product struct {
	productDo productDo

	ALL                  field.Asterisk
	ID                   field.Int32   // 商品ID，主键，自增
	Name                 field.String  // 商品名称
	Description          field.String  // 商品描述
	Category             field.String  // 商品类别
	Price                field.Float64 // 商品价格，精确到小数点后两位
	StockQuantity        field.Int32   // 商品库存数量
	CountryOfManufacture field.String  // 商品生产国家
	DateAdded            field.Time    // 商品上架时间
	LastUpdated          field.Time    // 商品信息最后更新时间
	UnitsSold            field.Int32   // 商品销售数量
	NumberOfReviews      field.Int32   // 商品评价数量
	AverageRating        field.Float64 // 商品平均评分，精确到小数点后两位

	fieldMap map[string]field.Expr
}

func (p product) Table(newTableName string) *product {
	p.productDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p product) As(alias string) *product {
	p.productDo.DO = *(p.productDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *product) updateTableName(table string) *product {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "ID")
	p.Name = field.NewString(table, "Name")
	p.Description = field.NewString(table, "Description")
	p.Category = field.NewString(table, "Category")
	p.Price = field.NewFloat64(table, "Price")
	p.StockQuantity = field.NewInt32(table, "StockQuantity")
	p.CountryOfManufacture = field.NewString(table, "CountryOfManufacture")
	p.DateAdded = field.NewTime(table, "DateAdded")
	p.LastUpdated = field.NewTime(table, "LastUpdated")
	p.UnitsSold = field.NewInt32(table, "UnitsSold")
	p.NumberOfReviews = field.NewInt32(table, "NumberOfReviews")
	p.AverageRating = field.NewFloat64(table, "AverageRating")

	p.fillFieldMap()

	return p
}

func (p *product) WithContext(ctx context.Context) IProductDo { return p.productDo.WithContext(ctx) }

func (p product) TableName() string { return p.productDo.TableName() }

func (p product) Alias() string { return p.productDo.Alias() }

func (p product) Columns(cols ...field.Expr) gen.Columns { return p.productDo.Columns(cols...) }

func (p *product) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *product) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["ID"] = p.ID
	p.fieldMap["Name"] = p.Name
	p.fieldMap["Description"] = p.Description
	p.fieldMap["Category"] = p.Category
	p.fieldMap["Price"] = p.Price
	p.fieldMap["StockQuantity"] = p.StockQuantity
	p.fieldMap["CountryOfManufacture"] = p.CountryOfManufacture
	p.fieldMap["DateAdded"] = p.DateAdded
	p.fieldMap["LastUpdated"] = p.LastUpdated
	p.fieldMap["UnitsSold"] = p.UnitsSold
	p.fieldMap["NumberOfReviews"] = p.NumberOfReviews
	p.fieldMap["AverageRating"] = p.AverageRating
}

func (p product) clone(db *gorm.DB) product {
	p.productDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p product) replaceDB(db *gorm.DB) product {
	p.productDo.ReplaceDB(db)
	return p
}

type productDo struct{ gen.DO }

type IProductDo interface {
	gen.SubQuery
	Debug() IProductDo
	WithContext(ctx context.Context) IProductDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductDo
	WriteDB() IProductDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProductDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProductDo
	Not(conds ...gen.Condition) IProductDo
	Or(conds ...gen.Condition) IProductDo
	Select(conds ...field.Expr) IProductDo
	Where(conds ...gen.Condition) IProductDo
	Order(conds ...field.Expr) IProductDo
	Distinct(cols ...field.Expr) IProductDo
	Omit(cols ...field.Expr) IProductDo
	Join(table schema.Tabler, on ...field.Expr) IProductDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductDo
	Group(cols ...field.Expr) IProductDo
	Having(conds ...gen.Condition) IProductDo
	Limit(limit int) IProductDo
	Offset(offset int) IProductDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProductDo
	Unscoped() IProductDo
	Create(values ...*model.Product) error
	CreateInBatches(values []*model.Product, batchSize int) error
	Save(values ...*model.Product) error
	First() (*model.Product, error)
	Take() (*model.Product, error)
	Last() (*model.Product, error)
	Find() ([]*model.Product, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Product, err error)
	FindInBatches(result *[]*model.Product, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Product) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProductDo
	Assign(attrs ...field.AssignExpr) IProductDo
	Joins(fields ...field.RelationField) IProductDo
	Preload(fields ...field.RelationField) IProductDo
	FirstOrInit() (*model.Product, error)
	FirstOrCreate() (*model.Product, error)
	FindByPage(offset int, limit int) (result []*model.Product, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productDo) Debug() IProductDo {
	return p.withDO(p.DO.Debug())
}

func (p productDo) WithContext(ctx context.Context) IProductDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productDo) ReadDB() IProductDo {
	return p.Clauses(dbresolver.Read)
}

func (p productDo) WriteDB() IProductDo {
	return p.Clauses(dbresolver.Write)
}

func (p productDo) Session(config *gorm.Session) IProductDo {
	return p.withDO(p.DO.Session(config))
}

func (p productDo) Clauses(conds ...clause.Expression) IProductDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productDo) Returning(value interface{}, columns ...string) IProductDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productDo) Not(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productDo) Or(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productDo) Select(conds ...field.Expr) IProductDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productDo) Where(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productDo) Order(conds ...field.Expr) IProductDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productDo) Distinct(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productDo) Omit(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productDo) Join(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productDo) Group(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productDo) Having(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productDo) Limit(limit int) IProductDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productDo) Offset(offset int) IProductDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProductDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productDo) Unscoped() IProductDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productDo) Create(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productDo) CreateInBatches(values []*model.Product, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productDo) Save(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productDo) First() (*model.Product, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Take() (*model.Product, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Last() (*model.Product, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Find() ([]*model.Product, error) {
	result, err := p.DO.Find()
	return result.([]*model.Product), err
}

func (p productDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Product, err error) {
	buf := make([]*model.Product, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productDo) FindInBatches(result *[]*model.Product, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productDo) Attrs(attrs ...field.AssignExpr) IProductDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productDo) Assign(attrs ...field.AssignExpr) IProductDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productDo) Joins(fields ...field.RelationField) IProductDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productDo) Preload(fields ...field.RelationField) IProductDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productDo) FirstOrInit() (*model.Product, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FirstOrCreate() (*model.Product, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FindByPage(offset int, limit int) (result []*model.Product, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productDo) Delete(models ...*model.Product) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productDo) withDO(do gen.Dao) *productDo {
	p.DO = *do.(*gen.DO)
	return p
}
