package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwtService_GenerateToken(t *testing.T) {
	tests := []struct {
		name           string
		userId         string
		jwtSecret      string
		expireTimeSec  int
		expectedError  bool
		validateResult bool
	}{
		{
			name:           "Success - Generate valid token",
			userId:         "test-user-123",
			jwtSecret:      "test-secret",
			expireTimeSec:  3600,
			expectedError:  false,
			validateResult: true,
		},
		{
			name:           "Success - Generate token with empty user ID",
			userId:         "",
			jwtSecret:      "test-secret",
			expireTimeSec:  3600,
			expectedError:  false,
			validateResult: true,
		},
		{
			name:           "Success - Generate token with short expiry",
			userId:         "test-user-123",
			jwtSecret:      "test-secret",
			expireTimeSec:  1,
			expectedError:  false,
			validateResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewJwtService(tt.jwtSecret, tt.expireTimeSec)
			token, err := service.GenerateToken(tt.userId)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)

				// Validate the generated token
				if tt.validateResult {
					jwtToken, err := service.ValidateToken(token)
					assert.NoError(t, err)
					assert.NotNil(t, jwtToken)
					assert.Equal(t, tt.userId, jwtToken.UserId)
					assert.True(t, jwtToken.ExpiredAt.After(time.Now()))
				}
			}
		})
	}
}

func TestJwtService_ValidateToken(t *testing.T) {
	tests := []struct {
		name          string
		token         string
		jwtSecret     string
		expireTimeSec int
		expectedError bool
	}{
		{
			name:          "Success - Validate valid token",
			token:         "", // Will be generated in the test
			jwtSecret:     "test-secret",
			expireTimeSec: 3600,
			expectedError: false,
		},
		{
			name:          "Failure - Invalid token format",
			token:         "invalid-token",
			jwtSecret:     "test-secret",
			expireTimeSec: 3600,
			expectedError: true,
		},
		{
			name:          "Failure - Token signed with different secret",
			token:         "", // Will be generated with different secret
			jwtSecret:     "test-secret",
			expireTimeSec: 3600,
			expectedError: true,
		},
		{
			name:          "Failure - Expired token",
			token:         "", // Will be generated with short expiry
			jwtSecret:     "test-secret",
			expireTimeSec: 1,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewJwtService(tt.jwtSecret, tt.expireTimeSec)
			var token string
			var err error

			switch tt.name {
			case "Success - Validate valid token":
				token, err = service.GenerateToken("test-user-123")
				assert.NoError(t, err)
			case "Failure - Token signed with different secret":
				differentService := NewJwtService("different-secret", tt.expireTimeSec)
				token, err = differentService.GenerateToken("test-user-123")
				assert.NoError(t, err)
			case "Failure - Expired token":
				token, err = service.GenerateToken("test-user-123")
				assert.NoError(t, err)
				time.Sleep(2 * time.Second) // Wait for token to expire
			default:
				token = tt.token
			}

			jwtToken, err := service.ValidateToken(token)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, jwtToken)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, jwtToken)
				assert.Equal(t, "test-user-123", jwtToken.UserId)
				assert.True(t, jwtToken.ExpiredAt.After(time.Now()))
			}
		})
	}
}

func TestJwtService_ValidateToken_ExpiredToken(t *testing.T) {
	service := NewJwtService("test-secret", 1)
	token, err := service.GenerateToken("test-user-123")
	assert.NoError(t, err)

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	jwtToken, err := service.ValidateToken(token)
	assert.Error(t, err)
	assert.Nil(t, jwtToken)
}

func TestJwtService_ValidateToken_InvalidClaims(t *testing.T) {
	service := NewJwtService("test-secret", 3600)
	token, err := service.GenerateToken("test-user-123")
	assert.NoError(t, err)

	// Modify the token to make it invalid
	invalidToken := token + "invalid"

	jwtToken, err := service.ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, jwtToken)
}
