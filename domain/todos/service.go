package todos

import (
	"strconv"

	"github.com/jadahbakar/skyshi-todolist/util/logger"
)

type srv struct {
	repo Repository
}

type Service interface {
	Create(*PostReq) (*Todo, error)
	FindById(int64) (*Todo, error)
	Update(string, *PatchReq) (*Todo, error)
	Delete(string) (int64, error)
	FindAll(string) ([]Todo, error)
	FindTodoById(string) (*Todo, error)
}

func NewService(r Repository) Service {
	return &srv{repo: r}
}

func (s *srv) Create(req *PostReq) (*Todo, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	res, err := s.FindById(id)
	return res, err

}

func (s *srv) FindById(id int64) (*Todo, error) {
	res, err := s.repo.GetById(id)
	return res, err
}

func (s *srv) Update(param string, req *PatchReq) (*Todo, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return nil, err
	}

	resid, err := s.repo.Update(id, req)
	if err != nil {
		return nil, err
	}
	res, err := s.FindById(resid)
	return res, err

}

func (s *srv) Delete(param string) (int64, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return 0, err
	}

	resid, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return resid, err
}

func (s *srv) FindAll(param string) ([]Todo, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return nil, err
	}

	res, err := s.repo.GetAll(id)
	if err != nil {
		logger.Error("error get data")
		return nil, err
	}
	return res, nil
}

func (s *srv) FindTodoById(param string) (*Todo, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return nil, err
	}

	res, err := s.FindById(id)
	return res, err
}
