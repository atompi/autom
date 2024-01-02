package options

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Version string = "v0.0.1"

type LogOptions struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type CoreOptions struct {
	Threads int        `yaml:"threads"`
	Log     LogOptions `yaml:"log"`
}

type Options struct {
	Core CoreOptions `yaml:"core"`
}

func NewOptions() (opts Options) {
	optsSource := viper.AllSettings()
	err := createOptions(optsSource, &opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create options failed:", err)
		os.Exit(1)
	}
	return
}
