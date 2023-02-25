package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/game"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	"github.com/dineshgowda24/tic-tac-toe/source/player/human"
	"github.com/dineshgowda24/tic-tac-toe/source/player/smart_computer"
	"github.com/manifoldco/promptui"
)

func Cli() {
	for {
		fmt.Println("Welcome to Tic Tac Toe XOXO!")

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
			playerOne = smart_computer.NewSmartComputer(player.X)
			fmt.Printf("%s will play with %s\n", playerOne.Name(), playerOne.Move().String())
			name, _ := collectPlayerName()
			playerTwo = human.NewHuman(player.O, name, os.Stdin)
			fmt.Printf("%s will play with %s\n", playerTwo.Name(), playerTwo.Move().String())

		} else {
			nameOne, _ := collectPlayerName()
			playerOne = human.NewHuman(player.X, nameOne, os.Stdin)
			fmt.Printf("%s will play with %s\n", playerOne.Name(), player.X)
			nameTwo, _ := collectPlayerName()
			playerTwo = human.NewHuman(player.O, nameTwo, os.Stdin)
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
		Label: "Select board size",
		Items: []string{"3x3", "4x4", "5x5"},
	}

	_, boardSize, err := boardPrompt.Run()

	if err != nil {
		fmt.Printf("board prompt failed %v\n", err)
		return -1, errors.New("board prompt failed")
	}

	fmt.Printf("You choose %q\n", boardSize)

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
		Label: "Select game type",
		Items: []string{"Single", "MultiPlayer"},
	}

	_, gameType, err := gameTypePrompt.Run()

	if err != nil {
		fmt.Printf("Game type prompt failed %v\n", err)
		return "", errors.New("game type prompt failed")
	}

	switch gameType {
	case "Single":
		return "s", nil
	case "Multiplayer":
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
		return "", nil
	}

	return result, nil
}

// shouldContinuePlaying asks for consent to play again after game is finished
func shouldContinuePlaying() bool {
	fmt.Println("Do you want to play again?")
	fmt.Println("Enter [Y]es or [N]o")
	for {
		var stillPlaying string
		fmt.Scanf("%s", &stillPlaying)
		stillPlaying = strings.ToLower(strings.TrimSpace(stillPlaying))
		switch stillPlaying {
		case "n":
			return false
		case "y":
			return true
		default:
			fmt.Println("Hmm.. something doesn't look right, try again'")
		}
	}
}
