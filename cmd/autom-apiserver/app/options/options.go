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

type EtcdTlsOptions struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
	Ca   string `yaml:"ca"`
}

type EtcdOptions struct {
	Endpoints   []string       `yaml:"endpoints"`
	DialTimeout int            `yaml:"dial_timeout"`
	Tls         EtcdTlsOptions `yaml:"tls"`
}

type APIServerOptions struct {
	Listen string      `yaml:"listen"`
	Path   string      `yaml:"path"`
	RBAC   RBACOptions `yaml:"rbac"`
	Etcd   EtcdOptions `yaml:"etcd"`
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
