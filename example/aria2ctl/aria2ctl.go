package main

import (
	"aria2jsonrpc"
	"fmt"
	"time"

)

const (
	server = "http://localhost:6800/jsonrpc"
)

func main() {
	var client *aria2jsonrpc.Client
	client = aria2jsonrpc.New(server)
	fmt.Println(client.GetVersion())

	downloadUri := "http://dldir1.qq.com/qqfile/QQforMac/QQ_V4.0.1.dmg"

	var parameter = make(map[string]string, 1)

	/*
		parameter["gid"] = "061f67a7513bdb3e"

	*/

	g, err := client.AddUri(downloadUri, parameter)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	//g := "061f67a7513bdb3e"
	fmt.Println(g)
	fmt.Println(client.TellActive())
	fmt.Println(client.ForcePause(g))
	// fmt.Println(client.Unpause(g))
	time.Sleep(1 * time.Second)
	//fmt.Println(client.TellStatus(g))
	time.Sleep(1 * time.Second)
	//fmt.Println(client.Pause(g))
	//time.Sleep(1 * time.Second)
	//fmt.Println(client.PauseAll())
	//fmt.Println(client.Remove(g))
	//fmt.Println(client.TellActive())
}
