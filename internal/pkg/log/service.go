package log

type logStore interface {
	InsertAll(data []Log) error
	GetAll() ([]Log, error)
	Get(id string) (*Log, error)
	Delete(id string) error
}

type Service struct {
	store logStore
}

func New(store logStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) AddLogs(data []Log) error {
	return s.store.InsertAll(data)
}
