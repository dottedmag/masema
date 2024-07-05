package main

import (
	"flag"
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
	} `toml:"mail"`

	App struct {
		ServerURL string `toml:"server_url"`
		Token     string `toml:"token"`
	} `toml:"app"`
}

func defaultConfigFile() string {
	return xdgConfigDir() + "/masema/config.toml"
}

func loadConfig(configFileName string) config {
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
	var configFileName, dbFileName string
	flag.StringVar(&configFileName, "config-file", defaultConfigFile(), "Config file location")
	flag.StringVar(&dbFileName, "db-file", defaultDBFile(), "DB file location")
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	c := loadConfig(configFileName)

	db, err := openDB(dbFileName)
	if err != nil {
		fatal("failed to open database: %s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fatal("failed to close database: %s", err)
		}
	}()

	switch flag.Arg(0) {
	case "fetch":
		fetch(c, db)
	case "send":
		send(c, db)
	default:
		usage()
	}
}
