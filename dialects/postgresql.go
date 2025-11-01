package dialects

import (
	"fmt"

	"github.com/ductran999/dbkit/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreSQLDialect struct {
	cfg config.PostgreSQLConfig
}

// NewPostgreSQLDialect returns a new PostgreSQL dialect with config.
func NewPostgreSQLDialect(cfg config.PostgreSQLConfig) Dialect {
	return &postgreSQLDialect{cfg: cfg}
}

// Open opens a PostgreSQL database connection.
func (d *postgreSQLDialect) Open() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(d.buildDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(d.cfg.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(d.cfg.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(d.cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(d.cfg.ConnMaxIdleTime)

	return db, nil
}

func (d *postgreSQLDialect) buildDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		d.cfg.Host, d.cfg.Username, d.cfg.Password, d.cfg.Database,
		d.cfg.Port, d.cfg.SSLMode, d.cfg.TimeZone,
	)
}
