package main

import (
	"os"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/interfaces/cli"
	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/usecases"
)

func main() {
    gameUseCase := usecases.NewGameUseCase()
    cliHandler := cli.NewCLIHandler(gameUseCase)
    cliHandler.Start(os.Stdin, os.Stdout)
}
