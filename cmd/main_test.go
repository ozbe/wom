package main

import (
	"context"
	"strings"
	"testing"

	"github.com/ozbe/wom"
	"github.com/ozbe/wom/testutil"
)

func TestApp(t *testing.T) {
	input := wom.Input{
		Context: context.Background(),
	}

	cfg := config{
		getPod: func(_ context.Context, name string) string {
			return strings.Repeat(name, 2)
		},
	}
	app := newCliBuilder(cfg)

	testCases := map[string]struct {
		args        []string
		assertPanic bool
	}{
		"pod get": {
			args: []string{"test", "pod", "get", "jo"},
		},
		"svc get": {
			args: []string{"test", "pod", "get", "jo"},
		},
		"svc set": {
			args:        []string{"test", "svc", "set"},
			assertPanic: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output := &testutil.Output{}

			if tc.assertPanic {
				defer testutil.AssertPanic(t, output)()
			}

			app.Build(input, output).Run(tc.args)

			if tc.assertPanic {
				t.Fatal("Test did not panic")
			}

			testutil.AssertGoldenFile(t, *output)
		})
	}
}
