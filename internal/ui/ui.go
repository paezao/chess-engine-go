package ui

import (
	"bufio"
	"fmt"
	"github.com/paezao/chess-engine-go/internal/logger"
	"io"
	"os"
	"strings"
)

func UI() (from, to chan string) {
	from = make(chan string)
	to = make(chan string)

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				logger.Info("stdin -> Input: '" + text + "'")
				from <- text
			}
		}
	}()

	go func() {
		for output := range to {
			fmt.Println(output)
		}
	}()

	return
}
