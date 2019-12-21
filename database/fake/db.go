package fake

import (
	"github.com/sakari-ai/moirai/database"
	"testing"
)

func WithDbOptions(options ...database.Option) DbTestingOption {
	return func(ops *dbTestingOptions) {
		ops.dbOptions = options
	}
}

func WithModels(models ...interface{}) DbTestingOption {
	return func(ops *dbTestingOptions) {
		ops.models = models
	}
}

type DbTestingOption func(ops *dbTestingOptions)
type dbTestingOptions struct {
	dbOptions []database.Option
	models    []interface{}
}

// PrepareForTesting prepares database for testing
// includes creating database engine and database schemas of this project.
func PrepareForTesting(t *testing.T, options ...DbTestingOption) database.DBEngine {
	var args = &dbTestingOptions{}
	for _, opt := range options {
		opt(args)
	}

	mdb, err := database.OpenInMemorySqlite(args.dbOptions...)
	if err != nil {
		t.Fatal(err)
	}

	if len(args.models) > 0 {
		err := mdb.AutoMigrate(args.models...).Error()
		if err != nil {
			t.Fatal(err)
		}
	}

	return mdb
}
