package game

import (
	"github.com/thiago18l/hexagonal-go/pkg/clock"
	"github.com/thiago18l/hexagonal-go/pkg/random"
	port "github.com/thiago18l/hexagonal-go/src/core/ports"
)

type service struct {
	random     random.Random
	clock      clock.Clock
	repository port.GameRepository
}
