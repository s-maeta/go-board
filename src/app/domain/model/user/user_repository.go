package user

type IUserRepository interface {
	Create(user *User) error
	Delete(user *User) error
	FindForLoginUser(email Email, password Password) *User
	FindForUniqueId(uniqueId string) *User
}
