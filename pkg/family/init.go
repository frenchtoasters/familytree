package family

import (
	"github.com/robfig/cron/v3"
	flag "github.com/spf13/pflag"
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

func (c *FamilyConfig) AddFlags(fs *flag.FlagSet) {
	fs.StringSliceVar(&c.Parents, "parents", defaultParents, "Parents of Family")
	fs.StringSliceVar(&c.Children, "children", defultChildren, "Children of Family")
}
