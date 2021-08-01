package cmd

import (
	"context"
	"io/ioutil"

	"github.com/frenchtoasters/familytree/pkg/family"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

type treeOptions struct {
	ConfigFile string
	Version    bool
	LogLevel   uint32
	Logger     *logrus.Logger
	Config     *family.TreeComponentConfig
}

// newTreeOptions returns a new Options object.
func newTreeOptions() *treeOptions {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &treeOptions{
		LogLevel: 4,
		Version:  false,
		Config:   family.NewTreeComponentConfig(),
		Logger:   logger,
	}
}

func (o *treeOptions) validate() error {
	return o.Config.Validate()
}

func (o *treeOptions) addFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.ConfigFile, "config-file", o.ConfigFile, "path to the configuration file")
	fs.Uint32Var(&o.LogLevel, "log-level", o.LogLevel, "verbosity level of logs")
	o.Config.AddFlags(fs)
}

func (o *treeOptions) complete() {
	o.Config.Complete()
	o.Logger.SetLevel(logrus.Level(o.LogLevel))
}

func (o *treeOptions) loadFamilyFromFile() error {
	if len(o.ConfigFile) != 0 {
		data, err := ioutil.ReadFile(o.ConfigFile)
		if err != nil {
			return err
		}
		config := family.NewTreeComponentConfig()
		if err := yaml.Unmarshal(data, config); err != nil {
			return err
		}
		o.Config = config
	}
	return nil
}

func (o *treeOptions) run(ctx context.Context) error {
	familyTree, err := family.NewFamilyTree(o.Logger, o.Config)
	if err != nil {
		return err
	}
	return familyTree.Run(ctx)
}
