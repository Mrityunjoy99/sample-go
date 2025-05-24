package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// DummyClaims implements jwt.Claims but is not jwt.MapClaims
type DummyClaims struct{}

func (DummyClaims) Valid() error { return nil }

func (DummyClaims) GetExpirationTime() (*jwt.NumericDate, error) { return nil, nil }

func (DummyClaims) GetIssuedAt() (*jwt.NumericDate, error) { return nil, nil }

func (DummyClaims) GetNotBefore() (*jwt.NumericDate, error) { return nil, nil }

func (DummyClaims) GetIssuer() (string, error) { return "", nil }

func (DummyClaims) GetSubject() (string, error) { return "", nil }

func (DummyClaims) GetAudience() (jwt.ClaimStrings, error) { return nil, nil }

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

func TestJwtService_ValidateToken_MissingOrWrongTypeClaims(t *testing.T) {
	secret := "test-secret"
	expire := 3600
	service := NewJwtService(secret, expire)

	t.Run("missing userId claim", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			// "userId" is missing
			"exp": time.Now().Add(time.Duration(expire) * time.Second).Unix(),
		})
		tokenStr, err := token.SignedString([]byte(secret))
		assert.NoError(t, err)
		jwtToken, err := service.ValidateToken(tokenStr)
		assert.Error(t, err)
		assert.Nil(t, jwtToken)
		assert.Contains(t, err.Error(), "userId")
	})

	t.Run("missing exp claim", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": "test-user-123",
			// "exp" is missing
		})
		tokenStr, err := token.SignedString([]byte(secret))
		assert.NoError(t, err)
		jwtToken, err := service.ValidateToken(tokenStr)
		assert.Error(t, err)
		assert.Nil(t, jwtToken)
		assert.Contains(t, err.Error(), "exp")
	})

	t.Run("userId is not a string", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": 12345, // int instead of string
			"exp":    time.Now().Add(time.Duration(expire) * time.Second).Unix(),
		})
		tokenStr, err := token.SignedString([]byte(secret))
		assert.NoError(t, err)
		jwtToken, err := service.ValidateToken(tokenStr)
		assert.Error(t, err)
		assert.Nil(t, jwtToken)
		assert.Contains(t, err.Error(), "userId")
	})

	t.Run("exp is not a float", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": "test-user-123",
			"exp":    "not-a-float", // string instead of float
		})
		tokenStr, err := token.SignedString([]byte(secret))
		assert.NoError(t, err)
		jwtToken, err := service.ValidateToken(tokenStr)
		assert.Error(t, err)
		assert.Nil(t, jwtToken)
		assert.Contains(t, err.Error(), "exp")
	})
}

func TestJwtService_validateTokenFromParsedToken_InvalidClaimsType(t *testing.T) {
	service := NewJwtService("test-secret", 3600)
	// Create a jwt.Token with a claims type that is not jwt.MapClaims
	token := &jwt.Token{
		Claims: DummyClaims{}, // Not a MapClaims!
	}
	jwtToken, err := service.(*jwtService).validateTokenFromParsedToken(token)
	assert.Error(t, err)
	assert.Nil(t, jwtToken)
	assert.Contains(t, err.Error(), "invalid token claims")
}
