package connections

import (
	"fmt"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
)

// NewClickHouseConnection initializes and returns a new ClickHouse database connection.
func NewClickHouseConnection(cfg config.ClickHouseConfig) (*connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := dialects.NewClickHouseDialect(cfg).Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open ClickHouse connection: %w", err)
	}

	// Return the fully initialized connection.
	return &connection{db: db}, nil
}
