// Package dbkit provides a unified database abstraction layer for Go applications.
// It supports multiple database drivers with consistent interfaces and advanced features.
package dbkit

import (
	"context"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/connections"
	"gorm.io/gorm"
)

// Connection represents a database connection with advanced features
type Connection interface {
	// Core database operations
	DB() *gorm.DB
	Close() error
	Ping(ctx context.Context) error
}

type Config = config.Config
type PostgreSQLConfig = config.PostgreSQLConfig
type MySQLConfig = config.MySQLConfig
type ClickHouseConfig = config.ClickHouseConfig

var (
	ErrPostgresqlSSLMode = config.ErrPostgresqlSSLMode
	ErrMissingHost       = config.ErrMissingHost
	ErrInvalidPort       = config.ErrInvalidPort
	ErrMissingUsername   = config.ErrMissingUsername
	ErrMissingDatabase   = config.ErrMissingDatabase
)

func NewPostgreSQLConnection(cfg PostgreSQLConfig) (Connection, error) {
	return connections.NewPostgreSQLConnection(cfg)
}

func NewMySQLConnection(cfg MySQLConfig) (Connection, error) {
	return connections.NewMySQLConnection(cfg)
}

func NewClickHouseConnection(cfg ClickHouseConfig) (Connection, error) {
	return connections.NewClickHouseConnection(cfg)
}
