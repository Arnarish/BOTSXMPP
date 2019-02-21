package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	//TODO push to github and import the plugins

	cb "github.com/arnarish/botsxmpp"
)

var host, user, pass, room, name, protocol, logfile string
var crons strslice

func init() {
	flag.StringVar(&host, "host", "localhost:5222", "Hostname:port of the server")
	flag.StringVar(&user, "username", "mybot", "Username to the server(e.g. foo@bar.com)")
	flag.StringVar(&pass, "password", "allyourbase", "Password to log on to the server")
	flag.StringVar(&room, "room", "", "Room to join(e.g. #foobar@bar.com)")
	flag.StringVar(&name, "name", "cadpbot", "name of your bot")
	flag.StringVar(&protocol, "protocol", "xmpp", "Protocol to use")
	flag.StringVar(&logfile, "logfile", "/tmp/log", "path to log file")
	flag.Var(&crons, "job", "list of jobs")
}

func createBot(Plugins []cb.Plugin) cb.cadpbot {
	var bot cb.cadpbot
	bot = cb.cadpbot{
		xmpp.new(pass, room, name),
		Plugins,
	}
	return bot
}

func execPlugin(p cb.Plugin, m cb.Message, b cb.bot) {
	err := p.Execute(m, b)
	if err != nil {
		b.log(p.name() + " -> " + err.Error())
	}
}

func main() {
	flag.parse()
	chatlog := chatlog.chatLog{filename: logfile}

	plugins := cb.Plugin{
		//TODO fill this
	}

	bot := createBot(plugins)
	err := bot.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	for _, crn := range crons {
		parts := strings.Split(crn, "|")
		cron.NewCron(parts[2], crn, bot)
	}

	var msg cb.Message
	var plugin cb.Plugin
	for msg = range bot.Listen() {
		for _, plugin = range bot.plugins {
			go execPlugin(plugin, msg, bot)
		}
	}
}

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%s", s)
}

func (s *strslice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
