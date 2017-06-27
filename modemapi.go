package main

import (
	"log"
	"strconv"
	"strings"
	"sync"
)

var ModemAPI_Mutex sync.Mutex

func ModemAPI_Lock() {
	ModemAPI_Mutex.Lock()
}

func ModemAPI_Unlock() {
	ModemAPI_Mutex.Unlock()
}

func ModemAPI_GetBattery() (int, int) {
	Modem_WriteLine("AT+CBC")

	log.Println(Modem_ReadUntil('\n')) // blank
	info := Modem_ReadUntil('\n') // info
	log.Println(Modem_ReadUntil('\n')) // blank
	log.Println(Modem_ReadUntil('\n')) // OK

	parts := strings.Split(strings.TrimSpace(info), ",")
	percentage, _ := strconv.Atoi(parts[1])
	voltage, _ := strconv.Atoi(parts[2])

	return percentage, voltage
}

func ModemAPI_GetCarrier() (string) {
	Modem_WriteLine("AT+COPS?")

	log.Println(Modem_ReadUntil('\n')) // blank
	carrierLine := Modem_ReadUntil('\n') // +COPS: 0,0,"T-Mobile USA"
	log.Println(Modem_ReadUntil('\n')) // blank
	log.Println(Modem_ReadUntil('\n')) // OK

	s := strings.Split(carrierLine, "\"")
	return s[1][0:len(s[1])]
}

func ModemAPI_GetCCID() (string) {
	Modem_WriteLine("AT+CCID")

	log.Println(Modem_ReadUntil('\n')) // blank
	ccid := Modem_ReadUntil('\n') // ccid
	log.Println(Modem_ReadUntil('\n')) // blank
	log.Println(Modem_ReadUntil('\n')) // OK

	return strings.TrimSpace(ccid)
}