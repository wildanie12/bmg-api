package user

import (
	"fmt"
	_domain "gin-bmg-restful/entities/domain"
	_web "gin-bmg-restful/entities/web"

	"gorm.io/gorm"
)

// Repository represent main struct object
type Repository struct {
	db *gorm.DB
}

// NewRepository is an constructor to
// construct Repository 
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindAll return all user data based
// on what filter provided in the parameter 
// 	 filter format: { field, operator, value }
//   filter can be more than one since it is an array
func (repo Repository) FindAll(filters []map[string]string) ([]_domain.User, error) {
	users := []_domain.User{}

	// filter
	db := repo.db.Model(&_domain.User{})
	for _, filter := range filters {
		db = db.Where(fmt.Sprintf("%s %s ?", filter["field"], filter["operator"]), filter["value"])
	}

	// execute query
	tx := db.Find(&users)
	if tx.Error != nil {
		return []_domain.User{}, _web.Error{ Message: tx.Error.Error() }
	}
	
	return users, nil
}

// Find method can get the single user resource
// based on provided id parameter
func (repo Repository) Find(username string) (_domain.User, error) {
	user := _domain.User{}
	tx := repo.db.Where("username = ?", username).First(&user)
	if tx.RowsAffected <= 0 {
		return _domain.User{}, _web.Error{ Message: "User not found", Code: 400 }
	} else if tx.Error != nil {
		return _domain.User{}, _web.Error{ Message: tx.Error.Error(), Code: 500 }
	}
	return user, nil
}

// FindBy method finds single user using specified custom
// field and value provided in parameter
func (repo Repository) FindBy(field, value string) (_domain.User, error) {
	user := _domain.User{}
	tx := repo.db.Where(field + " = ?", value).First(&user)
	if tx.RowsAffected <= 0 {
		return _domain.User{}, _web.Error{ Message: "User not found", Code: 400 }
	} else if tx.Error != nil {
		return _domain.User{}, _web.Error{ Message: tx.Error.Error(), Code: 500 }
	}
	return user, nil
}

// Create user with newly provided user
// entity domain and persistence it to data source
func (repo Repository) Create(user _domain.User) (_domain.User, error) {
	tx := repo.db.Create(user)
	if tx.Error != nil {
		return _domain.User{}, _web.Error{ Message: tx.Error.Error(), Code: 500 }
	}
	return user, nil
}

// Update user with newly provided user
// entity domain and persistence it to data source
func (repo Repository) Update(user _domain.User, username string) (_domain.User, error) {
	tx := repo.db.Model(&_domain.User{}).Where("username = ?", username).Updates(user)
	if tx.Error != nil {
		return _domain.User{}, _web.Error{ Message: tx.Error.Error(), Code: 500 }
	}
	return user, nil
}

// Delete user by id provided in parameter
func (repo Repository) Delete(username string) error {
	tx := repo.db.Where("username = ?", username).Delete(_domain.User{})
	if tx.Error != nil {
		return _web.Error{ Message: tx.Error.Error(), Code: 500 }
	}
	return nil
}