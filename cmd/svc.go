package main

import (
	"context"
	"errors"

	cli "github.com/jawher/mow.cli"
	"github.com/ozbe/wom"
)

type CmdSvc struct {
	getSvc func(context.Context, string) string
}

func (c CmdSvc) Cmd(i wom.Input, o wom.Output) cli.CmdInitializer {
	return func(cmd *cli.Cmd) {
		cmd.Command("get", "", func(cmd *cli.Cmd) {
			name := cmd.StringArg("NAME", "", "")
			cmd.Action = func() {
				c.get(i.Context, o, *name)
			}
		})
		cmd.Command("set", "", cli.ActionCommand(func() {
			o.Fatal(1, errors.New("not implemented"))
		}))
	}
}

func (c CmdSvc) get(ctx context.Context, o wom.Output, name string) {
	o.Print(c.getSvc(ctx, name))
}
