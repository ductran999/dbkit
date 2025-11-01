package connections

import (
	"fmt"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
)

// NewMySQLConnection initializes and returns a new MySQL database connection.
func NewMySQLConnection(cfg config.MySQLConfig) (*connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := dialects.NewMySQLDialect(cfg).Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Return the fully initialized connection.
	return &connection{db: db}, nil
}
