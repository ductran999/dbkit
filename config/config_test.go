package config_test

import (
	"testing"
	"time"

	"github.com/ductran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name        string
		config      config.Config
		expectedErr error
	}{
		{
			name: "valid config",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Password: "testpass",
				Database: "testdb",
				TimeZone: "UTC",
			},
			expectedErr: nil,
		},
		{
			name: "valid config with minimal fields",
			config: config.Config{
				Host:     "127.0.0.1",
				Port:     3306,
				Username: "user",
				Database: "db",
			},
			expectedErr: nil,
		},
		{
			name: "missing host",
			config: config.Config{
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "empty host",
			config: config.Config{
				Host:     "",
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "whitespace only host",
			config: config.Config{
				Host:     "   ",
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "invalid port - zero",
			config: config.Config{
				Host:     "localhost",
				Port:     0,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "invalid port - negative",
			config: config.Config{
				Host:     "localhost",
				Port:     -1,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "invalid port - too high",
			config: config.Config{
				Host:     "localhost",
				Port:     65536,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "missing username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "empty username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "whitespace only username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "   ",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "missing database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "empty database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Database: "",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "whitespace only database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Database: "   ",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "valid port boundaries",
			config: config.Config{
				Host:     "localhost",
				Port:     1,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: nil,
		},
		{
			name: "valid port upper boundary",
			config: config.Config{
				Host:     "localhost",
				Port:     65535,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestPoolConfig_SetDefaults(t *testing.T) {
	tests := []struct {
		name     string
		input    config.PoolConfig
		expected config.PoolConfig
	}{
		{
			name:  "all zero values should use defaults",
			input: config.PoolConfig{},
			expected: config.PoolConfig{
				MaxIdleConnection: config.DefaultMaxIdleConnection,
				MaxOpenConnection: config.DefaultMaxOpenConnection,
				ConnMaxLifetime:   config.DefaultConnectionMaxLifetime,
				ConnMaxIdleTime:   config.DefaultConnectionMaxIdleTime,
			},
		},
		{
			name: "negative values should use defaults",
			input: config.PoolConfig{
				MaxIdleConnection: -5,
				MaxOpenConnection: -10,
				ConnMaxLifetime:   -time.Hour,
				ConnMaxIdleTime:   -30 * time.Minute,
			},
			expected: config.PoolConfig{
				MaxIdleConnection: config.DefaultMaxIdleConnection,
				MaxOpenConnection: config.DefaultMaxOpenConnection,
				ConnMaxLifetime:   config.DefaultConnectionMaxLifetime,
				ConnMaxIdleTime:   config.DefaultConnectionMaxIdleTime,
			},
		},
		{
			name: "MaxOpenConnection should be adjusted to match MaxIdleConnection when smaller",
			input: config.PoolConfig{
				MaxIdleConnection: 20,
				MaxOpenConnection: 10, // smaller than MaxIdleConnection
			},
			expected: config.PoolConfig{
				MaxIdleConnection: 20,
				MaxOpenConnection: 20, // should be adjusted to match MaxIdleConnection
				ConnMaxLifetime:   config.DefaultConnectionMaxLifetime,
				ConnMaxIdleTime:   config.DefaultConnectionMaxIdleTime,
			},
		},
		{
			name: "ConnMaxLifetime should be capped at 24 hours",
			input: config.PoolConfig{
				ConnMaxLifetime: 48 * time.Hour, // exceeds 24 hours
			},
			expected: config.PoolConfig{
				MaxIdleConnection: config.DefaultMaxIdleConnection,
				MaxOpenConnection: config.DefaultMaxOpenConnection,
				ConnMaxLifetime:   24 * time.Hour, // capped at 24 hours
				ConnMaxIdleTime:   config.DefaultConnectionMaxIdleTime,
			},
		},
		{
			name: "ConnMaxIdleTime should be capped at 1 hour",
			input: config.PoolConfig{
				ConnMaxIdleTime: 2 * time.Hour, // exceeds 1 hour
			},
			expected: config.PoolConfig{
				MaxIdleConnection: config.DefaultMaxIdleConnection,
				MaxOpenConnection: config.DefaultMaxOpenConnection,
				ConnMaxLifetime:   config.DefaultConnectionMaxLifetime,
				ConnMaxIdleTime:   time.Hour, // capped at 1 hour
			},
		},
		{
			name: "valid values within limits should remain unchanged",
			input: config.PoolConfig{
				MaxIdleConnection: 15,
				MaxOpenConnection: 25,
				ConnMaxLifetime:   12 * time.Hour,
				ConnMaxIdleTime:   30 * time.Minute,
			},
			expected: config.PoolConfig{
				MaxIdleConnection: 15,
				MaxOpenConnection: 25,
				ConnMaxLifetime:   12 * time.Hour,
				ConnMaxIdleTime:   30 * time.Minute,
			},
		},
		{
			name: "boundary values - exactly at limits",
			input: config.PoolConfig{
				MaxIdleConnection: 1,
				MaxOpenConnection: 1,
				ConnMaxLifetime:   24 * time.Hour, // exactly at limit
				ConnMaxIdleTime:   time.Hour,      // exactly at limit
			},
			expected: config.PoolConfig{
				MaxIdleConnection: 1,
				MaxOpenConnection: 1,
				ConnMaxLifetime:   24 * time.Hour,
				ConnMaxIdleTime:   time.Hour,
			},
		},
		{
			name: "mixed scenario - some defaults, some adjustments, some capping",
			input: config.PoolConfig{
				MaxIdleConnection: 0,              // should use default
				MaxOpenConnection: 5,              // valid but will be adjusted if less than default MaxIdle
				ConnMaxLifetime:   48 * time.Hour, // should be capped
				ConnMaxIdleTime:   0,              // should use default
			},
			expected: config.PoolConfig{
				MaxIdleConnection: config.DefaultMaxIdleConnection,
				MaxOpenConnection: max(config.DefaultMaxIdleConnection, 5), // adjusted if needed
				ConnMaxLifetime:   24 * time.Hour,                          // capped
				ConnMaxIdleTime:   config.DefaultConnectionMaxIdleTime,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := tt.input
			pc.SetDefaults()

			if pc.MaxIdleConnection != tt.expected.MaxIdleConnection {
				t.Errorf("MaxIdleConnection = %d, expected %d", pc.MaxIdleConnection, tt.expected.MaxIdleConnection)
			}
			if pc.MaxOpenConnection != tt.expected.MaxOpenConnection {
				t.Errorf("MaxOpenConnection = %d, expected %d", pc.MaxOpenConnection, tt.expected.MaxOpenConnection)
			}
			if pc.ConnMaxLifetime != tt.expected.ConnMaxLifetime {
				t.Errorf("ConnMaxLifetime = %v, expected %v", pc.ConnMaxLifetime, tt.expected.ConnMaxLifetime)
			}
			if pc.ConnMaxIdleTime != tt.expected.ConnMaxIdleTime {
				t.Errorf("ConnMaxIdleTime = %v, expected %v", pc.ConnMaxIdleTime, tt.expected.ConnMaxIdleTime)
			}
		})
	}
}

func TestPoolConfig_SetDefaults_MaxOpenConnectionAdjustment(t *testing.T) {
	// Specific test for the MaxOpenConnection adjustment logic
	testCases := []struct {
		name            string
		maxIdle         int
		maxOpen         int
		expectedMaxOpen int
	}{
		{"MaxOpen > MaxIdle - no adjustment", 10, 20, 20},
		{"MaxOpen = MaxIdle - no adjustment", 10, 10, 10},
		{"MaxOpen < MaxIdle - should adjust", 20, 10, 20},
		{"MaxOpen = 0, MaxIdle > 0 - should use default then adjust", 15, 0, max(config.DefaultMaxOpenConnection, 15)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pc := config.PoolConfig{
				MaxIdleConnection: tc.maxIdle,
				MaxOpenConnection: tc.maxOpen,
			}
			pc.SetDefaults()

			if pc.MaxOpenConnection != tc.expectedMaxOpen {
				t.Errorf("Expected MaxOpenConnection to be %d, got %d", tc.expectedMaxOpen, pc.MaxOpenConnection)
			}
		})
	}
}

func TestPoolConfig_SetDefaults_TimeCapping(t *testing.T) {
	// Specific tests for time value capping
	testCases := []struct {
		name             string
		inputLifetime    time.Duration
		inputIdleTime    time.Duration
		expectedLifetime time.Duration
		expectedIdleTime time.Duration
	}{
		{
			name:             "Both times exceed limits",
			inputLifetime:    48 * time.Hour,
			inputIdleTime:    2 * time.Hour,
			expectedLifetime: 24 * time.Hour,
			expectedIdleTime: time.Hour,
		},
		{
			name:             "Lifetime at boundary",
			inputLifetime:    24*time.Hour + time.Nanosecond,
			inputIdleTime:    30 * time.Minute,
			expectedLifetime: 24 * time.Hour,
			expectedIdleTime: 30 * time.Minute,
		},
		{
			name:             "IdleTime at boundary",
			inputLifetime:    12 * time.Hour,
			inputIdleTime:    time.Hour + time.Nanosecond,
			expectedLifetime: 12 * time.Hour,
			expectedIdleTime: time.Hour,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pc := config.PoolConfig{
				MaxIdleConnection: 10, // Set valid values to avoid defaults
				MaxOpenConnection: 20,
				ConnMaxLifetime:   tc.inputLifetime,
				ConnMaxIdleTime:   tc.inputIdleTime,
			}
			pc.SetDefaults()

			if pc.ConnMaxLifetime != tc.expectedLifetime {
				t.Errorf("Expected ConnMaxLifetime to be %v, got %v", tc.expectedLifetime, pc.ConnMaxLifetime)
			}
			if pc.ConnMaxIdleTime != tc.expectedIdleTime {
				t.Errorf("Expected ConnMaxIdleTime to be %v, got %v", tc.expectedIdleTime, pc.ConnMaxIdleTime)
			}
		})
	}
}
