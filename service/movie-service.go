package service

import (
	"errors"
	"github.com/matiascfgm/restAPI/entity"
	"github.com/matiascfgm/restAPI/repository"
	"math/rand"
	"strconv"
)

type MovieService interface {
	Validate(movie *entity.Movie) error
	Create(movie *entity.Movie) (*entity.Movie, error)
	FindAll() ([]entity.Movie, error)
}

type service struct{}

var (
	repo = repository.NewFirestorePostRepository()
)

func NewMovieService() MovieService {
	return &service{}
}

func (*service) Validate(movie *entity.Movie) error {
	if movie == nil {
		err := errors.New("The movie is empty")
		return err
	}
	if movie.Title == "" {
		err := errors.New("The movie title is empty")
		return err
	}
	return nil
}
func (*service) Create(movie *entity.Movie) (*entity.Movie, error) {
	movie.ID = strconv.Itoa(rand.Intn(1000))
	return repo.Save(movie)
}
func (*service) FindAll() ([]entity.Movie, error) {
	return repo.FindAll()
}
