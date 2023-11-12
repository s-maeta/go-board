package pre_user

type IPreUserRepository interface {
	Create(preUser PreUser) error
	FindForToken(token Token) (PreUser, error)
}
