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
		return app.HttpError("Can't process task at this time, try again", app.HttpStatus.InternalServerError)
	}
	users, error := us.UserRepository.FindAll()
	if error != nil {
		return app.HttpErrorWithLog("Error occurred, try again later", app.HttpStatus.ExpectationFailed, error)
	}
	return app.HttpSuccessWithData("User successfully fetched", app.HttpStatus.OK, users)
}
