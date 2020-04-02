package alienattach

import (
	attach "gitlab.fbk168.com/gamedevjp/backend-utility/utility/attach"
)

// UserAttach ...
type UserAttach struct {
	userID  int64
	dataMap map[int64]map[int64]*attach.Info
}

// LoadData ...
func (us *UserAttach) LoadData() {
	// test Data
	us.dataMap = make(map[int64]map[int64]*attach.Info)
	//---
	// redis load data

	// if fail sql load data
}

// Get ...
func (us *UserAttach) Get(attachkind int64, attachtype int64) *attach.Info {
	if _, ok := (*us.GetType(attachkind))[attachtype]; !ok {
		us.SetValue(attachkind, attachtype, "", 0)
	}
	return us.dataMap[attachkind][attachtype]
}

// GetType ...
func (us *UserAttach) GetType(attachkind int64) *map[int64]*attach.Info {
	if _, ok := us.dataMap[attachkind]; !ok {
		us.dataMap[attachkind] = make(map[int64]*attach.Info)
	}
	result := us.dataMap[attachkind]
	return &result
}

// SetDBValue ...
func (us *UserAttach) SetDBValue(attachKind, attachType int64, SValue string, IValue int64) {

	if att, ok := (*us.GetType(attachKind))[attachType]; !ok {
		att = attach.NewInfo(attachKind, attachType, true)
		att.SetSValue(SValue)
		att.SetIValue(IValue)
		us.dataMap[attachKind][attachType] = att
	} else {
		att.SetSValue(SValue)
		att.SetIValue(IValue)
	}
}

// SetValue ...
func (us *UserAttach) SetValue(attachKind, attachType int64, SValue string, IValue int64) {

	if att, ok := (*us.GetType(attachKind))[attachType]; !ok {
		att = attach.NewInfo(attachKind, attachType, false)
		att.SetSValue(SValue)
		att.SetIValue(IValue)
		us.dataMap[attachKind][attachType] = att
	} else {
		att.SetSValue(SValue)
		att.SetIValue(IValue)
	}
}

// SetAttach ...
func (us *UserAttach) SetAttach(info *attach.Info) {
	us.dataMap[info.GetKind()][info.GetTypes()] = info
}

// Save ...
func (us *UserAttach) Save() {

}

// Clear ...
func (us *UserAttach) Clear() {

}
