package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func APIRoute(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("token") != config.Server.Token {
			WriteJSON(w, ErrorResponse{"error", "token_invalid"})
			return
		}

		h(w, r, ps)
	}
}

func Routes_Init() {
	router := httprouter.New()

	// root directory
	router.GET("/", APIRoute(APIRoute_Main))

	// modem directory
	router.GET("/modem/battery", APIRoute(APIRoute_Modem_Battery))
	router.GET("/modem/carrier", APIRoute(APIRoute_Modem_Carrier))
	router.GET("/modem/ccid", APIRoute(APIRoute_Modem_CCID))

	// sms directory
	router.GET("/sms/send", APIRoute(APIRoute_SMS_Send))

	http.Handle("/", router)
}