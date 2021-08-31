package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

const CookieName = "id"
const ExpiresIn = 2 * time.Hour

var ErrSessionIdNotFound = errors.New("provided session id could not be matched to session")
var ErrSessionExpired = errors.New("session has expired")

type Manager interface {
	NewSession() (string, error)
	ValidateSession(string) error
}

type inMemoryManager struct {
	sessions map[string]time.Time
	mutex *sync.Mutex
}

func NewInMemoryManager() Manager {
	return &inMemoryManager{
		sessions: make(map[string]time.Time),
		mutex: &sync.Mutex{},
	}
}

func (i *inMemoryManager) NewSession() (string, error)  {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	sessionId := base64.URLEncoding.EncodeToString(b)

	i.mutex.Lock()
	i.sessions[sessionId] = time.Now().Add(ExpiresIn)
	i.mutex.Unlock()

	return sessionId, nil
}

func (i *inMemoryManager) ValidateSession(sessionId string) error {
	i.mutex.Lock()
	exp, ok := i.sessions[sessionId]
	i.mutex.Unlock()
	if !ok {
		return ErrSessionIdNotFound
	}

	if time.Now().After(exp) {
		i.mutex.Lock()
		delete(i.sessions, sessionId)
		i.mutex.Unlock()

		return ErrSessionExpired
	}

	return nil
}
