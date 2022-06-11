package user_test

import (
	_domain "gin-bmg-restful/entities/domain"
	_web "gin-bmg-restful/entities/web"
	_userRepository "gin-bmg-restful/repositories/user"
	_userService "gin-bmg-restful/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)


func TestFindAll(t *testing.T) {
	t.Run("case-success", func(t *testing.T) {		
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("FindAll").Return(_userRepository.SampleData, nil)

		userService := _userService.NewService(userMock)
		users, err := userService.FindAll([]map[string]string{})
		assert.Nil(t, err)
		assert.NotEmpty(t, users)
	})

	t.Run("case-repo-fail", func(t *testing.T) {		
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("FindAll").Return([]_domain.User{}, _web.Error{})

		userService := _userService.NewService(userMock)
		users, err := userService.FindAll([]map[string]string{})
		assert.NotNil(t, err)
		assert.Empty(t, users)
	})
}

func TestFind(t *testing.T) {
	t.Run("case-success", func(t *testing.T) {		
		sampleUser := _userRepository.SampleData[0]
		sampleResponse := _web.UserResponse {
			Username: sampleUser.Username,
			Name: sampleUser.Name,
			Email: sampleUser.Email,
			ReferralCode: sampleUser.ReferralCode,
		}
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("Find").Return(_userRepository.SampleData[0], nil)

		userService := _userService.NewService(userMock)
		user, err := userService.Find("ahmad")
		assert.Nil(t, err)
		assert.Equal(t, sampleResponse, user)
	})

	t.Run("case-repo-fail", func(t *testing.T) {		
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("Find").Return(_domain.User{}, _web.Error{})

		userService := _userService.NewService(userMock)
		userRes, err := userService.Find("ahmad")
		assert.NotNil(t, err)
		assert.Equal(t, _web.UserResponse{}, userRes)
	})
}

func TestUpdateProfile(t *testing.T) {
	t.Run("case-success", func(t *testing.T) {		
		sampleUser := _userRepository.SampleData[0]
		sampleInput := _web.UserRequest{
			Email: "ahmadmailbaru@mail.com",
			Password: "ahmadpassbaru",
		}
		sampleResponse := _web.UserResponse {
			Username: sampleUser.Username,
			Name: sampleUser.Name,
			Email: sampleInput.Email,
			ReferralCode: sampleUser.ReferralCode,
		}
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		newPass, _ := bcrypt.GenerateFromPassword([]byte(sampleInput.Password), bcrypt.DefaultCost)
		sampleUser.Email = sampleInput.Email
		sampleUser.Password = string(newPass)
		userMock.Mock.On("Update").Return(sampleUser, nil)

		userService := _userService.NewService(userMock)
		user, err := userService.UpdateProfile(sampleInput, sampleUser.Username)
		assert.Nil(t, err)
		assert.Equal(t, sampleResponse, user)
	})
}

func TestCheckReferral(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sampleUser := _userRepository.SampleData[0]
		sampleResponse := _web.UserResponse {
			Username: sampleUser.Username,
			Name: sampleUser.Name,
			Email: sampleUser.Email,
			ReferralCode: sampleUser.ReferralCode,
		}
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("FindBy").Return(sampleUser, nil)
		userService := _userService.NewService(userMock)
		// Ahmad's referral code
		ownerRes, err := userService.CheckReferral(_web.ReferralRequest{
			Referral: "ref-3310",
		}, "budi")

		assert.Nil(t, err)
		assert.Equal(t, sampleResponse, ownerRes)
	})
	t.Run("own-referral", func(t *testing.T) {
		sampleUser := _userRepository.SampleData[0]
		userMock := _userRepository.NewRepositoryMock(&mock.Mock{})
		userMock.Mock.On("FindBy").Return(sampleUser, nil)
		userService := _userService.NewService(userMock)
		// Ahmad's referral code
		ownerRes, err := userService.CheckReferral(_web.ReferralRequest{
			Referral: "ref-3310",
		}, "ahmad")

		assert.Error(t, err)
		assert.Equal(t, _web.UserResponse{}, ownerRes)
	})
}