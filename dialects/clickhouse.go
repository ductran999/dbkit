package dialects

import (
	"fmt"
	"net"
	"strconv"
	"time"

	std_ck "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ductran999/dbkit/config"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type clickHouseDialect struct {
	conf config.ClickHouseConfig
}

// NewClickHouseDialect returns a new ClickHouse dialect with config.
func NewClickHouseDialect(conf config.ClickHouseConfig) Dialect {
	return &clickHouseDialect{
		conf: conf,
	}
}

// Open opens a ClickHouse database connection.
func (d *clickHouseDialect) Open() (*gorm.DB, error) {

	sqlDB := std_ck.OpenDB(&std_ck.Options{
		Addr: []string{net.JoinHostPort(d.conf.Host, strconv.Itoa(d.conf.Port))},
		Auth: std_ck.Auth{
			Database: d.conf.Database,
			Username: d.conf.Username,
			Password: d.conf.Password,
		},
		Settings: std_ck.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Debug:       true,
	})

	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		Conn: sqlDB,
	}))

	if err != nil {
		return nil, fmt.Errorf("failed to open ClickHouse connection: %w", err)
	}

	return db, nil
}
