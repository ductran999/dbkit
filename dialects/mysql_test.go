package dialects_test

import (
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
	"github.com/stretchr/testify/require"
)

func TestMySQLDialect(t *testing.T) {
	tests := []struct {
		name        string
		config      config.MySQLConfig
		expectedErr string
	}{
		{
			name: "wrong connect information",
			config: config.MySQLConfig{
				Config: config.Config{
					Port:     4953,
					Host:     "localhost",
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "Asia/Ho_Chi_Minh",
				},
			},
			expectedErr: "failed to open MySQL connection",
		},
		{
			name: "valid config",
			config: config.MySQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     3306,
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "Asia/Ho_Chi_Minh",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := dialects.NewMySQLDialect(tc.config).Open()
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
