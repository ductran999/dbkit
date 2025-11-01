package config_test

import (
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestPostgresqlConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      config.PostgreSQLConfig
		expectedErr error
	}{
		{
			name: "valid config default disable ssl",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "UTC",
				},
			},
			expectedErr: nil,
		},
		{
			name: "valid config",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "test",
					Password: "test",
					Database: "testdb",
					TimeZone: "UTC",
				},
				SSLMode: config.PgSSLVerifyFull,
			},
			expectedErr: nil,
		},
		{
			name: "invalid ssl mode",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "UTC",
				},
				SSLMode: "unknown",
			},
			expectedErr: nil,
		},
		{
			name: "invalid base config",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Port:     5432,
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "UTC",
				},
			},
			expectedErr: config.ErrMissingHost,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
