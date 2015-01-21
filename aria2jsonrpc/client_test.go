package aria2jsonrpc

import (
	"fmt"
	"testing"
	"time"
)

var client *Client

func init() {
	server := "http://localhost:6800/jsonrpc"
	client = New(server)

}

func TestA(t *testing.T) {

	fmt.Println(client.GetVersion())
	downloadUri := "http://dldir1.qq.com/qqfile/QQforMac/QQ_V4.0.1.dmg"

	g, err := client.AddUri(downloadUri)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(g)
	fmt.Println(client.TellActive())
	// fmt.Println(client.ForcePause(g))
	// fmt.Println(client.Unpause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(client.TellStatus(g))
	// time.Sleep(1 * time.Second)
	// fmt.Println(client.Pause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(client.PauseAll())
	fmt.Println(client.Remove(g))
	fmt.Println(client.TellActive())
}
