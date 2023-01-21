package main

import "os"

func xdgDataDir() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome != "" {
		return dataHome
	}
	return os.Getenv("HOME") + "/.local/share"
}

func xdgConfigDir() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome != "" {
		return configHome
	}
	return os.Getenv("HOME") + "/.config"
}
