package repository

import (
	"CleanCode/entidades_negocio"
	"CleanCode/user/contracts"
	"CleanCode/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

//MockDB ...
type MockDB struct {
	mock.Mock
}

//Data ...
func (mockDB *MockDB) Data() map[int]entidades_negocio.EUser {
	return StructureDB()
}

//StructureDB ...
func StructureDB() map[int]entidades_negocio.EUser {
	return map[int]entidades_negocio.EUser{
		0: entidades_negocio.EUser{
			ID:    1,
			Name:  "Test",
			Email: "Test@",
			Age:   18,
		},
	}
}

//MockRepository ...
type MockRepository struct {
	mock.Mock
}

//GetUserByIDMock ...
func (mockRepo *MockRepository) GetUserByIDMock(id int) (entidades_negocio.EUser, error) {
	args := mockRepo.Called(id)
	return args.Get(0).(entidades_negocio.EUser), args.Error(1)
}

//Initialize testing

type RUserTest struct {
	RUserMock *MockRepository
	IUser     contracts.IRepository
}

func setup() *RUserTest {
	mockDB := &MockDB{}
	mockDB.On("Data").Return(map[int]entidades_negocio.EUser{})
	return &RUserTest{
		RUserMock: &MockRepository{},
		IUser:     repository.NewRUserRepo(mockDB),
	}
}

func TestGetUserByID(t *testing.T) {
	var userID int = 0
	test := setup()
	test.RUserMock.On("GetUserByIDMock", userID).Return(entidades_negocio.EUser{}, nil)
	response, err := test.IUser.GetUserByID(userID)

	assert.Empty(t, err)
	assert.NotEmpty(t, response)
}
