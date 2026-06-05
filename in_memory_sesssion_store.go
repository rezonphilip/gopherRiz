package gopherriz

import (
	"sync"
	"time"
)

type Session struct {
	UserId    string
	ExpiresAt time.Time
}

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func (s *SessionManager) Get(key string) (Session, bool) {
	s.mu.RLock()
	val, ok := s.sessions[key]
	s.mu.RUnlock()

	if !ok || time.Now().After(val.ExpiresAt) {
		return Session{}, false
	}

	return val, ok

}

func (s *SessionManager) Set(session Session, key string) {
	s.mu.Lock()
	s.sessions[key] = session
	s.mu.Unlock()
}
