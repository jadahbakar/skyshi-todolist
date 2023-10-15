package activity

import (
	"strconv"
	"strings"

	"github.com/jadahbakar/skyshi-todolist/util/errorlib"
	"github.com/jadahbakar/skyshi-todolist/util/logger"
)

type srv struct {
	repo Repository
}

type Service interface {
	Create(*PostReq) (*Activity, error)
	Update(string, string) (*Activity, error)
	Delete(string) (int64, error)
	FindById(int64) (*Activity, error)
	FindActId(string) (*Activity, error)
	FindAll() ([]Activity, error)
}

func NewService(r Repository) Service {
	return &srv{repo: r}
}

func (s *srv) Create(req *PostReq) (*Activity, error) {
	trimTitle := strings.Trim(req.Title, " \t\n")
	trimEmail := strings.Trim(req.Email, " \t\n")
	if trimTitle == "" {
		return nil, errorlib.WrapErr(nil, errorlib.ErrorCodeInvalidArgument, "title cannot be null")
	}
	if trimEmail == "" {
		// return nil, errors.New("email cannot be null")
		return nil, errorlib.WrapErr(nil, errorlib.ErrorCodeInvalidArgument, "email cannot be null")
	}

	id, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	res, err := s.FindById(id)
	return res, err

}

func (s *srv) Update(param string, title string) (*Activity, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return nil, errorlib.WrapErr(nil, errorlib.ErrorCodeInvalidArgument, "error parsing id")
	}

	res, err := s.FindById(id)
	if err != nil {
		return nil, err
	}

	resid, err := s.repo.Update(int64(res.Id), title)
	if err != nil {
		return nil, err
	}

	result, err := s.FindById(resid)
	return result, err

}

func (s *srv) Delete(param string) (int64, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, errorlib.WrapErr(nil, errorlib.ErrorCodeInvalidArgument, "error parsing id")
	}

	_, err = s.FindById(id)
	if err != nil {
		return 0, err
	}

	resid, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return resid, err
}

func (s *srv) FindById(id int64) (*Activity, error) {
	res, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *srv) FindActId(param string) (*Activity, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error("error parsing id")
		return nil, errorlib.WrapErr(nil, errorlib.ErrorCodeInvalidArgument, "error parsing id")
	}

	res, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return res, err

}

func (s *srv) FindAll() ([]Activity, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		logger.Error("error get data")
		return nil, err
	}
	return res, nil
}
