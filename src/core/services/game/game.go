package game

import (
	"errors"

	"github.com/thiago18l/hexagonal-go/pkg/clock"
	"github.com/thiago18l/hexagonal-go/pkg/random"
	"github.com/thiago18l/hexagonal-go/src/core/domain"
	port "github.com/thiago18l/hexagonal-go/src/core/ports"
)

type service struct {
	random     random.Random
	clock      clock.Clock
	repository port.GameRepository
}

func NewService(random random.Random, clock clock.Clock, repository port.GameRepository) *service {
	return &service{random: random, clock: clock, repository: repository}
}

// Get retrieves the game belonging to the userID given and with gameID given
func (srv *service) Get(userID string, gameID string) (domain.Game, error) {
	game, err := srv.repository.Get(userID, gameID)
	if err != nil {
		return domain.Game{}, errors.New("internal error")
	}
	if game == nil {
		return domain.Game{}, errors.New("not found")
	}
	return *game, nil
}

// GetAll retrieves all the games belonging to the userID given
func (srv *service) GetAll(userID string) ([]domain.Game, error) {
	games, err := srv.repository.GetAll(userID)
	if err != nil {
		return nil, errors.New("internal error")
	}
	return games, nil
}

// Create creates a new game for the user and settings given
func (srv *service) Create(userID string, settings domain.GameSettings) (domain.Game, error) {
	game := domain.Game{
		ID:       srv.random.GenerateID(),
		UserID:   userID,
		Settings: settings,
		Board:    domain.NewEmptyBoard(settings.Rows, settings.Columns),
		State:    domain.GameStateNew,
	}
	if err := srv.repository.Save(game); err != nil {
		return domain.Game{}, errors.New("failed at saving game into repository")
	}
	return game, nil
}
