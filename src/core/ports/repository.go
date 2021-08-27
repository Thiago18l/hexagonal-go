package port

import "github.com/thiago18l/hexagonal-go/src/core/domain"

type GameRepository interface {
	Get(userID string, gameID string) (*domain.Game, error)
	GetAll(userID string) ([]domain.Game, error)
	Save(game domain.Game) error
}
