package auth_test

import (
	_web "gin-bmg-restful/entities/web"
	_userRepository "gin-bmg-restful/repositories/user"
	_authService "gin-bmg-restful/services/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		userSample := _userRepository.SampleData[0]		
		userResponseSample := _web.UserResponse{
			Username: userSample.Username,
			Name: userSample.Name,
			Email: userSample.Email,
			ReferralCode: userSample.ReferralCode,
		}
		userRepo := _userRepository.NewRepositoryMock(&mock.Mock{})
		userRepo.Mock.On("Find").Return(userSample, nil)

		authService := _authService.NewService(userRepo)
		response, err := authService.Login(_web.AuthRequest{
			Username: "ahmad",
			Password: "ahmad123",
		})

		assert.NotEmpty(t, response.Token)
		assert.Equal(t, userResponseSample, response.User)
		assert.Nil(t, err)
	})
}

func TestRegister(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		userSample := _userRepository.SampleData[0]
		userResponseSample := _web.UserResponse{
			Username: userSample.Username,
			Name: userSample.Name,
			Email: userSample.Email,
			ReferralCode: userSample.ReferralCode,
		}

		userRepo := _userRepository.NewRepositoryMock(&mock.Mock{})
		userRepo.Mock.On("Create").Return(userSample, nil)

		userService := _authService.NewService(userRepo)
		response, err := userService.Register(_web.UserRequest{
			Username: userSample.Username,
			Email: userSample.Email,
			Password: "ahmad123",
			Name: userSample.Name,
		})	

		assert.Nil(t, err)
		assert.NotEmpty(t, response.Token)
		assert.Equal(t, userResponseSample, response.User)
	})
	t.Run("validation-error", func(t *testing.T) {
		userSample := _userRepository.SampleData[0]

		userRepo := _userRepository.NewRepositoryMock(&mock.Mock{})
		userRepo.Mock.On("Create").Return(userSample, nil)

		userService := _authService.NewService(userRepo)
		response, err := userService.Register(_web.UserRequest{
			Username: userSample.Username,
			Email: userSample.Email,
			Password: "",
			Name: userSample.Name,
		})	

		assert.NotNil(t, err)
		assert.Empty(t, response.Token)
		assert.Equal(t, _web.UserResponse{}, response.User)
	})
}
