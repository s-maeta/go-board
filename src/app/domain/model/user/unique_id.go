package user

import "github.com/rs/xid"

type UniqueId string

func NewUniqueId() (*UniqueId, error) {
	uid := xid.New()

	newUniqueId := UniqueId(uid.String())
	return &newUniqueId, nil
}
