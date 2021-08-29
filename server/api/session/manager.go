package session

import (
	"crypto/rand"
	"time"
)

const CookieName = "id"
const ExpiresIn = 2 * time.Hour

type Manager interface {
	NewSession() (string, error)
}

type InMemoryManager struct {
	sessions map[string]time.Time
}

func (i InMemoryManager) NewSession() (string, error)  {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	sessionId := string(b)
	i.sessions[sessionId] = time.Now().Add(ExpiresIn)
	return sessionId, nil
}
