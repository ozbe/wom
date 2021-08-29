package main

import (
	"context"
	"os"

	"github.com/ozbe/wom"
)

func main() {
	ctx := context.Background()
	input := wom.Input{Context: ctx}
	output := wom.NewDefaultOutput()

	cfg := config{
		getPod: func(_ context.Context, name string) string {
			return name + "-pod"
		},
		getSvc: func(_ context.Context, name string) string {
			return name + "-svc"
		},
	}
	newCliBuilder(cfg).
		Build(input, output).
		Run(os.Args)
}

type config struct {
	getPod func(context.Context, string) string
	getSvc func(context.Context, string) string
}

func newCliBuilder(cfg config) wom.CliBuilder {
	return *wom.NewCliBuilder("test", "").
		Cmd("pod", "", CmdPod{
			getPod: cfg.getPod,
		}).
		Cmd("svc", "", CmdSvc{
			getSvc: cfg.getSvc,
		})
}
