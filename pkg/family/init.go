package family

import (
	"flag"

	"github.com/robfig/cron/v3"
)

var (
	defaultFamilyName = "Family"
	defaultParents    = make([]string, 0)
	defultChildren    = make([]string, 0)
)

func NewTreeComponentConfig() *TreeComponentConfig {
	return &TreeComponentConfig{
		FamilyName: defaultFamilyName,
	}
}

func NewFamilyConfig() *FamilyConfig {
	return &FamilyConfig{
		Parents:  defaultParents,
		Children: defultChildren,
	}
}

func (c *TreeComponentConfig) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.FamilyName, "family-name", c.FamilyName, "family name to use for tree")
}

func (c *TreeComponentConfig) Validate() error {
	if _, err := cron.ParseStandard(c.FamilyName); err != nil {
		return err
	}
	return nil
}
