package connections_test

import (
	"context"
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/connections"
	"github.com/stretchr/testify/require"
)

func TestMySQLConnection(t *testing.T) {
	conf := config.MySQLConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     3306,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := connections.NewMySQLConnection(conf)
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

func TestMySQLConnectionFailed(t *testing.T) {
	conf := config.MySQLConfig{
		Config: config.Config{
			Port:     6432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	// test config missing host
	_, err := connections.NewMySQLConnection(conf)
	require.ErrorIs(t, err, config.ErrMissingHost)

	// Test connection failed cause wrong port
	conf.Host = "localhost"
	conn, err := connections.NewMySQLConnection(conf)
	require.ErrorContains(t, err, "failed to open MySQL connection")
	require.Nil(t, conn)
}
