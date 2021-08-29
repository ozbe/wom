package testutil

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertPanic(t *testing.T, output *Output) func() {
	t.Helper()

	return func() {
		r := recover()
		if r == nil {
			t.Fatal("Test did not panic")
		}
		AssertGoldenFile(t, *output)
	}
}

var update = flag.Bool("update", false, "update golden files")

func AssertGoldenFile(t *testing.T, output Output) {
	t.Helper()

	got := output.buf.Bytes()
	goldenFile := "testdata/" + t.Name() + ".golden"

	if *update {
		path := filepath.Dir(goldenFile)
		_ = os.MkdirAll(path, 0755)

		err := os.WriteFile(goldenFile, got, 0777)
		if err != nil {
			t.Fatalf("Error writing golden file %s: %s", goldenFile, err)
		}

		return
	}

	want, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("Error reading golden file: %s", err)
	}

	assert.Equal(t, want, got)
}

type Output struct {
	buf bytes.Buffer
}

func (o *Output) Print(a ...interface{}) {
	fmt.Fprintln(&o.buf, a...)
}

func (o *Output) Error(a ...interface{}) {
	fmt.Fprintln(&o.buf, a...)
}

func (o *Output) Fatal(code int, a ...interface{}) {
	o.Error(a...)
	panic(code)
}
