package services

import "errors"

type IUserService interface {
	GetName(userid int) string
	DeleteUser(userid int) error
}

type UserService struct {
}

func (u UserService) DeleteUser(userid int) error {
	if userid == 101 {
		return errors.New("无权限")
	}

	return nil
}

func (u UserService) GetName(userid int) string {
	if userid == 101 {
		return "Tom"
	}

	return "guest"
}
