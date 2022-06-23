package engine

import (
	"crypto/sha1"
	"fmt"
)

type Command interface {
	Execute(handler Handler)
}

type PrintCommand struct {
	Arg string
}

func (c *PrintCommand) Execute(loop Handler) {
	fmt.Println(c.Arg)
}

type Sha1Command struct {
	Arg string
}

func (c *Sha1Command) Execute(loop Handler) {
	hash := sha1.New()
	hash.Write([]byte(c.Arg))
	sum := hash.Sum(nil)
	line := fmt.Sprintf("sha1 for \"%v\": %v", c.Arg, string(sum))
	loop.Post(&PrintCommand{Arg: line})
}
