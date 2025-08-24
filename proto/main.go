package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type VariableInfo struct {
	ItemID     uint64
	AffiShopID uint64
	OrderID    uint64
}
type InitSessionVariable struct {
	UserID           *uint64
	AccountType      *int32
	MainSubAccountID *uint64 // 主子账号独立系统，由前端传入提供
	UserLoginStatus  *int32  // 0 not logged in ;1 logged in
	SessionID        *string
	OrderID          *uint64
	ItemID           *uint64
	AffiShopID       *uint64
}

func main() {
	initVariable := &InitSessionVariable{
		UserID:     proto.Uint64(1),
		AffiShopID: proto.Uint64(23135),
	}
	variableInfoStr, err1 := json.Marshal(initVariable)
	println(string(variableInfoStr))
	fmt.Println(err1)

	variable := &VariableInfo{}
	err2 := json.Unmarshal(variableInfoStr, variable)
	println(variable.ItemID, variable.AffiShopID, variable.OrderID)
	fmt.Println(err2)

}
