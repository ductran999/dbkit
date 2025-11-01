package dbkit_test

import (
	"context"
	"testing"

	"github.com/ductran999/dbkit"
	"github.com/ductran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQLConnection(t *testing.T) {
	config := dbkit.PostgreSQLConfig{
		Config: dbkit.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := dbkit.NewPostgreSQLConnection(config)
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

func TestMySQLConnection(t *testing.T) {
	config := config.MySQLConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     3306,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := dbkit.NewMySQLConnection(config)
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

func TestClickHouseConnection(t *testing.T) {
	config := config.ClickHouseConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     9000,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := dbkit.NewClickHouseConnection(config)
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
