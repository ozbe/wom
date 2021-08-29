package wom

import (
	"context"
	"fmt"
	"io"
	"os"

	cli "github.com/jawher/mow.cli"
)

type Input struct {
	Context context.Context
}

type Output interface {
	Print(...interface{})
	Error(...interface{})
	Fatal(int, ...interface{})
}

type Print func(...interface{})
type Fatal func(int, ...interface{})

type Exiter func(int)

type DefaultOutput struct {
	stdout io.Writer
	stderr io.Writer
	exiter Exiter
}

func NewDefaultOutput() DefaultOutput {
	return DefaultOutput{
		os.Stdin,
		os.Stderr,
		cli.Exit,
	}
}

func (o DefaultOutput) Print(a ...interface{}) {
	fmt.Fprintln(o.stdout, a...)
}

func (o DefaultOutput) Error(a ...interface{}) {
	fmt.Fprintln(o.stderr, a...)
}

func (o DefaultOutput) Fatal(code int, a ...interface{}) {
	o.Error(a...)
	o.exiter(1)
}

type cmdInitializer func(*cli.Cli, Input, Output)

type CliBuilder struct {
	name, desc string
	cmders     []cmdInitializer
}

func NewCliBuilder(name, desc string) *CliBuilder {
	return &CliBuilder{
		name: name,
		desc: desc,
	}
}

func (a *CliBuilder) Cmd(name, desc string, c Cmder) *CliBuilder {
	a.cmders = append(a.cmders, func(app *cli.Cli, i Input, o Output) {
		app.Command(name, desc, c.Cmd(i, o))
	})
	return a
}

type Cmder interface {
	Cmd(Input, Output) cli.CmdInitializer
}

func (a CliBuilder) Build(i Input, o Output) *cli.Cli {
	app := cli.App(a.name, a.desc)
	for _, r := range a.cmders {
		r(app, i, o)
	}
	return app
}
