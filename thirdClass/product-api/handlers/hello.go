package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello :)")

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "The url is malformed", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hi %s! How've you been?\n", data)
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
