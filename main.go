package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
)

func main() {
	config := pubnub.NewConfig()
	config.SubscribeKey = "sub-key"
	config.PublishKey = "pub-key"
	pn := pubnub.NewPubNub(config)
	fmt.Println("Welcome to Space Race!")
	newLobby("", "", pn) // Create lobby for a new game.
}
