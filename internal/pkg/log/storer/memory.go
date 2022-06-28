package storer

import (
	"errors"
	"time"

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
	return toLogs(m.LogMap), nil
}

// GetAllByTime gets all the logs between the given time range
func (m *MemoryStore) GetAllByTime(from time.Time, to time.Time) ([]log.Log, error) {
	result := make([]log.Log, 0, len(m.LogMap))
	for _, row := range m.LogMap {
		if row.T.After(from) && row.T.Before(to) {
			result = append(result, row)
		}
	}

	return result, nil
}

// InsertAll creates new logs
func (m *MemoryStore) InsertAll(data []log.Log) ([]log.Log, error) {
	for _, row := range data {
		if _, ok := m.LogMap[row.ID.String()]; ok {
			// Deduplicate
			continue
		}
		m.LogMap[row.ID.String()] = row
	}

	return toLogs(m.LogMap), nil
}

func toLogs(input map[string]log.Log) []log.Log {
	result := make([]log.Log, 0, len(input))

	for _, row := range input {
		result = append(result, row)
	}

	return result
}
