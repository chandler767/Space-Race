# Space Race

### Multiplayer racing game powered by [PubNub](https://www.pubnub.com/?devrel_gh=Space-Race).

<img alt="Multiplayer Gaming Revenue 2016" align="right" src="https://www.superdataresearch.com/wp-content/uploads/2016/08/SuperData-MMO-MOBA-Report-2016-1.png" height="330">

Playing a game alone is not nearly as fun as playing with your friends and that's why online games like massively multiplayer online games (MMOs) are a huge revenue stream that continues to grow rapidly every year. In June 2018 [consumers spent an estimated $9.1 billion digitally across all platforms [on MMOs]](https://www.superdataresearch.com/us-digital-games-market), compared to $7.9 billion last June. The design of online games can range from simple text-based environments to incredibly complex virtual worlds. A few other examples of popular online multiplayer games types include first-person shooters (FPS),  real-time strategy games (RTS), and multiplayer online battle arena games (MOBA). 

This project is intended to show how PubNub can be used to create a simple terminal based game. [PubNub is perfect for powering multiplayer games](https://pubnub.com/multiplayer-gaming/?devrel_gh=Space-Race) because it makes it easier to build a fast and scalable multiplayer game. In this game PubNub is used to manage the players in the game lobby and to transmit game data. The game is built in Go using the [PubNub Go SDK](https://www.pubnub.com/docs/go/pubnub-go-sdk?devrel_gh=Space-Race) and is cross platform. You can learn more about how this game was built from the blog post (coming soon).

[<img alt="PubNub usage in Space Race" align="center" src="https://i.imgur.com/ODdp1kY.png" height="150">](https://pubnub.com/?devrel_gh=Space-Race)

## Getting Started

### Quick Start

The fastest way to get started is to [download and run the appropriate binary for your OS](https://github.com/chandler767/Space-Race/releases/latest). You don't need to install Go or any dependencies to run the game. See the 'How To Play' section below to get racing.

[<img alt="Space Race Demo" align="center" src="https://i.imgur.com/CROuwN6.gif" height="150">](https://pubnub.com/?devrel_gh=Space-Race)

### Building From Source

- Want to learn more about this project or build a clone from scratch? Check out the blog post (coming soon).

- You’ll first need to sign up for a [PubNub account](https://dashboard.pubnub.com/signup/?devrel_gh=Space-Race). Once you sign up, you can get your unique PubNub keys from the [PubNub Developer Portal](https://admin.pubnub.com/?devrel_gh=Space-Race).

<a href="https://dashboard.pubnub.com/signup?devrel_gh=Space-Race">
    <img alt="PubNub Signup" src="https://i.imgur.com/og5DDjf.png" width=260 height=97/>
</a>

1. Install the latest version of [Go](https://golang.org/) and setup your $GOPATH.

2. Download the dependencies.
```bash
go get github.com/pubnub/go
go get github.com/gosuri/uiprogress
go get github.com/nsf/termbox-go
```

3. Clone the Repo.
```bash
git clone https://github.com/chandler767/Space-Race
```

4. Get your unique PubNub keys from the [PubNub Developer Portal](https://admin.pubnub.com/?devrel_gh=Space-Race). If you don't have a PubNub account, you can [sign up for a PubNub account](https://dashboard.pubnub.com/signup/?devrel_gh=Space-Race) for free.

5. Open [main.go](https://github.com/chandler767/Space-Race/blob/master/main.go) and replace "pub-key" and "sub-key" with your keys.

6. Build and run the game.
```bash
make run
```
OR
```bash
make build
./space-race
```

[<img alt="Space Race Game" align="center" src="https://i.imgur.com/z9RMA6K.png" height="150">](https://pubnub.com/?devrel_gh=Space-Race)

## How To Play

1. The first player to enter a lobby becomes the host. The second player to enter becomes the guest.

2. Both players need to enter the same lobby name.

3. When both players have joined a lobby the game will start after a 3 second delay (get your fingers ready).

4. After the game starts, alternate pressing SPACE and the RIGHT ARROW KEY to advance your progress bar.

5. The first player to 100% wins the game.

6. You can leave the game while playing by pressing ESC. If you leave the game then the other player wins automatically. 

[<img alt="Space Race Game" align="center" src="https://i.imgur.com/Rml08Qq.png" height="330">](https://pubnub.com/?devrel_gh=Space-Race)

### Limitations 

- If the terminal window is too small, or if the font is too big, the progress bars won't render correctly. Make the window bigger and then press command+k to refresh. 
- The lobby is kept simple by design so it can be used as a seed project. The lobby may not always be able to start a game in some edge cases. If you have problems starting a game try restarting the game for both players and use a new lobby name. 
