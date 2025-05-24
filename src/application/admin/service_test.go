package admin

import (
	"errors"
	"testing"
	"time"

	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	mockjwt "github.com/Mrityunjoy99/sample-go/src/mocks/domain/service"
	"github.com/stretchr/testify/assert"
)

func TestService_GenerateToken(t *testing.T) {
	mockJwt := new(mockjwt.JwtService)
	svc := NewService(mockJwt)

	t.Run("success", func(t *testing.T) {
		mockJwt.On("GenerateToken", "user-123").Return("token-abc", nil)
		token, err := svc.GenerateToken("user-123")
		assert.NoError(t, err)
		assert.Equal(t, "token-abc", token)
		mockJwt.AssertCalled(t, "GenerateToken", "user-123")
	})

	t.Run("failure", func(t *testing.T) {
		errExpected := errors.New("signing error")
		mockJwt.On("GenerateToken", "user-err").Return("", errExpected)
		token, err := svc.GenerateToken("user-err")
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, errExpected, err)
		mockJwt.AssertCalled(t, "GenerateToken", "user-err")
	})
}

func TestService_ValidateToken(t *testing.T) {
	mockJwt := new(mockjwt.JwtService)
	svc := NewService(mockJwt)

	t.Run("success", func(t *testing.T) {
		exp := time.Now().Add(time.Hour)
		mockJwt.On("ValidateToken", "token-abc").Return(&entity.JwtToken{
			UserId:    "user-123",
			ExpiredAt: exp,
		}, nil)
		resp, err := svc.ValidateToken("token-abc")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "user-123", resp.UserId)
		assert.Equal(t, exp.Format(time.RFC850), resp.ExpiredAt)
		mockJwt.AssertCalled(t, "ValidateToken", "token-abc")
	})

	t.Run("failure", func(t *testing.T) {
		errExpected := errors.New("invalid token")
		mockJwt.On("ValidateToken", "bad-token").Return((*entity.JwtToken)(nil), errExpected)
		resp, err := svc.ValidateToken("bad-token")
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errExpected, err)
		mockJwt.AssertCalled(t, "ValidateToken", "bad-token")
	})
}
