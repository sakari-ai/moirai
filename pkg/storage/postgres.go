package storage

import (
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type PostgresStorage struct {
	DB database.DBEngine `inject:"database"`
}

func (p PostgresStorage) WriteSchema(schema *model.Schema) error {
	return p.DB.Model(&model.Schema{}).Create(schema).Error()
}

func (p PostgresStorage) GetSchema(id uuid.UUID) (*model.Schema, error) {
	m := new(model.Schema)
	m.ID = id
	res := p.DB.Model(&model.Schema{}).Where("id = ?", id).
		First(m)

	return m, res.Error()
}

func NewStorage() *PostgresStorage {
	return new(PostgresStorage)
}
