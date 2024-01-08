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
	Mode    string     `yaml:"mode"`
	Threads int        `yaml:"threads"`
	Log     LogOptions `yaml:"log"`
}

type RBACOptions struct {
	Model string `yaml:"model"`
}

type APIServerOptions struct {
	Listen string      `yaml:"listen"`
	Path   string      `yaml:"path"`
	RBAC   RBACOptions `yaml:"rbac"`
}

type Options struct {
	Core      CoreOptions      `yaml:"core"`
	APIServer APIServerOptions `yaml:"apiserver"`
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
