package user

import (
	"gin-bmg-restful/deliveries/validations"
	_domain "gin-bmg-restful/entities/domain"
	_web "gin-bmg-restful/entities/web"
	_userRepository "gin-bmg-restful/repositories/user"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// Service main struct
type Service struct {
	userRepo _userRepository.RepositoryInterface
	validate *validator.Validate
}

// NewService return new instance of service struct object
func NewService(userRepo _userRepository.RepositoryInterface) *Service {
	return &Service{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

// FindAll return list of user in correlation with
// provided filter to filter data output
func (service Service) FindAll(filter []map[string]string) ([]_web.UserResponse, error) {

	users, err := service.userRepo.FindAll(filter)
	if err != nil {
		return []_web.UserResponse{}, err
	}

	usersResponse := []_web.UserResponse{}
	copier.Copy(&usersResponse, &users)
	return usersResponse, nil
}

// Find single user based on username 
// passed on paramter value
func (service Service) Find(username string) (_web.UserResponse, error) {
	user, err := service.userRepo.Find(username)
	if err != nil {
		return _web.UserResponse{}, err
	}

	userResponse := _web.UserResponse{}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

// UpdateProfile users
func (service Service) UpdateProfile(request _web.UserRequest, username string) (_web.UserResponse, error) {

	user := _domain.User{}
	copier.Copy(&user, &request)

	if request.Password != "" {
		password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		user.Password = string(password)
	}

	user, err := service.userRepo.Update(_domain.User{}, request.Username)
	if err != nil {
		return _web.UserResponse{}, err
	}

	userResponse := _web.UserResponse{}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

// CheckReferral checks input referral by another user
// Referral cannot be used by the owner
func (service Service) CheckReferral(request _web.ReferralRequest, username string) (_web.UserResponse, error) {
	err := validations.ValidateReferralRequest(service.validate, request)
	if err != nil {
		return _web.UserResponse{}, err
	}

	owner, err := service.userRepo.FindBy("referral", request.Referral)
	if err != nil {
		return _web.UserResponse{}, _web.Error{
			Message: "Referral code is invalid",
			Code: 400,
		}
	}
	if owner.Username == username {
		return _web.UserResponse{}, _web.Error{
			Message: "Referral code is invalid, this is your own referral code",
			Code: 400,
		}
	}

	ownerResponse := _web.UserResponse{}
	copier.Copy(&ownerResponse, &owner)
	return ownerResponse, nil
}