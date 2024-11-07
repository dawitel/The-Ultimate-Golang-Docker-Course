package usecases

import (
	"errors"
	"math/rand"
	"time"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/entities"
)

type GameUseCaseImpl struct{}

func NewGameUseCase() GameUseCase {
	return &GameUseCaseImpl{}
}

func (g *GameUseCaseImpl) StartGame(maxAttempts int) entities.Game {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // Random number between 1 and 100
	return entities.Game{TargetNumber: target, MaxAttempts: maxAttempts, Attempts: 0}
}

func (g *GameUseCaseImpl) MakeGuess(game *entities.Game, guess int) (string, error) {
	if game.Attempts >= game.MaxAttempts {
		return "", errors.New("game over, maximum attempts reached")
	}

	game.Attempts++
	if guess < game.TargetNumber {
		return "Too low!", nil
	} else if guess > game.TargetNumber {
		return "Too high!", nil
	}
	return "Correct! You guessed the number!", nil
}
