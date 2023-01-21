package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: masema [fetch|send]\n")
	os.Exit(2)
}

type config struct {
	Mail struct {
		From string `toml:"from"`
		To   string `toml:"to"`
	} `toml"mail"`

	App struct {
		ServerURL string `toml:"server_url"`
		Token     string `toml:"token"`
	} `toml:"app"`
}

func loadConfig() config {
	configFileName := xdgConfigDir() + "/masema/config.toml"
	fh, err := os.Open(configFileName)
	if err != nil {
		fatal("failed to open config file: %s", err)
	}
	defer fh.Close()

	var c config

	if _, err := toml.NewDecoder(fh).Decode(&c); err != nil {
		fatal("failed to parse config file %s: %s", configFileName, err)
	}

	return c
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	c := loadConfig()

	db, err := openDB()
	if err != nil {
		fatal("failed to open database: %s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fatal("failed to close database: %s", err)
		}
	}()

	switch os.Args[1] {
	case "fetch":
		fetch(c, db)
	case "send":
		send(c, db)
	default:
		usage()
	}
}
