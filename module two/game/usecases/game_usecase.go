package usecases

import (
	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/entities"
)

type GameUseCase interface {
	StartGame(maxAttempts int) entities.Game
	MakeGuess(game *entities.Game, guess int) (string, error)
}
