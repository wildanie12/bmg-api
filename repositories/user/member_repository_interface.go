package user

import "gin-bmg-restful/entities/domain"

// RepositoryInterface provide several contracts
// Regarding abstraction of a user data source in repository layer
type RepositoryInterface interface {
	
	// FindAll return all user data based
	// on what filter provided in the parameter 
	// filter format:
	//   { field, operator, value }
	//   filter can be more than one since it is an array
	FindAll(filter []map[string]string) ([]domain.User, error)
	
	// Find method can get the single user resource
	// based on provided id parameter
	Find(username string) (domain.User, error)

	// FindBy method finds single user using specified custom
	// field and value provided in parameter
	FindBy(field, value string) (domain.User, error)

	// Update user with newly provided user
	// entity domain and persistence it to data source
	Create(user domain.User) (domain.User, error)

	// Update user with newly provided user
	// entity domain and persistence it to data source
	Update(user domain.User, username string) (domain.User, error)

	// Delete user by id provided in parameter
	Delete(username string) error
}