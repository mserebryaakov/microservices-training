package handler

import (
	"fmt"
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

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Goodbye world")
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprint(rw, "Goodbye user")
}
