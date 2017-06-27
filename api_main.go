package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func APIRoute_Main(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "hello there\n\ni am a server\n\ni can send text messages")
}