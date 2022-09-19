package rpc

import (
	"context"
	"encoding/json"

	console_pbs "e.coding.net/itdesk/weixin/golib/pbs/console"
	"e.coding.net/itdesk/weixin/golib/utils"
)

func (a *Api) DepDetailRpc(data string) (interface{}, error) {
	req := console_pbs.DepDetailReq{}
	json.Unmarshal([]byte(data), &req)
	// fmt.Println(data)
	// fmt.Println(req)

	rsp, err := utils.ConsoleClient().DepDetailRpc(context.Background(), &req)
	// fmt.Println(err)
	// fmt.Println(rsp)
	return rsp, err
}
