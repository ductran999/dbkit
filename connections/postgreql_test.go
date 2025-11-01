package connections_test

import (
	"context"
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/connections"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQLConnection(t *testing.T) {
	pgConf := config.PostgreSQLConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := connections.NewPostgreSQLConnection(pgConf)
	require.NoError(t, err)

	// Test Ping to DB
	err = conn.Ping(context.Background())
	require.NoError(t, err)

	// verify db instance
	db := conn.DB()
	require.NotNil(t, db)

	conn.Close()

	// Test Ping to DB
	err = conn.Ping(context.Background())
	require.ErrorContains(t, err, "database is closed")
}

func TestPostgreSQLConnectionFailed(t *testing.T) {
	pgConf := config.PostgreSQLConfig{
		Config: config.Config{
			Port:     5433,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
		SSLMode: config.PgSSLDisable, // or whatever the correct SSLMode constant is
	}

	// test config missing host
	_, err := connections.NewPostgreSQLConnection(pgConf)
	require.ErrorIs(t, err, config.ErrMissingHost)

	// Test connection failed cause wrong port
	pgConf.Host = "localhost"
	conn, err := connections.NewPostgreSQLConnection(pgConf)
	require.ErrorContains(t, err, "failed to open PostgreSQL connection")
	require.Nil(t, conn)
}
