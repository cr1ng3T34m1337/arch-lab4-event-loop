package parser

import (
	"fmt"
	"strings"

	"github.com/cr1ng3T34m1337/arch-lab4-event-loop/pkg/engine"
)

func Parse(line string) []engine.Command {
	parts := strings.Fields(line)
	switch parts[0] {
	case "print":
		str := strings.Join(parts[1:], " ")
		return []engine.Command{&engine.PrintCommand{Arg: str}}
	case "sha1":
		cmds := make([]engine.Command, 0)
		for _, v := range parts[1:] {
			cmds = append(cmds, &engine.Sha1Command{Arg: v})
		}
		return cmds
	default:
		errMessage := fmt.Sprintf("no such instruction: %v", parts[0])
		return []engine.Command{&engine.PrintCommand{Arg: errMessage}}
	}
}
