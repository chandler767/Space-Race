package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pubnub "github.com/pubnub/go"
)

func newLobby(channel string, username string, pn *pubnub.PubNub) {
	var (
		guestUsername string
		hostUsername  string
	)
	channel, username = userInput(channel, username)
	lobbylistener := pubnub.NewListener()
	data := make(map[string]interface{})
	go func() {
		for {
			select {
			case status := <-lobbylistener.Status:
				switch status.Category {
				case pubnub.PNConnectedCategory:
					res, _, err := pn.HereNow(). // Check if there is already a host in the game.
									Channels([]string{channel}).
									Execute()
					if err != nil {
						panic(err)
					}
					if res.Channels[0].Occupancy != 0 { // Player will be guest. Send username to host.
						data["guestusername"] = username
						pn.Publish().
							Channel(channel).
							Message(data).
							Execute()
					} else {
						fmt.Println("Waiting for guest...")
					}
				case pubnub.PNDisconnectedCategory:
					fmt.Println("Error: Connection to game lobby was lost.")
					os.Exit(0)
				}
			case message := <-lobbylistener.Message:
				if msg, ok := message.Message.(map[string]interface{}); ok {
					if val, ok := msg["guestusername"]; ok { // The host receives the guest username then the host sends the host username and starts a game.
						guestUsername = val.(string)
						data["hostusername"] = username
						pn.Publish().
							Channel(channel).
							Message(data).
							Execute()
					}
					if val, ok := msg["hostusername"]; ok { // When the guest receives the host username then the game is ready to start.
						hostUsername = val.(string)
						pn.RemoveListener(lobbylistener)
						countdown() // Countdown before starting game.
						startGame(hostUsername, guestUsername, pn)
						return
					}

				}
			}
		}
	}()
	pn.AddListener(lobbylistener)
	pn.Subscribe().
		Channels([]string{channel}).
		Execute()
}

func userInput(channel string, username string) (string, string) {
	var (
		newChannel  string
		newUsername string
		err         error
	)
	reader := bufio.NewReader(os.Stdin)
	if channel == "" { // Ask for channel name.
		fmt.Print("Enter Channel Name: ")
	} else {
		fmt.Print("Enter Channel Name (" + channel + "): ")
	}
	newChannel, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	newChannel = strings.Replace(newChannel, "\n", "", -1) // Convert CRLF to LF.
	if newChannel != "" {                                  // Use last channel name when the player does not provide a channel.
		channel = newChannel
	}
	if username == "" { // Ask for username.
		fmt.Print("Enter Username: ")
	} else {
		fmt.Print("Enter Username (" + username + "): ")
	}
	newUsername, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	newUsername = strings.Replace(newUsername, "\n", "", -1) // Convert CRLF to LF.
	if newUsername != "" {                                   // Use last username when the player does not provide a username.
		username = newUsername
	}
	if (channel == "") || (username == "") { // The player must have a channel and username.
		fmt.Println("You Must Provide a Channel and Username! ")
		userInput(channel, username) // Start over.
	}
	return channel, username
}
