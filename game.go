package main

import (
	"fmt"
	"os"
	"time"

	pubnub "github.com/pubnub/go"

	"github.com/gosuri/uiprogress"
	term "github.com/nsf/termbox-go"
)

func countdown(hostName string, guestName string) {
	fmt.Println("The game is about to start between " + hostName + " and " + guestName + "!")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("3... ")
	time.Sleep(1 * time.Second)
	fmt.Print("2... ")
	time.Sleep(1 * time.Second)
	fmt.Print("1...  ")
	time.Sleep(1 * time.Second)
	fmt.Print("Race!!")
	time.Sleep(500 * time.Millisecond)
}

func startGame(isHost bool, lobby string, hostName string, guestName string, pn *pubnub.PubNub) {
	var (
		err      error
		spaced   bool
		progress int
	)
	data := make(map[string]interface{})

	countdown(hostName, guestName) // Countdown before starting game.

	gamelistener := pubnub.NewListener()

	err = term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	uiprogress.Start()
	hostBar := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	hostBar.AppendFunc(func(b *uiprogress.Bar) string {
		if isHost {
			return hostName + " (you)"
		}
		return hostName + " (host)"
	})
	guestBar := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	guestBar.AppendFunc(func(b *uiprogress.Bar) string {
		if !isHost {
			return guestName + " (you)"
		}
		return guestName + " (guest)"
	})

	go func() {
		for {
			select {
			case message := <-gamelistener.Message:
				if msg, ok := message.Message.(map[string]interface{}); ok {
					if val, ok := msg["guestProgress"]; ok { // The host receives the guest username then the host sends the host username and starts a game.
						guestBar.Set(int(val.(float64)))
					}
					if val, ok := msg["hostProgress"]; ok { // When the guest receives the host username then the game is ready to start.
						hostBar.Set(int(val.(float64)))
					}
				}
			}
		}
	}()

	pn.AddListener(gamelistener)
	pn.Subscribe().
		Channels([]string{lobby}).
		Execute()

	fmt.Println("Alternate pressing SPACE and the RIGHT ARROW KEY to race!")

keyPressListenerLoop:
	for {
		event := term.PollEvent()
		switch {
		case event.Key == term.KeyEsc:
			uiprogress.Stop()
			term.Close()
			pn.RemoveListener(gamelistener)
			pn.Unsubscribe().
				Channels([]string{lobby}).
				Execute()
			fmt.Println("")
			fmt.Println("Thanks for playing!")
			os.Exit(0)
			break keyPressListenerLoop
		case event.Key == term.KeySpace:
			if !spaced {
				progress = progress + 1
				if isHost {
					if hostBar.Width < hostBar.Total {
						data["hostProgress"] = progress
						pn.Publish().
							Channel(lobby).
							Message(data).
							Execute()
					}
				} else {
					if guestBar.Width < guestBar.Total {
						data["guestProgress"] = progress
						pn.Publish().
							Channel(lobby).
							Message(data).
							Execute()
					}
				}
				spaced = true
			}
		case event.Key == term.KeyArrowRight:
			spaced = false
		}
	}

}
