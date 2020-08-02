package users

import (
	"fmt"
	"github.com/brragul/bookstore_users-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result := userDb[user.Id]
	if result == nil{
		return errors.NewNotFound(fmt.Sprintf("userID: %d not found", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreate = result.DateCreate

	return nil
}

func (user *User) Save() *errors.RestErr{
	current := userDb[user.Id]
	if current != nil{
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists",user.Id))
	}
	userDb[user.Id] = user
	return nil
}