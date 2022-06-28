package log

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type logStore interface {
	InsertAll(data []Log) ([]Log, error)
	GetAll() ([]Log, error)
	GetAllByTime(from time.Time, to time.Time) ([]Log, error)
	Get(id string) (*Log, error)
	Delete(id string) error
}

// swagger:model
type LogMsg struct {
	ID string `json:"id"`
	L  string `json:"@l"`
	M  string `json:"@m"`
	T  string `json:"@t"`
}

type Service struct {
	store logStore
}

func New(store logStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateLogs(data []Log) ([]Log, error) {
	return s.store.InsertAll(data)
}

func (s *Service) GetLogs(from time.Time, to time.Time) ([]Log, error) {
	if from.IsZero() && to.IsZero() {
		return s.store.GetAll()
	}

	return s.store.GetAllByTime(from, to)
}

func ToLogs(data []LogMsg) ([]Log, error) {
	result := []Log{}
	for _, row := range data {
		// Format date string to standard date layout
		dateString := strings.Replace(row.T, ":", ".", -1)
		date, err := time.Parse("2006-01-02 15.04.05.000000", dateString)
		if err != nil {
			return nil, err
		}

		log := Log{
			ID: uuid.MustParse(row.ID),
			L:  row.L,
			M:  row.M,
			T:  date,
		}

		result = append(result, log)
	}

	return result, nil
}
