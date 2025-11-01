package dialects_test

import (
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQLDialect(t *testing.T) {
	tests := []struct {
		name        string
		config      config.PostgreSQLConfig
		expectedErr string
	}{
		{
			name: "wrong connect information",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Port:     4953,
					Host:     "localhost",
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "Asia/Ho_Chi_Minh",
				},
				SSLMode: config.PgSSLDisable,
			},
			expectedErr: "failed to open PostgreSQL connection",
		},
		{
			name: "valid config",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "Asia/Ho_Chi_Minh",
				},
				SSLMode: "disable",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := dialects.NewPostgreSQLDialect(tc.config).Open()
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
