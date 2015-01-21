package aria2jsonrpc

import (
	"github.com/kdar/httprpc"
)

func Call(address, method string, params, reply interface{}) error {

	jsonRpcVersion := RpcVersion

	err := httprpc.CallJson(jsonRpcVersion, address, method, params, &reply)
	return err

}
