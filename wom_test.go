package wom

import (
	"bytes"
	"errors"
	"testing"

	cli "github.com/jawher/mow.cli"
	"github.com/ozbe/wom/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDefaultOutput(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var exitCallCount int
	code := 1

	o := DefaultOutput{
		stdout: &stdout,
		stderr: &stderr,
		exiter: func(c int) {
			assert.Equal(t, code, c)
			exitCallCount++
		},
	}

	o.Print("foo", "bar")
	assert.Equal(t, "foo bar\n", stdout.String())

	o.Error("foo", "baz")
	assert.Equal(t, "foo baz\n", stderr.String())

	stderr.Reset()
	o.Fatal(code, errors.New("unexpected error"))
	assert.Equal(t, 1, exitCallCount, "exit call count")
	assert.Equal(t, "unexpected error\n", stderr.String())
}

func ExampleCliBuilder() {
	panic("not implemented")
}

func TestCliBuilder(t *testing.T) {
	input := Input{}
	output := testutil.Output{}

	cb := NewCliBuilder("name", "desc")

	cb.Cmd("get", "", funcCmdr(func(_ Input, o Output) cli.CmdInitializer {
		return cli.ActionCommand(func() {
			o.Print("get")
		})
	}))

	cb.Build(input, &output).Run([]string{"name", "get"})

	testutil.AssertGoldenFile(t, output)
}

type funcCmdr func(Input, Output) cli.CmdInitializer

func (c funcCmdr) Cmd(i Input, o Output) cli.CmdInitializer {
	return c(i, o)
}
