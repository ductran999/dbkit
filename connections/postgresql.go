package connections

import (
	"fmt"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
)

// NewPostgreSQLConnection initializes and returns a new PostgreSQL database connection.
func NewPostgreSQLConnection(cfg config.PostgreSQLConfig) (*connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := dialects.NewPostgreSQLDialect(cfg).Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	// Return the fully initialized connection.
	return &connection{db: db}, nil
}
