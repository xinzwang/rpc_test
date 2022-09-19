package rpc

import (
	"context"
	"encoding/json"

	console_pbs "e.coding.net/itdesk/weixin/golib/pbs/console"
	"e.coding.net/itdesk/weixin/golib/utils"
)

func (a *Api) PosDetailRpc(data string) (interface{}, error) {
	req := console_pbs.PosDetailReq{}
	json.Unmarshal([]byte(data), &req)
	rsp, err := utils.ConsoleClient().PosDetailRpc(context.Background(), &req)
	return rsp, err
}
