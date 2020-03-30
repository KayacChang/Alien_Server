package alienattach

import "github.com/YWJSonic/ServerUtility/attach"

// NewUserAttach ...
func NewUserAttach(userID int64) *UserAttach {
	attach := &UserAttach{
		userID:  userID,
		dataMap: make(map[int64]map[int64]*attach.Info),
	}
	// attach.InitData(userID)
	return attach
}
