package storer

import (
	"errors"

	"github.com/wanghantao11/log-receiver/internal/pkg/log"
)

type MemoryStore struct {
	LogMap map[string]log.Log
}

func NewMemory() *MemoryStore {
	return &MemoryStore{
		LogMap: map[string]log.Log{},
	}
}

// Delete a specific log
func (m *MemoryStore) Delete(id string) error {
	_, found := m.LogMap[id]
	if !found {
		return errors.New("log not found")
	}
	delete(m.LogMap, id)
	return nil
}

// Get gets a specific log by id
func (m *MemoryStore) Get(id string) (*log.Log, error) {
	row, found := m.LogMap[id]
	if !found {
		return nil, errors.New("log not found")
	}
	return &row, nil
}

// GetAll gets all the logs
func (m *MemoryStore) GetAll() ([]log.Log, error) {
	values := make([]log.Log, 0, len(m.LogMap))
	return values, nil
}

// InsertAll creates new logs
func (m *MemoryStore) InsertAll(data []log.Log) error {
	for _, row := range data {
		m.LogMap[row.ID.String()] = row
	}
	return nil
}
