package main

import (
	"encoding/json"
	"github.com/thoj/go-ircevent"
	"log"
	"net/http"
)

const ircserver = "irc.lebib.org:6667"
const nick = "bb2"

var ircconn irc.Connection

type Bib struct {
	Status bool `json:"status"`
}

var bib *Bib

func main() {
	// bib
	bib = &Bib{Status: false}
	// Serveur web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(bib)
	})
	// IRC
	ircconn := irc.IRC(nick, nick)
	err := ircconn.Connect(ircserver)
	/* Décommenter pour debug
	ircconn.VerboseCallbackHandler = true
	ircconn.Debug = true
	*/
	if err != nil {
		log.Fatal(err)
	}
	ircconn.AddCallback("PRIVMSG", func(e *irc.Event) {
		m := "Hein ? Liste des commandes: 'ouvre', 'ferme'"
		switch e.Message() {
		case "ouvre":
			bib.Status = true
			m = "Le bib est ouvert"
			break
		case "ferme":
			bib.Status = false
			m = "Le bib est fermé"
			break
		}
		ircconn.Privmsg(e.Nick, m)
	})
	log.Fatal(http.ListenAndServe(":55555", nil))
}
