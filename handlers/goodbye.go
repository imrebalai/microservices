package handlers

import (
	"log"
	"net/http"
)

// inspided by Nic Jackson Part2 https://www.youtube.com/watch?v=hodOppKJm5Y&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}
