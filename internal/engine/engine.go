package engine

import (
	"fmt"
	"github.com/paezao/chess-engine-go/internal/logger"
	"strings"
)

type State struct {
	Quit    bool
	UCIMode bool
}

func Engine(fromUI, toUI chan string) {
	state := State{Quit: false, UCIMode: false}

	var uiCmd string

	for !state.Quit {
		select {
		case uiCmd = <-fromUI:
			logger.Info("Input -> Engine: '" + uiCmd + "'")
		}

		uiCmd := strings.Split(uiCmd, " ")
		uiCmd[0] = strings.TrimSpace(strings.ToLower(uiCmd[0]))

		switch uiCmd[0] {
		case "uci":
			handleUci(&state, toUI)
		case "debug":
			fmt.Println("UI to Engine -> debug")
		case "isready":
			fmt.Println("UI to Engine -> isready")
		case "setoption":
			fmt.Println("UI to Engine -> setoption")
		case "register":
			fmt.Println("UI to Engine -> register")
		case "ucinewgame":
			fmt.Println("UI to Engine -> ucinewgame")
		case "position":
			fmt.Println("UI to Engine -> position")
		case "go":
			fmt.Println("UI to Engine -> go")
		case "stop":
			fmt.Println("UI to Engine -> stop")
		case "ponderhit":
			fmt.Println("UI to Engine -> ponderhit")
		case "quit", "q":
			fmt.Println("UI to Engine -> quit")
			state.Quit = true
			continue
		}
	}
	return
}

func handleUci(state *State, toUI chan string) {
	state.UCIMode = true
	toUI <- "id name Chess Engine GO"
	toUI <- "id author Marcelo Paez Sequeira"
	toUI <- "uciok"
}
