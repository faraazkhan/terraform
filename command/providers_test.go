package command

import (
	"os"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestProviders(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if err := os.Chdir(testFixturePath("providers")); err != nil {
		t.Fatalf("err: %s", err)
	}
	defer os.Chdir(cwd)

	ui := new(cli.MockUi)
	c := &ProvidersCommand{
		Meta: Meta{
			Ui: ui,
		},
	}

	args := []string{}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	output := ui.OutputWriter.String()
	if !strings.Contains(output, "provider.foo") {
		t.Errorf("output missing provider.foo\n\n%s", output)
	}
	if !strings.Contains(output, "provider.bar") {
		t.Errorf("output missing provider.bar\n\n%s", output)
	}
	if !strings.Contains(output, "provider.baz") {
		t.Errorf("output missing provider.baz\n\n%s", output)
	}
}
