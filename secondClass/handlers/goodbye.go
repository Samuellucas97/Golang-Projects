package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Println("Goodbye :(")

	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "The url is malformed", http.StatusBadRequest)
		return
	}

	res.Write([]byte("Bye Bye...\n"))
}