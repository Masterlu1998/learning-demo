package main

import (
	"rpcdemo/service"
	"rpcdemo/client"
	"fmt"
	"time"
)

func main() {
	service.StartRPCService()
	time.Sleep(3 * time.Second)
	client, err := client.GetRpcClient()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	args := &service.Args{ A: 7, B: 8 }
	var result int
	err = client.Call("Arith.GetMessage", args, &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

}