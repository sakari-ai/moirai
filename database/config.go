package database

import (
	"fmt"

	"github.com/sakari-ai/moirai/database/gorm"
)

type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Migrate  string
	Debug    bool
}

func (c Config) DSN() gorm.DSN {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Name,
		c.Password,
	)
	return gorm.DSN(dsn)
}
