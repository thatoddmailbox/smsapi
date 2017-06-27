package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes_Init() {
	router := httprouter.New()

	// root directory
	router.GET("/", APIRoute_Main)

	http.Handle("/", router)
}