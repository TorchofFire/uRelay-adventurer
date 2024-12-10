package connections

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/nacl/sign"
)

func unlockSignedMessage(publicKey, encryptedMessage string) (string, time.Time, error) {
	pubKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to decode publicKey: %w", err)
	}
	encMessage, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to decode encryptedMessage: %w", err)
	}

	openedMessage, ok := sign.Open(nil, encMessage, (*[32]byte)(pubKey))
	if !ok {
		return "", time.Time{}, errors.New("failed to verify signed message")
	}
	unlockedMessage := string(openedMessage)
	parts := strings.SplitN(unlockedMessage, "|", 2)
	if len(parts) < 2 {
		return "", time.Time{}, errors.New("invalid message format: expected 'timestamp|message'")
	}

	stringTimestamp, payload := parts[0], parts[1]

	timestamp, err := strconv.ParseInt(stringTimestamp, 10, 64)
	if err != nil {
		return "", time.Time{}, errors.New("invalid timestamp: expected a unix timestamp")
	}

	messageTime := time.Unix(timestamp, 0)

	return payload, messageTime, nil
}

func signMessage(privateKey, message string) (string, error) {
	privKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode privateKey: %w", err)
	}
	if len(privKey) != 64 {
		return "", errors.New("invalid privateKey length: expected 64 bytes")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fullMessage := fmt.Sprintf("%s|%s", timestamp, message)

	signedMessage := sign.Sign(nil, []byte(fullMessage), (*[64]byte)(privKey))

	return base64.StdEncoding.EncodeToString(signedMessage), nil
}
