package user

type UserGetRequest struct {
	UserId int `from:user json:"user" binding:"required"`
}
