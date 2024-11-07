package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/interfaces/cli"
	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/usecases"
)

func TestCLIHandler_Start(t *testing.T) {
	input := "50\n75\n25\n"
	output := &bytes.Buffer{}

	gameUseCase := usecases.NewGameUseCase()
	cliHandler := cli.NewCLIHandler(gameUseCase)

	cliHandler.Start(strings.NewReader(input), output)

	if !strings.Contains(output.String(), "Welcome to the Guessing Game!") {
		t.Errorf("Expected welcome message, got: %s", output.String())
	}
}
