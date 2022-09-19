package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"rpc_test/rpc"
	"time"

	"e.coding.net/itdesk/weixin/golib/config"
	"e.coding.net/itdesk/weixin/golib/utils"
)

var (
	funcMap map[string]reflect.Value
)

func parse() (string, string, string) {
	var envs string
	var funcName string
	var reqData string
	flag.StringVar(&envs, "envs", "dev", "choose envs")
	flag.StringVar(&funcName, "func", "", "choose rpc name interface")
	flag.StringVar(&reqData, "req_data", "{}", "request data.")
	flag.Parse()
	return envs, funcName, reqData
}

func main() {
	envs, funcName, reqData := parse()

	if envs == "dev" {
		utils.SetRpcConfig(&config.RpcConfig{
			WeixinAddr:  "localhost:13002",
			ConsoleAddr: "localhost:13012",
		})
	} else if envs == "test" {
		utils.SetRpcConfig(&config.RpcConfig{
			WeixinAddr:  "123.57.180.218:8602",
			ConsoleAddr: "123.57.180.218:8612",
		})
	}

	var valFunc rpc.Api
	val := reflect.ValueOf(&valFunc)

	f := val.MethodByName(funcName)
	if !f.IsValid() {
		fmt.Println("Invalid func name.")
		return
	}

	t1 := time.Now()
	res := f.Call([]reflect.Value{reflect.ValueOf(reqData)})
	t2 := time.Now()

	rsp, err := res[0].Interface(), res[1].Interface()

	if err != nil {
		fmt.Println("\nError ! ! !")
		fmt.Println(err.(error).Error())
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(rsp)

	fmt.Println()
	fmt.Println("Runtime:", t2.Sub(t1), "")
	fmt.Println("Response", buffer.String())

	return
}
