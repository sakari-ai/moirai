package database

import (
	"database/sql"

	lru "github.com/hashicorp/golang-lru"

	"github.com/sakari-ai/moirai/database/gorm"
)

// DBEngine is an interface which DB implements. By then all test are well organized
type DBEngine interface {
	Close() error
	DB() *sql.DB
	New() DBEngine
	LogMode(enable bool) DBEngine
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) DBEngine
	Filter(query interface{}, by interface{}) DBEngine
	Or(query interface{}, args ...interface{}) DBEngine
	Not(query interface{}, args ...interface{}) DBEngine
	Limit(value int) DBEngine
	Offset(value int) DBEngine
	Order(value string, reorder ...bool) DBEngine
	Select(query interface{}, args ...interface{}) DBEngine
	Omit(columns ...string) DBEngine
	Group(query string) DBEngine
	Having(query string, values ...interface{}) DBEngine
	Joins(query string, args ...interface{}) DBEngine
	Unscoped() DBEngine
	Attrs(attrs ...interface{}) DBEngine
	Assign(attrs ...interface{}) DBEngine
	First(out interface{}, where ...interface{}) DBEngine
	Last(out interface{}, where ...interface{}) DBEngine
	Find(out interface{}, where ...interface{}) DBEngine
	Scan(dest interface{}) DBEngine
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) DBEngine
	Count(value interface{}) DBEngine
	Related(value interface{}, foreignKeys ...string) DBEngine
	FirstOrInit(out interface{}, where ...interface{}) DBEngine
	FirstOrCreate(out interface{}, where ...interface{}) DBEngine
	Update(attrs ...interface{}) DBEngine
	Updates(values interface{}, ignoreProtectedAttrs ...bool) DBEngine
	UpdateColumn(attrs ...interface{}) DBEngine
	UpdateColumns(values interface{}) DBEngine
	Save(value interface{}) DBEngine
	Create(value interface{}) DBEngine
	Delete(value interface{}, where ...interface{}) DBEngine
	Raw(sql string, values ...interface{}) DBEngine
	Exec(sql string, values ...interface{}) DBEngine
	Model(value interface{}) DBEngine
	Table(name string) DBEngine
	Debug() DBEngine
	Begin() DBEngine
	Commit() DBEngine
	Rollback() DBEngine
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(values ...interface{}) DBEngine
	DropTable(values ...interface{}) DBEngine
	DropTableIfExists(values ...interface{}) DBEngine
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) DBEngine
	ModifyColumn(column string, typ string) DBEngine
	DropColumn(column string) DBEngine
	AddIndex(indexName string, column ...string) DBEngine
	AddUniqueIndex(indexName string, column ...string) DBEngine
	RemoveIndex(indexName string) DBEngine
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) DBEngine
	Preload(column string, conditions ...interface{}) DBEngine
	Set(name string, value interface{}) DBEngine
	InstantSet(name string, value interface{}) DBEngine
	Get(name string) (value interface{}, ok bool)
	Association(column string) Association
	AddError(err error) error
	GetErrors() (errors []error)

	// extra
	Error() error
	RowsAffected() int64
	Cache() Cache
}

type DB struct {
	db      *gorm.DB
	SkCache Cache `inject:"innete_cache"`
}

// Open is a drop-in replacement for Open()
func Open(cfg Config) (DBEngine, error) {
	db, err := gorm.Open(cfg.DSN())
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

type options struct {
	cache Cache
}
type Option func(o *options)

func WithCache(cache Cache) Option {
	return func(o *options) {
		o.cache = cache
	}
}

// Creates an sqlite memory db which will be database.DBEngine
// If WithCache is not provided, lru cache will be created as default.
func OpenInMemorySqlite(opts ...Option) (DBEngine, error) {
	db, err := gorm.OpenDialects("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	args := &options{}
	for _, opt := range opts {
		opt(args)
	}

	if args.cache == nil {
		// init default cache
		lruCache, _ := lru.NewARC(10)
		args.cache = lruCache
	}

	return &DB{
		db:      db,
		SkCache: args.cache,
	}, nil
}

// wrap wraps gorm.DB in an interface
func wrap(db *gorm.DB, cache Cache) DBEngine {
	return &DB{
		db:      db,
		SkCache: cache,
	}
}

func (d *DB) Cache() Cache {
	return d.SkCache
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) DB() *sql.DB {
	return d.db.DB()
}

func (d *DB) New() DBEngine {
	return wrap(d.db.New(), d.Cache())
}

func (d *DB) LogMode(enable bool) DBEngine {
	return wrap(d.db.LogMode(enable), d.Cache())
}

func (d *DB) SingularTable(enable bool) {
	d.db.SingularTable(enable)
}

func (d *DB) Where(query interface{}, args ...interface{}) DBEngine {
	return wrap(d.db.Where(query, args...), d.Cache())
}

func (d *DB) Filter(query interface{}, arg interface{}) DBEngine {
	if arg != nil && arg != "" {
		return wrap(d.db.Where(query, arg), d.Cache())
	}
	return d
}

func (d *DB) Or(query interface{}, args ...interface{}) DBEngine {
	return wrap(d.db.Or(query, args...), d.Cache())
}

func (d *DB) Not(query interface{}, args ...interface{}) DBEngine {
	return wrap(d.db.Not(query, args...), d.Cache())
}

func (d *DB) Limit(value int) DBEngine {
	if value == 0 {
		value = 10
	}
	return wrap(d.db.Limit(value), d.Cache())
}

func (d *DB) Offset(value int) DBEngine {
	return wrap(d.db.Offset(value), d.Cache())
}

func (d *DB) Order(value string, reorder ...bool) DBEngine {
	return wrap(d.db.Order(value, reorder...), d.Cache())
}

func (d *DB) Select(query interface{}, args ...interface{}) DBEngine {
	return wrap(d.db.Select(query, args...), d.Cache())
}

func (d *DB) Omit(columns ...string) DBEngine {
	return wrap(d.db.Omit(columns...), d.Cache())
}

func (d *DB) Group(query string) DBEngine {
	return wrap(d.db.Group(query), d.Cache())
}

func (d *DB) Having(query string, values ...interface{}) DBEngine {
	return wrap(d.db.Having(query, values...), d.Cache())
}

func (d *DB) Joins(query string, args ...interface{}) DBEngine {
	return wrap(d.db.Joins(query, args...), d.Cache())
}

func (d *DB) Unscoped() DBEngine {
	return wrap(d.db.Unscoped(), d.Cache())
}

func (d *DB) Attrs(attrs ...interface{}) DBEngine {
	return wrap(d.db.Attrs(attrs...), d.Cache())
}

func (d *DB) Assign(attrs ...interface{}) DBEngine {
	return wrap(d.db.Assign(attrs...), d.Cache())
}

func (d *DB) First(out interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.First(out, where...), d.Cache())
}

func (d *DB) Last(out interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.Last(out, where...), d.Cache())
}

func (d *DB) Find(out interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.Find(out, where...), d.Cache())
}

func (d *DB) Scan(dest interface{}) DBEngine {
	return wrap(d.db.Scan(dest), d.Cache())
}

func (d *DB) Row() *sql.Row {
	return d.db.Row()
}

func (d *DB) Rows() (*sql.Rows, error) {
	return d.db.Rows()
}

func (d *DB) ScanRows(rows *sql.Rows, result interface{}) error {
	return d.db.ScanRows(rows, result)
}

func (d *DB) Pluck(column string, value interface{}) DBEngine {
	return wrap(d.db.Pluck(column, value), d.Cache())
}

func (d *DB) Count(value interface{}) DBEngine {
	return wrap(d.db.Count(value), d.Cache())
}

func (d *DB) Related(value interface{}, foreignKeys ...string) DBEngine {
	return wrap(d.db.Related(value, foreignKeys...), d.Cache())
}

func (d *DB) FirstOrInit(out interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.FirstOrInit(out, where...), d.Cache())
}

func (d *DB) FirstOrCreate(out interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.FirstOrCreate(out, where...), d.Cache())
}

func (d *DB) Update(attrs ...interface{}) DBEngine {
	return wrap(d.db.Update(attrs...), d.Cache())
}

func (d *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) DBEngine {
	return wrap(d.db.Updates(values, ignoreProtectedAttrs...), d.Cache())
}

func (d *DB) UpdateColumn(attrs ...interface{}) DBEngine {
	return wrap(d.db.UpdateColumn(attrs...), d.Cache())
}

func (d *DB) UpdateColumns(values interface{}) DBEngine {
	return wrap(d.db.UpdateColumns(values), d.Cache())
}

func (d *DB) Save(value interface{}) DBEngine {
	return wrap(d.db.Save(value), d.Cache())
}

func (d *DB) Create(value interface{}) DBEngine {
	return wrap(d.db.Create(value), d.Cache())
}

func (d *DB) Delete(value interface{}, where ...interface{}) DBEngine {
	return wrap(d.db.Delete(value, where...), d.Cache())
}

func (d *DB) Raw(sql string, values ...interface{}) DBEngine {
	return wrap(d.db.Raw(sql, values...), d.Cache())
}

func (d *DB) Exec(sql string, values ...interface{}) DBEngine {
	return wrap(d.db.Exec(sql, values...), d.Cache())
}

func (d *DB) Model(value interface{}) DBEngine {
	return wrap(d.db.Model(value), d.Cache())
}

func (d *DB) Table(name string) DBEngine {
	return wrap(d.db.Table(name), d.Cache())
}

func (d *DB) Debug() DBEngine {
	return wrap(d.db.Debug(), d.Cache())
}

func (d *DB) Begin() DBEngine {
	return wrap(d.db.Begin(), d.Cache())
}

func (d *DB) Commit() DBEngine {
	return wrap(d.db.Commit(), d.Cache())
}

func (d *DB) Rollback() DBEngine {
	return wrap(d.db.Rollback(), d.Cache())
}

func (d *DB) NewRecord(value interface{}) bool {
	return d.db.NewRecord(value)
}

func (d *DB) RecordNotFound() bool {
	return d.db.RecordNotFound()
}

func (d *DB) CreateTable(values ...interface{}) DBEngine {
	return wrap(d.db.CreateTable(values...), d.Cache())
}

func (d *DB) DropTable(values ...interface{}) DBEngine {
	return wrap(d.db.DropTable(values...), d.Cache())
}

func (d *DB) DropTableIfExists(values ...interface{}) DBEngine {
	return wrap(d.db.DropTableIfExists(values...), d.Cache())
}

func (d *DB) HasTable(value interface{}) bool {
	return d.db.HasTable(value)
}

func (d *DB) AutoMigrate(values ...interface{}) DBEngine {
	return wrap(d.db.AutoMigrate(values...), d.Cache())
}

func (d *DB) ModifyColumn(column string, typ string) DBEngine {
	return wrap(d.db.ModifyColumn(column, typ), d.Cache())
}

func (d *DB) DropColumn(column string) DBEngine {
	return wrap(d.db.DropColumn(column), d.Cache())
}

func (d *DB) AddIndex(indexName string, columns ...string) DBEngine {
	return wrap(d.db.AddIndex(indexName, columns...), d.Cache())
}

func (d *DB) AddUniqueIndex(indexName string, columns ...string) DBEngine {
	return wrap(d.db.AddUniqueIndex(indexName, columns...), d.Cache())
}

func (d *DB) RemoveIndex(indexName string) DBEngine {
	return wrap(d.db.RemoveIndex(indexName), d.Cache())
}

func (d *DB) Preload(column string, conditions ...interface{}) DBEngine {
	return wrap(d.db.Preload(column, conditions...), d.Cache())
}

func (d *DB) Set(name string, value interface{}) DBEngine {
	return wrap(d.db.Set(name, value), d.Cache())
}

func (d *DB) InstantSet(name string, value interface{}) DBEngine {
	return wrap(d.db.InstantSet(name, value), d.Cache())
}

func (d *DB) Get(name string) (interface{}, bool) {
	return d.db.Get(name)
}

func (d *DB) Association(column string) Association {
	return newAssociation(d.db.Association(column))
}

func (d *DB) AddForeignKey(field string, dest string, onDelete string, onUpdate string) DBEngine {
	return wrap(d.db.AddForeignKey(field, dest, onDelete, onUpdate), d.Cache())
}

func (d *DB) AddError(err error) error {
	return d.db.AddError(err)
}

func (d *DB) GetErrors() (errors []error) {
	return d.db.GetErrors()
}

func (d *DB) RowsAffected() int64 {
	return d.db.RowsAffected
}

func (d *DB) Error() error {
	return d.db.Error
}
