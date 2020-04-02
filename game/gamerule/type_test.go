package gamerule

import (
	"fmt"
	"testing"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
)

func TestRule_GameRequest(t *testing.T) {
	mmi := make(map[string]map[string]interface{})
	mmi["1"] = make(map[string]interface{}, 0)
	// att := gameattach.NewAttach(5)
	// att.SetValue(7, 5, "", 1)
	// att.SetValue(7, 6, "", 2)
	// att.SetValue(7, 7, "", 3)

	// uaa := att.(*gameattach.UserAttach)
	mmi["1"]["2"] = 3
	fmt.Println(foundation.JSONToString(mmi))
}
