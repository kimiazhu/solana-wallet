package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

func main() {
	// create a RPC client
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	// get the current running Solana version
	response, err := c.GetVersion(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("version", response.SolanaCore)

	balance, err := c.GetBalance(context.TODO(), "9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an")
	if err != nil {
		panic(err)
	}

	fmt.Println(balance)
}
