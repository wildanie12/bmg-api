package auth_test

import (
	"gin-bmg-restful/entities/web"
	_userRepository "gin-bmg-restful/repositories/user"
	_authService "gin-bmg-restful/services/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		userSample := _userRepository.SampleData[0]		
		userResponseSample := web.UserResponse{
			Username: userSample.Username,
			Name: userSample.Name,
			Email: userSample.Email,
			ReferralCode: userSample.ReferralCode,
		}
		userRepo := _userRepository.NewRepositoryMock(&mock.Mock{})
		userRepo.Mock.On("Find").Return(userSample, nil)

		authService := _authService.NewService(userRepo)
		response, err := authService.Login(web.AuthRequest{
			Username: "ahmad",
			Password: "ahmad123",
		})

		assert.NotEmpty(t, response.Token)
		assert.Equal(t, userResponseSample, response.User)
		assert.Nil(t, err)
	})
}