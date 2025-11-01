package dialects

import (
	"fmt"
	"net/url"

	"github.com/ductran999/dbkit/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDialect struct {
	cfg config.MySQLConfig
}

// NewMySQLDialect returns a new MySQL dialect with config.
func NewMySQLDialect(cfg config.MySQLConfig) Dialect {
	return &mysqlDialect{cfg: cfg}
}

// Open opens a MySQL database connection.
func (d *mysqlDialect) Open() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(d.buildDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
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

func (d *mysqlDialect) buildDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		url.QueryEscape(d.cfg.Username),
		url.QueryEscape(d.cfg.Password),
		d.cfg.Host,
		d.cfg.Port,
		d.cfg.Database,
		url.QueryEscape(d.cfg.TimeZone),
	)
}
