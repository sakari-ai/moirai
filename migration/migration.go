package migration

type MigrateRollback interface {
	Migrate() error
	RollbackLast() error
}

type Migration struct {
	mr MigrateRollback
}

func New(opts ...Option) *Migration {
	m := &Migration{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Migration) Migrate() error {
	if m.mr != nil {
		return m.mr.Migrate()
	}
	return nil
}

func (m *Migration) RollbackLast() error {
	return m.mr.RollbackLast()
}
