package rpc

import (
	"context"
	"encoding/json"

	console_pbs "e.coding.net/itdesk/weixin/golib/pbs/console"
	"e.coding.net/itdesk/weixin/golib/utils"
)

func (a *Api) CompanyDetailRpc(data string) (interface{}, error) {
	req := console_pbs.CompanyDetailReq{}
	json.Unmarshal([]byte(data), &req)

	rsp, err := utils.ConsoleClient().CompanyDetailRpc(context.Background(), &req)
	return rsp, err
}
