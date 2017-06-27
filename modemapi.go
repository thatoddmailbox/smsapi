package main

import (
	"strconv"
	"strings"
	"sync"
	"time"
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

	Modem_ReadUntil('\n') // blank
	info := Modem_ReadUntil('\n') // info
	Modem_ReadUntil('\n') // blank
	Modem_ReadUntil('\n') // OK

	parts := strings.Split(strings.TrimSpace(info), ",")
	percentage, _ := strconv.Atoi(parts[1])
	voltage, _ := strconv.Atoi(parts[2])

	return percentage, voltage
}

func ModemAPI_GetCarrier() (string) {
	Modem_WriteLine("AT+COPS?")

	Modem_ReadUntil('\n') // blank
	carrierLine := Modem_ReadUntil('\n') // +COPS: 0,0,"T-Mobile USA"
	Modem_ReadUntil('\n') // blank
	Modem_ReadUntil('\n') // OK

	s := strings.Split(carrierLine, "\"")
	return s[1][0:len(s[1])]
}

func ModemAPI_GetCCID() (string) {
	Modem_WriteLine("AT+CCID")

	Modem_ReadUntil('\n') // blank
	ccid := Modem_ReadUntil('\n') // ccid
	Modem_ReadUntil('\n') // blank
	Modem_ReadUntil('\n') // OK

	return strings.TrimSpace(ccid)
}

func ModemAPI_SendSMS(to string, message string) bool {
	if Modem_GetReply("AT+CMGF=1") != "OK" {
		return false
	}

	// tell the modem who to send it to
	Modem_WriteLine("AT+CMGS=\"" + to + "\"")

	// write directly so there's no extra newline
	Modem_Port.Write([]byte(message))
	Modem_Port.Write([]byte{'\x1a'})

	Modem_ReadUntil('\n') // blank
	Modem_ReadUntil('\n') // >

	time.Sleep(2*time.Second) // wait for it to send

	Modem_ReadUntil('\n') // +CMGS: 17
	Modem_ReadUntil('\n') // blank
	if strings.TrimSpace(Modem_ReadUntil('\n')) == "OK" {
		return true
	} else {
		return false
	}
}