package service

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "tamirlan.kadyrbekov@gmail.com",
		password: "testing",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
