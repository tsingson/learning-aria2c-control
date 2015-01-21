package aria2jsonrpc

import (
	"fmt"
	"testing"
)

func Test_Call(t *testing.T) {
	var reply interface{}

	server := "http://localhost:6800/jsonrpc"
	downloadUri := "http://dldir1.qq.com/qqfile/QQforMac/QQ_V4.0.1.dmg"

	err := Call(server, addUri, []interface{}{[]string{downloadUri}}, &reply)
	fmt.Printf("%s %v\n", reply, err)
}
