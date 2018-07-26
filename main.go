package main

import (
	"fmt"
	"os"

	"github.com/gosuri/uiprogress"
	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	fmt.Println("Press SPACE to race!")
	uiprogress.Start()

	bar1 := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	bar1.PrependFunc(func(b *uiprogress.Bar) string {
		return "Host: "
	})

	bar2 := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	bar2.PrependFunc(func(b *uiprogress.Bar) string {
		return "Guest: "
	})

	fmt.Println("game:")

keyPressListenerLoop:
	for {
		event := term.PollEvent()
		switch {
		case event.Key == term.KeyEsc:
			term.Close()
			break keyPressListenerLoop
		case event.Key == term.KeySpace:
			bar2.Incr()
		case event.Key == term.KeyEnter:
			bar1.Incr()
		}
	}
	uiprogress.Stop()

	fmt.Println("Thanks for playing!")

	fmt.Println("New game?")

	os.Exit(0)
}
