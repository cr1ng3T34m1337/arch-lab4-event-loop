package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cr1ng3T34m1337/arch-lab4-event-loop/pkg/engine"
	"github.com/cr1ng3T34m1337/arch-lab4-event-loop/pkg/parser"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open("instructions.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmds := parser.Parse(commandLine)
			for _, cmd := range cmds {
				eventLoop.Post(cmd)
			}
		}
	} else {
		log.Fatal(err)
	}
	eventLoop.AwaitFinish()
}
