package pre_user

type PreUser struct {
	UniqueId UniqueId
	Email    Email
	Password Password
	Token    Token
}

func NewPreUser(
	UniqueId UniqueId,
	Email Email,
	Password Password,
	Token Token,
) *PreUser {
	return &PreUser{
		UniqueId: UniqueId,
		Email:    Email,
		Password: Password,
		Token:    Token,
	}
}

func CreatePreUser(
	Email,
	Password string,
) (*PreUser, error) {

	uniqueId, err := NewUniqueId()
	if err != nil {
		return nil, err
	}

	email, err := NewEmail(Email)
	if err != nil {
		return nil, err
	}

	password, err := NewPassword(Password)
	if err != nil {
		return nil, err
	}

	token, err := NewToken(*uniqueId)
	if err != nil {
		return nil, err
	}

	newPreUser := PreUser{
		*uniqueId,
		*email,
		*password,
		*token,
	}

	return &newPreUser, nil
}
