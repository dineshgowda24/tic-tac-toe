package main

import (
	"fmt"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/server"
)

func main() {
	srv := server.NewServer()
	srv.Serve()
}

//func main() {
//	for {
//		fmt.Println("Welcome to Tic Tac Toe XOXO!")
//		fmt.Println("Please select board size ie size X size - size=3,4,5,6")
//
//		// get size & player types
//		size := collectSize()
//		gameType := collectGameType()
//
//		// set up players
//		var playerOne, playerTwo player.Player
//		if gameType == "s" {
//			playerOne = smart_computer.NewSmartComputer(player.X)
//			fmt.Printf("%s will play with %s\n", playerOne.Name(), playerOne.Move().String())
//			playerTwo = human.NewHuman(player.O, collectPlayerName(), os.Stdin)
//			fmt.Printf("%s will play with %s\n", playerTwo.Name(), playerTwo.Move().String())
//
//		} else {
//			playerOne = human.NewHuman(player.X, collectPlayerName(), os.Stdin)
//			fmt.Printf("%s will play with %s\n", playerOne.Name(), player.X)
//			playerTwo = human.NewHuman(player.O, collectPlayerName(), os.Stdin)
//			fmt.Printf("%s will play with %s\n", playerTwo.Name(), playerTwo.Move().String())
//		}
//
//		// set up board
//		brd, err := board.NewBoard(size)
//		if err != nil {
//			fmt.Println("failed to initialize new board")
//			os.Exit(1)
//		}
//
//		//init game
//		gm := game.NewGame(brd, playerOne, playerTwo)
//		gm.Start()
//
//		if continuePlaying() {
//			continue
//		}
//		fmt.Println("See ya soon!")
//		return
//	}
//}

// collectSize returns size from STDIN
func collectSize() int {
	var size int
	for {
		fmt.Scanf("%d", &size)
		if size < 3 {
			fmt.Println("invalid size, that's ok try sizes greater than 2")
			continue
		}
		return size
	}
}

// collectGameType returns game type from STDIN
func collectGameType() string {
	var gameType string
	fmt.Println("Do you want to play single player or multi player?")
	fmt.Println("Select [S]ingle or [M]ulti-Player")
	for {
		fmt.Scanf("%s", &gameType)
		gameType = strings.ToLower(strings.TrimSpace(gameType))
		switch gameType {
		case "s", "m":
			return gameType
		default:
			fmt.Println("Hmm.. something doesn't look right, try again'")
		}
	}
}

// collectPlayerName returns move of player from STDIN
func collectPlayerName() string {
	var identifier string
	fmt.Println("Please enter your name")
	for {
		fmt.Scanf("%s", &identifier)
		identifier = strings.ToLower(strings.TrimSpace(identifier))
		switch len(identifier) {
		case 0:
			fmt.Println("Hmm.. something doesn't look right, try again'")
		default:
			return identifier
		}
	}
}

// continuePlaying asks for consent to play again after game is finished
func continuePlaying() bool {
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
