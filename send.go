package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/mail"
	"os"
	"os/exec"
	"strings"
)

func sendMessage(fromAddr, toAddr *mail.Address, body []byte) bool {
	cmd := exec.Command("/usr/sbin/sendmail", "-t", "-oi",
		"-f", fromAddr.Address, "-F", fromAddr.Name)
	cmd.Stdin = bytes.NewReader(body)

	if _, err := cmd.Output(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to send mail: %v (%s)\n", err, strings.TrimSpace(string(err.(*exec.ExitError).Stderr)))
		return false
	}
	return true
}

func send(c config, db *sql.DB) {
	fromAddr, err := mail.ParseAddress(c.Mail.From)
	if err != nil {
		fatal("failed to parse 'from' address: %v", err)
	}
	toAddr, err := mail.ParseAddress(c.Mail.To)
	if err != nil {
		fatal("failed to parse 'to' address: %v", err)
	}

	for _, msg := range unsentMessages(db) {
		if sendMessage(fromAddr, toAddr, formatMessage(msg, fromAddr, toAddr)) {
			markMessageAsSent(db, msg.id)
		}
	}
}
