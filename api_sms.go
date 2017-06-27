package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func APIRoute_SMS_Send(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	to := r.FormValue("to")
	msg := r.FormValue("msg")

	if to == "" || msg == "" {
		WriteJSON(w, ErrorResponse{"error", "missing_params"})
		return
	}

	ModemAPI_Lock()
	result := ModemAPI_SendSMS(to, msg)
	ModemAPI_Unlock()
	
	if result {
		WriteJSON(w, StatusResponse{"ok"})
	} else {
		WriteJSON(w, ErrorResponse{"error", "modem_unhappy"})
	}
}