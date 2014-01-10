package main

import (
	"fmt"
	"github.com/djworth/irc"
	"strings"
)

func dispatch(client *irc.IrcClient, message string) {

	if strings.Contains(message, " btc ") {
		price := CurrentBTCPrice("USD")
		client.SendMessage(fmt.Sprintf("Current BTC price is:", price))
	}

}

func main() {

	client := irc.NewIrcClient()

	client.Host = "irc.freenode.net"
	client.Port = 6667
	client.Channel = "#fireside"
	client.Nick = "badboy_gobot"
	client.CallBack = dispatch

	client.Join()

}
