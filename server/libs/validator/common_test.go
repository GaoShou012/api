package libs_validator

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestUsername(t *testing.T) {
	username := "abc*"
	if err := Username(username); err != nil {
		log.Error("username:", err)
	}
}

func TestPassword(t *testing.T) {
	password := "312s"
	if err := Password(password); err != nil {
		log.Error("password:", err)
	}
}

func TestPhone(t *testing.T) {
	phone := "138138001380"
	if err := Phone(phone); err != nil {
		log.Error("phone:", err)
	}
}
