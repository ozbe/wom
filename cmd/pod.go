package main

import (
	"context"

	cli "github.com/jawher/mow.cli"
	"github.com/ozbe/wom"
)

type CmdPod struct {
	getPod func(context.Context, string) string
}

func (c CmdPod) Cmd(i wom.Input, o wom.Output) cli.CmdInitializer {
	return func(cmd *cli.Cmd) {
		cmd.Command("get", "", func(cmd *cli.Cmd) {
			name := cmd.StringArg("NAME", "", "")
			cmd.Action = func() {
				c.get(i.Context, o.Print, *name)
			}
		})
	}
}

func (c CmdPod) get(ctx context.Context, print wom.Print, name string) {
	print(c.getPod(ctx, name))
}
