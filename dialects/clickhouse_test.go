package dialects_test

import (
	"testing"

	"github.com/ductran999/dbkit/config"
	"github.com/ductran999/dbkit/dialects"
	"github.com/stretchr/testify/require"
)

func TestClickHouseDialect(t *testing.T) {
	tests := []struct {
		name        string
		config      config.ClickHouseConfig
		expectedErr string
	}{
		{
			name: "wrong connect information",
			config: config.ClickHouseConfig{
				Config: config.Config{
					Port:     4953,
					Host:     "localhost",
					Username: "test",
					Password: "test",
					Database: "dbkit_test",
					TimeZone: "Asia/Ho_Chi_Minh",
				},
			},
			expectedErr: "failed to open ClickHouse connection",
		},
		{
			name: "valid config",
			config: config.ClickHouseConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     9000,
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
			_, err := dialects.NewClickHouseDialect(tc.config).Open()
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
