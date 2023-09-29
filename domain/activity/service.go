package activity

type srv struct {
	repo Repository
}

type Service interface {
	Create(*PostReq) (*Activity, error)
}

func NewService(r Repository) Service {
	return &srv{repo: r}
}

func (s *srv) Create(req *PostReq) (*Activity, error) {
	res, err := s.repo.Create(req)
	return res, err
}
