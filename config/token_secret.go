package config

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateTokenSecret() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)

	if err != nil {
		log.Fatalf("Error generating token secret, %v", err)
	}

	return hex.EncodeToString(token)
}
