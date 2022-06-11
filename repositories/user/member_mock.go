package user

import (
	"gin-bmg-restful/entities/domain"
	_domain "gin-bmg-restful/entities/domain"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var ahmadPassword, _ = bcrypt.GenerateFromPassword([]byte("ahmad123"), bcrypt.DefaultCost)
var budiPassword, _ = bcrypt.GenerateFromPassword([]byte("budi123"), bcrypt.DefaultCost)

// SampleData contains mock user data that can be used
// for testing purposes
var SampleData = []_domain.User{
	{
		Username: "ahmad",
		Password: string(ahmadPassword),
		Name: "Ahmad Cahyo",
		Email: "ahmad@mail.com",
		ReferralCode: "ref-3310",
	},
	{
		Username: "budi",
		Password: string(budiPassword),
		Name: "Budi Yanto",
		Email: "budi@mail.com",
		ReferralCode: "ref-3311",
	},
}

// RepositoryMock represent main struct object
type RepositoryMock struct {
	Mock *mock.Mock
}

// NewRepositoryMock is an constructor to
// construct RepositoryMock
func NewRepositoryMock(mock *mock.Mock) *RepositoryMock {
	return &RepositoryMock{
		Mock: mock,
	}
}

// FindAll return all user data based
// on what filter provided in the parameter 
// 	 filter format: { field, operator, value }
//   filter can be more than one since it is an array
func (repo RepositoryMock) FindAll(filters []map[string]string) ([]_domain.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

// Find method can get the single user resource
// based on provided id parameter
func (repo RepositoryMock) Find(id int) (_domain.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

// FindBy method finds single user using specified custom
// field and value provided in parameter
func (repo RepositoryMock) FindBy(field, value string) (_domain.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

// Create user with newly provided user
// entity domain and persistence it to data source
func (repo RepositoryMock) Create(user _domain.User) (_domain.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

// Update user with newly provided user
// entity _domain and persistence it to data source
func (repo RepositoryMock) Update(user _domain.User) (_domain.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

// Delete user by id provided in parameter
func (repo RepositoryMock) Delete(id int) error {
	args := repo.Mock.Called()
	return args.Error(0)
}