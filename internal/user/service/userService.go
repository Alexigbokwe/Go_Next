package service

import (
	"goNext/app"
	"goNext/internal/user/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository `inject:"type"`
}

func (us *UserService) GetUsers() app.HttpResponseType[any] {
	if us.UserRepository == nil {
		return app.HttpErrorResponse("Can't process task at this time, try again", app.HttpStatus.InternalServerError, nil, nil)
	}
	users, error := us.UserRepository.FindAll()
	if error != nil {
		return app.HttpErrorResponse("Error occurred, try again later", app.HttpStatus.ExpectationFailed, nil, error)
	}
	return app.HttpSuccessResponse("User successfully fetched", app.HttpStatus.OK, users)
}
