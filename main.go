package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

//	"github.com/gosuri/uiprogress"
//	term "github.com/nsf/termbox-go"
)

func newGame(channel string, username string) {
	reader := bufio.NewReader(os.Stdin)
	var (
		err error
		newChannel string
		newUsername string 
	)
	
	if (channel == "") { // Ask for channel name.
		fmt.Print("Enter Channel Name: ")
	} else {
		fmt.Print("Enter Channel Name ("+channel+"): ")
	}
	newChannel, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	newChannel = strings.Replace(newChannel, "\n", "", -1) // Convert CRLF to LF.
	if newChannel != "" { // Use last channel name when the user does not provide a channel. 
		channel = newChannel
	}

	if (username == "") { // Ask for username.
		fmt.Print("Enter Username: ")
	} else {
		fmt.Print("Enter Username ("+username+"): ")
	}
	newUsername, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	newUsername = strings.Replace(newUsername, "\n", "", -1) // Convert CRLF to LF.
	if newUsername != "" {  // Use last username when the user does not provide a username. 
		username = newUsername
	}

	if ((channel == "") || (username == "")) { // The user must have a channel and username.
		fmt.Println("You Must Provide a Channel and Username! ")
		newGame(channel, username) // Start over.
		return
	}



}

func main() {

	var (
		channel string
		username string
	)

	fmt.Println("Welcome to Space Race!")
	
	newGame(channel, username)

	os.Exit(0)
}
