package main

import (
	"github.com/paezao/chess-engine-go/internal/engine"
	"github.com/paezao/chess-engine-go/internal/logger"
	"github.com/paezao/chess-engine-go/internal/ui"
)

func main() {
	logger.InitLogger()
	logger.Info("Starting Chess Engine")
	fromUI, toUI := ui.UI()
	engine.Engine(fromUI, toUI)
}
