package topic

import (
	"fmt"
	"sync"
)

type InMemoryTopic struct {
	id   string
	mut  sync.RWMutex
	data [][]byte
}

var _ Topic = (*InMemoryTopic)(nil)

func NewInMemoryTopic(id string) *InMemoryTopic {
	return &InMemoryTopic{
		data: make([][]byte, 0),
		id:   id,
	}
}

func (s *InMemoryTopic) Insert(b []byte) (int, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *InMemoryTopic) Extract(offset int) ([]byte, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	if offset < 0 {
		return nil, fmt.Errorf(
			"IndexOutOfBound :: offset must be non-negative, provided (%d)",
			offset,
		)
	}

	if offset >= len(s.data) {
		return nil, fmt.Errorf(
			"IndexOutOfBound :: provided (%d), length (%d)",
			offset, len(s.data),
		)
	}

	return s.data[offset], nil
}

func (s *InMemoryTopic) ExtractLatest() ([]byte, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.data[len(s.data)-1], nil
}
