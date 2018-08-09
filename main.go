package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available cpu cores.
	config := pubnub.NewConfig()
	config.SubscribeKey = "sub-c-8dd422dc-9a7e-11e8-b377-126307b646dc"
	config.PublishKey = "pub-c-083cb2bb-500d-4bfc-910b-228bed5b5d71"
	pn := pubnub.NewPubNub(config)
	fmt.Println("Welcome to Space Race!")
	newLobby("", "", pn) // Create lobby for a new game.
}
