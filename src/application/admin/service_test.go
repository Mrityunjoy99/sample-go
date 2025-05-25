package admin

import (
	"testing"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	mockjwt "github.com/Mrityunjoy99/sample-go/src/mocks/domain/service"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GenerateToken(t *testing.T) {
	mockJwt := new(mockjwt.JwtService)
	expireTimeSec := 3600
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
}
