package main

import (
	"bufio"
	"log"
	"time"
	"strings"

	"github.com/tarm/serial"
)

var Modem_Port *serial.Port
var Modem_Reader *bufio.Reader

func Modem_ReadUntil(targetChar byte) string {
	reply, err := Modem_Reader.ReadBytes(targetChar)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("<--", strings.TrimSpace(string(reply)))
	return string(reply)
}

func Modem_WriteLine(line string) {
	log.Println("-->", line)
	Modem_Port.Write([]byte(line + "\n"))
}

func Modem_Discard() {
	i, _ := Modem_Reader.Discard(Modem_Reader.Buffered())
	log.Println("discard", i)
}

func Modem_GetReply(command string) string {
	Modem_Discard()
	Modem_WriteLine(command)
	time.Sleep(1*time.Millisecond)
	Modem_ReadUntil('\n') // skip first blank line
	time.Sleep(1*time.Millisecond)
	return strings.TrimSpace(Modem_ReadUntil('\n'))
}

func Modem_Init() {
	var err error

	c := &serial.Config{Name: config.Server.SerialPort, Baud: 115200, ReadTimeout: time.Second * 2}
	Modem_Port, err = serial.OpenPort(c)
	if err != nil {
		log.Println("Failed to open given port for modem:")
		log.Fatal(err)
	}

	log.Println("Initializing modem...")

	Modem_Port.Flush()
	Modem_Reader = bufio.NewReader(Modem_Port)

	// start up the auto-bauder
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("ATE0") // disable echo

	Modem_Port.Flush()

	time.Sleep(10*time.Millisecond)

	Modem_Reader.ReadBytes('\n')
	Modem_Reader.ReadBytes('\n')
	Modem_Reader.ReadBytes('\n')
	Modem_Reader.ReadBytes('\n')
	Modem_Reader.ReadBytes('\n')
	Modem_Discard()

	if Modem_GetReply("AT") != "OK" {
		log.Fatal("Didn't get an OK to AT!")
	}

	log.Printf("Connected to SIM800 on port %s", config.Server.SerialPort)
}