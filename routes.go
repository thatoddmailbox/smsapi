package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes_Init() {
	router := httprouter.New()

	// root directory
	router.GET("/", APIRoute_Main)

	// modem directory
	router.GET("/modem/battery", APIRoute_Modem_Battery)
	router.GET("/modem/carrier", APIRoute_Modem_Carrier)
	router.GET("/modem/ccid", APIRoute_Modem_CCID)

	http.Handle("/", router)
}