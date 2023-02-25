package server

import (
	"bufio"
	"log"
	"net"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/game"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	"github.com/dineshgowda24/tic-tac-toe/source/player/human"
)

type Server struct {
	queue chan *PlayerMeta
}

func New() *Server {
	return &Server{
		queue: make(chan *PlayerMeta, 1),
	}
}

type PlayerMeta struct {
	name string
	conn net.Conn
}

func (s *Server) Serve() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started and listening on 8080")
	go s.consume()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		_, err = conn.Write([]byte("Welcome to tic tac toe XOXO!\n"))
		if err != nil {
			log.Println(err.Error())
			conn.Close()
			continue
		}

		_, err = conn.Write([]byte("Please enter your name\n"))
		if err != nil {
			log.Println(err.Error())
			conn.Close()
			continue
		}

		reader := bufio.NewReader(conn)
		name, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			conn.Close()
			continue
		}
		s.queue <- &PlayerMeta{
			name: name,
			conn: conn,
		}
	}
}

func (s *Server) consume() {
	var players []player.Player
	for meta := range s.queue {
		if len(players) == 0 {
			playerOne := human.New(player.X, meta.name, meta.conn)
			players = append(players, playerOne)
			continue
		}
		playerTwo := human.New(player.O, meta.name, meta.conn)
		brd, _ := board.NewBoard(3)
		gm := game.NewGame(brd, players[0], playerTwo)
		go gm.Start()

		players = []player.Player{}
	}
}
