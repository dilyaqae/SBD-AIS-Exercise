package main

import (
	"exc8/client"
	"exc8/server"
	"time"
)

func main() {
	// start the server in the background
	go func() {
		server.StartGrpcServer()
	}()

	time.Sleep(1 * time.Second)
	// todo start client
	cl, err := client.NewGrpcClient()
	if err != nil {
		panic(err)
	}

	err = cl.Run()
	if err != nil {
		panic(err)
	}

	println("Orders complete!")
}
