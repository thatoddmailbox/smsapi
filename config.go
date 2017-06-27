package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var config Config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int
	Token string
}

func Config_CreateNew() {
	newConfig := `# smsapi configuration
[server]
Port = 4725
Token = "hello"`
	err := ioutil.WriteFile("config.toml", []byte(newConfig), 0644)
	if err != nil {
		panic(err)
	}
}

func Config_Init() {
	if _, err := os.Stat("config.toml"); err != nil {
		Config_CreateNew() // create new config to be parsed
	}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
}