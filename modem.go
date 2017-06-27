package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

var Modem_Port *serial.Port

func Modem_ReadUntil(targetChar byte) string {
	outBuf := []byte{}
	for {
		buf := make([]byte, 128)
		n, err := Modem_Port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < n; i++ {
			c := buf[i]
			outBuf = append(outBuf, c)
			if c == targetChar {
				// done, found the target
				return string(outBuf)
			}
		}
	}
}

func Modem_WriteLine(line string) {
	Modem_Port.Write([]byte(line + "\n"))
}

func Modem_GetReply(command string) string {
	Modem_WriteLine(command)
	return Modem_ReadUntil('\n')
}

func Modem_Init() {
	var err error

	c := &serial.Config{Name: config.Server.SerialPort, Baud: 115200}
	Modem_Port, err = serial.OpenPort(c)
	if err != nil {
		log.Println("Failed to open given port for modem:")
		log.Fatal(err)
	}

	log.Println("Initializing modem...")

	Modem_Port.Flush()

	// start up the auto-bauder
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("AT")
	Modem_WriteLine("ATE0") // disable echo

	Modem_Port.Flush()

	time.Sleep(10*time.Millisecond)

	log.Println(Modem_GetReply("AT"))
	log.Println(Modem_GetReply("AT+COPS?"))
	log.Printf("Connected to SIM800 on port %s", config.Server.SerialPort)
}