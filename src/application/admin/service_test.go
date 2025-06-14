package admin

import (
	"testing"

	"time"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	mockjwt "github.com/Mrityunjoy99/sample-go/src/mocks/domain/service"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	mockJwt := new(mockjwt.JwtService)
	expireTimeSec := 3600

	t.Run("new_service", func(t *testing.T) {
		svc := NewService(mockJwt, expireTimeSec)
		assert.NotNil(t, svc)
		assert.Implements(t, (*Service)(nil), svc)
	})

	svc := NewService(mockJwt, expireTimeSec)

	t.Run("success", func(t *testing.T) {
		userId := "user-123"
		userType := "USER"
		mockJwt.On("GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeUser
		})).Return("token-abc", nil)

		token, err := svc.GenerateToken(userId, userType)
		assert.NoError(t, err)
		assert.Equal(t, "token-abc", token)
		mockJwt.AssertCalled(t, "GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeUser
		}))
	})

	t.Run("failure", func(t *testing.T) {
		userId := "user-err"
		userType := "USER"
		errExpected := genericerror.NewInternalErrByErr(assert.AnError)
		mockJwt.On("GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeUser
		})).Return("", errExpected)

		token, err := svc.GenerateToken(userId, userType)
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, errExpected, err)
		mockJwt.AssertCalled(t, "GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeUser
		}))
	})

	t.Run("invalid_user_type", func(t *testing.T) {
		userId := "user-123"
		userType := "INVALID"
		
		token, err := svc.GenerateToken(userId, userType)
		assert.Error(t, err)
		assert.Equal(t, constant.ErrorCodeBadRequest, err.GetCode())
		assert.Equal(t, "", token)
		mockJwt.AssertNotCalled(t, "GenerateToken")
	})

	t.Run("manager_user_type", func(t *testing.T) {
		userId := "manager-123"
		userType := "MANAGER"
		mockJwt.On("GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeManager
		})).Return("token-manager", nil)

		token, err := svc.GenerateToken(userId, userType)
		assert.NoError(t, err)
		assert.Equal(t, "token-manager", token)
		mockJwt.AssertCalled(t, "GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeManager
		}))
	})

	t.Run("admin_user_type", func(t *testing.T) {
		userId := "admin-123"
		userType := "ADMIN"
		mockJwt.On("GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeAdmin
		})).Return("token-admin", nil)

		token, err := svc.GenerateToken(userId, userType)
		assert.NoError(t, err)
		assert.Equal(t, "token-admin", token)
		mockJwt.AssertCalled(t, "GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeAdmin
		}))
	})

	t.Run("token_expiration", func(t *testing.T) {
		userId := "user-exp"
		userType := "USER"
		mockJwt.On("GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.UserId == userId && token.UserType == constant.UserTypeUser
		})).Return("token-exp", nil)

		token, err := svc.GenerateToken(userId, userType)
		assert.NoError(t, err)
		assert.Equal(t, "token-exp", token)

		// Verify the token expiration time is set correctly
		mockJwt.AssertCalled(t, "GenerateToken", mock.MatchedBy(func(token *entity.JwtToken) bool {
			return token.ExpiredAt.Sub(time.Now()).Seconds() >= float64(expireTimeSec-1)
		}))
	})
}
