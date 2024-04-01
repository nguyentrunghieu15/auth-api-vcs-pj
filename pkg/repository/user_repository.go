package repository

import (
	"fmt"
	"log"

	"github.com/nguyentrunghieu15/common-vcs-prj/db/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	IsExsitsUserByEmail(string) (bool, *model.User, error)
	FindUserById(int) (*model.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func (u *UserRepository) IsExsitsUserByEmail(email string) (bool, *model.User, error) {
	if u.Db == nil {
		log.Printf("User Repository: Database connetion is nil\n")
		return false, nil, fmt.Errorf("User Repository: Database connetion is nil")
	}
	var user = &model.User{}
	result := u.Db.Where("email = ?", email).First(user)
	if result.Error != nil {
		log.Printf("User Repository: %v\n", result.Error)
		return false, nil, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil, nil
	}

	return true, user, nil
}

func (u *UserRepository) FindUserById(id int) (*model.User, error) {
	if u.Db == nil {
		log.Printf("User Repository: Database connetion is nil\n")
		return nil, fmt.Errorf("User Repository: Database connetion is nil")
	}
	var user = &model.User{}
	result := u.Db.First(user, id)
	if result.Error != nil {
		log.Printf("User Repository: %v\n", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}
