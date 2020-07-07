package services

type IUserService interface {
	GetName (userid int) string
}

type UserService struct {

}

func (u UserService) GetName(userid int) string {
	if userid == 101 {
		return "admin"
	}

	return "guest"
}
