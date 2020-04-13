package alienattach

import (
	"database/sql"
	"fmt"

	attach "github.com/YWJSonic/ServerUtility/attach"
	"github.com/YWJSonic/ServerUtility/code"
	"gitlab.fbk168.com/gamedevjp/alien/server/game/db"
)

// Setting ...
type Setting struct {
	UserIDStr string
	Kind      int64
	DB        *sql.DB
}

// UserAttach ...
type UserAttach struct {
	userID    int64
	userIDStr string
	kind      int64
	db        *sql.DB
	dataMap   map[int64]map[int64]*attach.Info
}

// LoadData ...
func (us *UserAttach) LoadData() {
	// test Data
	us.dataMap = make(map[int64]map[int64]*attach.Info)
	//---
	// redis load data

	// if fail sql load data
	result, err := db.GetAttachKind(us.db, us.userIDStr, us.kind)
	if err.ErrorCode != code.OK {
		fmt.Println(err)
	}
	for _, row := range result {
		att := &attach.Info{
			Kind:     us.kind,
			Types:    row["Type"].(int64),
			IValue:   row["IValue"].(int64),
			IsDBData: true,
		}
		us.SetAttach(att)
		// us.dataMap[us.kind][att.Types] = &att
	}
	fmt.Println("done")
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
	info.IsDirty = true
	if _, ok := us.dataMap[info.GetKind()]; !ok {
		us.dataMap[info.GetKind()] = make(map[int64]*attach.Info)
	}
	us.dataMap[info.GetKind()][info.GetTypes()] = info
}

// Save ...
func (us *UserAttach) Save() {
	quarys := []string{}
	for kind, typeAtt := range us.dataMap {
		for t, att := range typeAtt {
			if att.IsDirty {
				quarys = append(quarys, fmt.Sprintf(us.setQuary(), us.userID, us.userIDStr, kind, t, att.IValue, att.IValue))
			}
		}
	}
	for _, quart := range quarys {
		Rows, err := us.db.Query(quart)

		if err != nil {
			fmt.Println(Rows, err)
		}
	}
}

// Clear ...
func (us *UserAttach) Clear() {

}

func (us *UserAttach) setQuary() string {
	return "INSERT INTO attach VALUES (%d, \"%s\", %d, %d, %d) ON DUPLICATE KEY UPDATE IValue = %d;\n"
}
