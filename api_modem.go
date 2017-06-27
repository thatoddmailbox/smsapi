package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BatteryResponse struct {
	Status string `json:"status"`
	Percentage int `json:"percentage"`
	Voltage int `json:"voltage"`
}

type CarrierResponse struct {
	Status string `json:"status"`
	Carrier string `json:"carrier"`
}

type CCIDResponse struct {
	Status string `json:"status"`
	CCID string `json:"ccid"`
}

func APIRoute_Modem_Battery(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ModemAPI_Lock()
	percentage, voltage := ModemAPI_GetBattery()
	ModemAPI_Unlock()
	
	WriteJSON(w, BatteryResponse{"ok", percentage, voltage})
}

func APIRoute_Modem_Carrier(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ModemAPI_Lock()
	carrier := ModemAPI_GetCarrier()
	ModemAPI_Unlock()
	
	WriteJSON(w, CarrierResponse{"ok", carrier})
}


func APIRoute_Modem_CCID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ModemAPI_Lock()
	ccid := ModemAPI_GetCCID()
	ModemAPI_Unlock()
	
	WriteJSON(w, CCIDResponse{"ok", ccid})
}