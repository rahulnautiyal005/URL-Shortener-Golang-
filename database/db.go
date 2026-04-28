package database

import (
	"errors"
	"sync"
	"time"
	"url-shortener/models"
)

type Store interface {
	Save(url *models.URL) error
	GetByCode(code string) (*models.URL, error)
	IncrementClick(code string) error
	GetNextID() (int64, error)
}

type MemoryStore struct {
	mu   sync.RWMutex
	urls map[string]*models.URL
	nextID int64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls:   make(map[string]*models.URL),
		nextID: 1000, // Start with a non-zero ID
	}
}

func (s *MemoryStore) GetNextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	return id, nil
}

func (s *MemoryStore) Save(url *models.URL) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	url.CreatedAt = time.Now()
	s.urls[url.ShortCode] = url
	return nil
}

func (s *MemoryStore) GetByCode(code string) (*models.URL, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.urls[code]
	if !ok {
		return nil, errors.New("URL not found")
	}

	if !url.ExpiresAt.IsZero() && time.Now().After(url.ExpiresAt) {
		return nil, errors.New("URL has expired")
	}

	return url, nil
}

func (s *MemoryStore) IncrementClick(code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if url, ok := s.urls[code]; ok {
		url.ClickCount++
		return nil
	}
	return errors.New("URL not found")
}
