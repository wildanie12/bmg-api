package user_test

import (
	_config "gin-bmg-restful/config"
	_domain "gin-bmg-restful/entities/domain"
	_memberRepository "gin-bmg-restful/repositories/user"
	_utils "gin-bmg-restful/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)



func destroy(db *gorm.DB) {
	tx := db.Where("1=1").Delete(&_domain.User{})
	if tx.Error != nil {
		panic(tx.Error)
	}
}

var config = _config.New()
var dbTest = _utils.NewSQLiteConnection(config)
var repository = _memberRepository.NewRepository(dbTest)
func init() {
	_utils.Migrate(dbTest)
}

func fillData() {
	repository.Create(_memberRepository.SampleData[0])
	repository.Create(_memberRepository.SampleData[1])
}

func TestFindAll(t *testing.T) {
	
	
	t.Run("success", func(t *testing.T) {
		
		destroy(dbTest)
		repository.Create(_memberRepository.SampleData[0])
		data, err := repository.FindAll([]map[string]string{})

		assert.Equal(t, []_domain.User{ _memberRepository.SampleData[0] }, data)
		assert.Nil(t, err)
	})
	t.Run("success-with-filter", func(t *testing.T) {
		destroy(dbTest)
		fillData()
		
		data, err := repository.FindAll([]map[string]string{
			{ "field": "name", "operator": "LIKE", "value": "%ahm%" },
		})
		dataAll, err := repository.FindAll([]map[string]string{})
		assert.Equal(t, 2, len(dataAll))
		assert.Equal(t, []_domain.User{ _memberRepository.SampleData[0] }, data)
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	destroy(dbTest)
	fillData()
	t.Run("success", func(t *testing.T) {
		user, err := repository.Find("ahmad")
		assert.Nil(t, err)
		assert.Equal(t, _memberRepository.SampleData[0], user)
	})
}

func TestFindBy(t *testing.T) {
	destroy(dbTest)
	fillData()
	t.Run("success", func(t *testing.T) {
		user, err := repository.FindBy("email", "ahmad@mail.com")
		assert.Nil(t, err)
		assert.Equal(t, _memberRepository.SampleData[0], user)
	})
}

func TestUpdate(t *testing.T) {
	destroy(dbTest)
	fillData()
	t.Run("success", func(t *testing.T) {
		userUpdate := _domain.User{
			Email: "ahmadlagi@mail.com",
		}
		_, err := repository.Update(userUpdate, "ahmad")
		assert.Nil(t, err)
	})
}


func TestDelete(t *testing.T) {
	destroy(dbTest)
	fillData()
	t.Run("success", func(t *testing.T) {
		err := repository.Delete("ahmad")

		assert.Nil(t, err)
	})
}