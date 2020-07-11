package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// inspided by Nic Jackson Part2 https://www.youtube.com/watch?v=hodOppKJm5Y&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s \n", d)
	fmt.Fprintf(rw, "Hello %s \n", d)

}
