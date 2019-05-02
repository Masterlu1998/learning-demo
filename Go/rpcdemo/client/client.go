package client

import(
	"net/rpc"
)

func GetRpcClient() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1122")
	if err != nil {
		return nil, err
	}
	return client, nil
}