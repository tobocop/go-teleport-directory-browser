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

func NewInMemoryManager() Manager {
	return &inMemoryManager{
		sessions: make(map[string]time.Time),
	}
}

type inMemoryManager struct {
	sessions map[string]time.Time
}

func (i inMemoryManager) NewSession() (string, error)  {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	sessionId := string(b)
	// TODO: Probably need to lock write / read with a mutex
	i.sessions[sessionId] = time.Now().Add(ExpiresIn)
	return sessionId, nil
}
