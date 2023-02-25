package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/color"
	"github.com/dineshgowda24/tic-tac-toe/source/game"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	player_factory "github.com/dineshgowda24/tic-tac-toe/source/player/factory"

	"github.com/manifoldco/promptui"
)

func Cli() {
	for {
		fmt.Println(color.Green("Welcome to Tic Tac Toe XOXO!"))

		size := getBoardSize()
		gameType := getGameType()

		// set up players
		var playerOne, playerTwo player.Player
		if gameType == "s" {
			playerOne = getPlayerByDifficulty()
			fmt.Println(color.WhiteItalic(fmt.Sprintf("%s will play with %s", playerOne.Name(), playerOne.Move().String())))
			playerTwo = getHumanPlayer(player.O)
		} else {
			playerOne = getHumanPlayer(player.X)
			playerTwo = getHumanPlayer(player.O)
		}

		// set up board
		brd, err := board.NewBoard(size)
		if err != nil {
			fmt.Println("failed to initialize new board")
			os.Exit(1)
		}

		//init game
		gm := game.NewGame(brd, playerOne, playerTwo)
		gm.Start()

		if continuePlaying() {
			continue
		}
		fmt.Println(color.Green("See ya soon!"))
		return
	}
}

func getHumanPlayer(move player.Move) player.Player {
	name, _ := getPlayerName()
	plr := player_factory.NewPlayer(player.Human, name, move, os.Stdin)
	fmt.Println(color.WhiteItalic(fmt.Sprintf("%s will play with %s", plr.Name(), plr.Move().String())))
	return plr
}

// getBoardSize returns size from STDIN
func getBoardSize() int {
	boardPrompt := promptui.Select{
		Label: color.Green("Select board size"),
		Items: []string{"3x3", "4x4", "5x5"},
	}

	_, boardSize, _ := boardPrompt.Run()
	switch boardSize {
	case "3x3":
		return 3
	case "4x4":
		return 4
	case "5x5":
		return 5
	default:
		return 3
	}
}

// collectGameType returns game type from STDIN
func getGameType() string {
	gameTypePrompt := promptui.Select{
		Label: color.Green("Select game type"),
		Items: []string{"Single", "MultiPlayer"},
	}

	_, gameType, _ := gameTypePrompt.Run()
	switch gameType {
	case "Single":
		return "s"
	case "MultiPlayer":
		return "m"
	default:
		return "s"
	}
}

// getPlayerName returns move of player from STDIN
func getPlayerName() (string, error) {
	prompt := promptui.Prompt{
		Label: "What should I call you?",
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
		Validate: func(input string) error {
			input = strings.ToLower(strings.TrimSpace(input))
			if len(input) < 1 {
				return errors.New("invalid name")
			}
			return nil
		},
	}

	return prompt.Run()
}

// shouldContinuePlaying asks for consent to play again after game is finished
func continuePlaying() bool {
	playingPrompt := promptui.Select{
		Label: color.Green("Play again"),
		Items: []string{"Yes", "No"},
	}

	_, result, _ := playingPrompt.Run()
	switch result {
	case "Yes":
		return true
	case "No":
		return false
	default:
		return false
	}
}

// getDifficultyLevel returns difficulty level from STDIN
func getDifficultyLevel() string {
	prompt := promptui.Select{
		Label: color.Green("Select difficulty level"),
		Items: []string{"Beginner", "Expert"},
	}

	_, level, _ := prompt.Run()
	switch level {
	case "Beginner":
		return "b"
	case "Expert":
		return "e"
	default:
		return "b"
	}
}

func getPlayerByDifficulty() player.Player {
	level := getDifficultyLevel()
	switch level {
	case "b":
		return player_factory.NewPlayer(player.RandomComputer, "", player.X, nil)
	case "e":
		return player_factory.NewPlayer(player.SmartComputer, "", player.X, nil)
	default:
		return player_factory.NewPlayer(player.RandomComputer, "", player.X, nil)
	}
}
