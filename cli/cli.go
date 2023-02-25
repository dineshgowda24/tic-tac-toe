package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/game"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	player_factory "github.com/dineshgowda24/tic-tac-toe/source/player/factory"
	"github.com/fatih/color"

	"github.com/manifoldco/promptui"
)

var green func(a ...interface{}) string = color.New(color.FgGreen).Add(color.BgBlack).Add(color.Bold).SprintFunc()

func Cli() {
	for {
		fmt.Println(green("Welcome to Tic Tac Toe XOXO!"))

		size, err := getBoardSize()
		if err != nil {
			fmt.Print("Board prompt failed")
			os.Exit(1)
		}

		gameType, err := getGameType()
		if err != nil {
			fmt.Print("Game prompt failed")
			os.Exit(1)
		}

		// set up players
		var playerOne, playerTwo player.Player
		if gameType == "s" {
			playerOne = getPlayerByDifficulty()
			fmt.Printf("%s will play with %s\n", playerOne.Name(), playerOne.Move().String())
			name, _ := collectPlayerName()
			playerTwo = player_factory.NewPlayer(player.Human, name, player.O, os.Stdin)
			fmt.Printf("%s will play with %s\n", playerTwo.Name(), playerTwo.Move().String())

		} else {
			nameOne, _ := collectPlayerName()
			playerOne = player_factory.NewPlayer(player.Human, nameOne, player.X, os.Stdin)
			fmt.Printf("%s will play with %s\n", playerOne.Name(), player.X)
			nameTwo, _ := collectPlayerName()
			playerTwo = player_factory.NewPlayer(player.Human, nameTwo, player.O, os.Stdin)
			fmt.Printf("%s will play with %s\n", playerTwo.Name(), playerTwo.Move().String())
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

		if shouldContinuePlaying() {
			continue
		}
		fmt.Println("See ya soon!")
		return
	}
}

// getBoardSize returns size from STDIN
func getBoardSize() (int, error) {
	boardPrompt := promptui.Select{
		Label: green("Select board size"),
		Items: []string{"3x3", "4x4", "5x5"},
	}

	_, boardSize, err := boardPrompt.Run()

	if err != nil {
		fmt.Printf("board prompt failed %v\n", err)
		return -1, errors.New("board prompt failed")
	}

	switch boardSize {
	case "3x3":
		return 3, nil
	case "4x4":
		return 4, nil
	case "5x5":
		return 5, nil
	}
	return -1, errors.New("invalid board size")
}

// collectGameType returns game type from STDIN
func getGameType() (string, error) {
	gameTypePrompt := promptui.Select{
		Label: green("Select game type"),
		Items: []string{"Single", "MultiPlayer"},
	}

	_, gameType, _ := gameTypePrompt.Run()
	switch gameType {
	case "Single":
		return "s", nil
	case "MultiPlayer":
		return "m", nil
	}
	return "", errors.New("invalid game type")
}

// collectPlayerName returns move of player from STDIN
func collectPlayerName() (string, error) {

	validate := func(input string) error {
		input = strings.ToLower(strings.TrimSpace(input))
		if len(input) < 1 {
			return errors.New("invalid name")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Enter your name",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Name prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

// shouldContinuePlaying asks for consent to play again after game is finished
func shouldContinuePlaying() bool {
	playingPrompt := promptui.Select{
		Label: green("Play again"),
		Items: []string{"Yes", "No"},
	}

	_, result, _ := playingPrompt.Run()
	fmt.Println(result)
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
		Label: green("Select difficulty level"),
		Items: []string{"Beginner", "Expert"},
	}

	_, level, _ := prompt.Run()
	switch level {
	case "Beginner":
		return "b"
	case "Expert":
		return "e"
	default:
		return ""
	}
}

func getPlayerByDifficulty() player.Player {
	level := getDifficultyLevel()
	var playerOne player.Player
	if level == "b" {
		playerOne = player_factory.NewPlayer(player.RandomComputer, "", player.X, nil)
	} else if level == "e" {
		playerOne = player_factory.NewPlayer(player.SmartComputer, "", player.X, nil)
	} else {
		fmt.Print("difficulty level prompt failed")
		os.Exit(1)
	}
	return playerOne
}
