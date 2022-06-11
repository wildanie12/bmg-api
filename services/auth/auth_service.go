package auth

import (
	_validations "gin-bmg-restful/deliveries/validations"
	_domain "gin-bmg-restful/entities/domain"
	_web "gin-bmg-restful/entities/web"
	_userRepository "gin-bmg-restful/repositories/user"
	"gin-bmg-restful/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// Service struct
type Service struct {
	userRepo _userRepository.RepositoryInterface
	validate *validator.Validate
}

// NewService generate new instance of Service
func NewService(userRepo _userRepository.RepositoryInterface) *Service {
	return &Service{
		userRepo: userRepo,
		validate: validator.New(),
	} 
}


// Login a registered user and returns JWT token
func (service Service) Login(authRequest _web.AuthRequest) (_web.AuthResponse, error) {
	// validate
	err := _validations.ValidateAuthRequest(service.validate, authRequest)
	if err != nil {
		return _web.AuthResponse{}, err
	}

	// get user data
	user, err := service.userRepo.Find(authRequest.Username)
	if err != nil {
		return _web.AuthResponse{}, _web.Error{
			Code: 400,
			Message: "your credentials is invalid",
		}
	}

	// compare password hash
	matchedPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))
	if matchedPassword != nil {
		return _web.AuthResponse{}, _web.Error{
			Code: 400,
			Message: "your credentials is invalid pass",
		}
	}

	// generate token
	token, _ := utils.CreateToken(user)
	userResponse := _web.UserResponse{}
	copier.Copy(&userResponse, &user)

	return _web.AuthResponse{
		Token: token,
		User: userResponse,
	}, nil
}

// Register a user
func (service Service) Register(request _web.UserRequest) (_web.AuthResponse, error) {
	// validate
	err := _validations.ValidateUserRequest(service.validate, request)
	if err != nil {
		return _web.AuthResponse{}, err
	}

	// convert to domain entity
	user := _domain.User{}
	copier.Copy(&user, &request)
	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	// repository action
	user, err = service.userRepo.Create(user)
	if err != nil {
		return _web.AuthResponse{}, err
	}

	// generate token
	token, _ := utils.CreateToken(user)
	userResponse := _web.UserResponse{}
	copier.Copy(&userResponse, &user)

	return _web.AuthResponse{
		Token: token,
		User: userResponse,
	}, nil
}
