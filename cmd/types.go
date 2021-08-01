package cmd

import "github.com/sirupsen/logrus"

var (
	version     bool
	logger      = logrus.New()
	emptyStruct struct{}
)
