package connections_test

import (
	"context"
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/connections"
	"github.com/stretchr/testify/require"
)

func TestClickHouseConnection(t *testing.T) {
	chConf := config.ClickHouseConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     9000,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := connections.NewClickHouseConnection(chConf)
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

func TestClickHouseConnectionFailed(t *testing.T) {
	chConf := config.ClickHouseConfig{
		Config: config.Config{
			Port:     6432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	// test config missing host
	_, err := connections.NewClickHouseConnection(chConf)
	require.ErrorIs(t, err, config.ErrMissingHost)

	// Test connection failed cause wrong port
	chConf.Host = "localhost"
	conn, err := connections.NewClickHouseConnection(chConf)
	require.ErrorContains(t, err, "failed to open ClickHouse connection")
	require.Nil(t, conn)
}
