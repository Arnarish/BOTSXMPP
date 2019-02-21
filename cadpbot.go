package cadpbot

import (
	"log"
)

// interface to connect to a chat service, listen for messages and send replies to them
type Bot interface {
	name() string
	send(msg string)
	reply(orig Message, msg string)
	connect() error
	listen() chan Message
	setLogger(*log.Logger)
	log(msg string)
}

// interface for an individual message as well as metadata
type Message interface {
	Body() string
	From() string
	Room() string
}

type Plugin interface {
	Name() string
	Execute(msg Message, bot Bot) error
}

type CADPbot struct {
	Bot
	Plugins []Plugin
}

func (g *CADPbot) MyBot() Bot {
	return g.Bot
}
