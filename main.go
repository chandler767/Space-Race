package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available cpu cores.
	config := pubnub.NewConfig()
	config.SubscribeKey = "sub-key"
	config.PublishKey = "pub-key"
	pn := pubnub.NewPubNub(config)
	fmt.Println("Welcome to Space Race!")
	newLobby("", "", pn) // Create lobby for a new game.
}
