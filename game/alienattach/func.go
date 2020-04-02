package alienattach

import (
	attach "gitlab.fbk168.com/gamedevjp/backend-utility/utility/attach"
)

// NewUserAttach ...
func NewUserAttach(userID int64) *UserAttach {
	attach := &UserAttach{
		userID:  userID,
		dataMap: make(map[int64]map[int64]*attach.Info),
	}
	// attach.InitData(userID)
	return attach
}
