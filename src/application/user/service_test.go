package user_test

import (
	"testing"

	"github.com/Mrityunjoy99/sample-go/src/application/user"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
	mock_repository "github.com/Mrityunjoy99/sample-go/src/mocks/infrastructure/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersTestSuite struct {
	suite.Suite
	service  user.Service
	userRepo *mock_repository.UserRepository
}

func (suite *UsersTestSuite) SetupTest() {
	suite.userRepo = mock_repository.NewUserRepository(suite.T())
	suite.service = user.NewService(suite.userRepo)
}

func (suite *UsersTestSuite) TestGetUserById() {
	// Arrange

	expectedUser := entity.User{
		BaseModel: database.BaseModel{
			Id: uuid.New(),
		},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jhon.doe@gmail.com",
		Phone:     "1234567890",
	}

	suite.userRepo.On("GetUserById", expectedUser.Id).Return(expectedUser, nil)
	actualUser, err := suite.service.GetUserById(expectedUser.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedUser.FirstName, actualUser.FirstName)
	assert.Equal(suite.T(), expectedUser.LastName, actualUser.LastName)
	assert.Equal(suite.T(), expectedUser.Email, actualUser.Email)
	assert.Equal(suite.T(), expectedUser.Phone, actualUser.Phone)
}

func (suite *UsersTestSuite) TestCreateUser() {
	// Arrange
	expectedUser := entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jhon.doe@gmail.com",
		Phone:     "1234567890",
	}

	createUserDto := user.CreateUserDto{
		FirstName: expectedUser.FirstName,
		LastName:  expectedUser.LastName,
		Email:     expectedUser.Email,
		Phone:     expectedUser.Phone,
	}

	suite.userRepo.On("CreateUser", expectedUser).Return(expectedUser, nil)
	actualUser, err := suite.service.CreateUser(createUserDto)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedUser.FirstName, actualUser.FirstName)
	assert.Equal(suite.T(), expectedUser.LastName, actualUser.LastName)
	assert.Equal(suite.T(), expectedUser.Email, actualUser.Email)
	assert.Equal(suite.T(), expectedUser.Phone, actualUser.Phone)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
