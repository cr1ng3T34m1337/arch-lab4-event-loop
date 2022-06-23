package parser

import (
	"testing"

	"github.com/cr1ng3T34m1337/arch-lab4-event-loop/pkg/engine"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	res := Parse("print test data")
	expected := []engine.Command{&engine.PrintCommand{Arg: "test data"}}
	assert.Equal(t, expected, res)
	res = Parse("sha1 hash1 hash2")
	expected = []engine.Command{&engine.Sha1Command{Arg: "hash1"}, &engine.Sha1Command{Arg: "hash2"}}
	assert.Equal(t, expected, res)
	res = Parse("notExisting skjgdiuwhughu sjdfub iudgfuyg")
	expected = []engine.Command{&engine.PrintCommand{Arg: "no such instruction: notExisting"}}
	assert.Equal(t, expected, res)
}
