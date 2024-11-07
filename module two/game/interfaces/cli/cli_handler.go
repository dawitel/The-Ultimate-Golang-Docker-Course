package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/guessing-ame/usecases"
)

type CLIHandler struct {
    useCase usecases.GameUseCase
}

func NewCLIHandler(u usecases.GameUseCase) *CLIHandler {
    return &CLIHandler{useCase: u}
}

func (h *CLIHandler) Start(input io.Reader, output io.Writer) {
    game := h.useCase.StartGame(5)
    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)

    fmt.Fprintln(writer, "Welcome to the Guessing Game!")
    writer.Flush()

    for {
        fmt.Fprint(writer, "Enter your guess: ")
        writer.Flush()

        if !scanner.Scan() {
            fmt.Fprintln(writer, "Error reading input.")
            break
        }

        input := strings.TrimSpace(scanner.Text())
        guess, err := strconv.Atoi(input)
        if err != nil {
            fmt.Fprintln(writer, "Invalid input. Please enter a number.")
            writer.Flush()
            continue
        }

        response, err := h.useCase.MakeGuess(&game, guess)
        if err != nil {
            fmt.Fprintln(writer, err)
            writer.Flush()
            break
        }

        fmt.Fprintln(writer, response)
        writer.Flush()

        if response == "Correct! You guessed the number!" {
            break
        }
    }
}
