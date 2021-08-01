package cmd

import (
	"runtime"

	ver "github.com/frenchtoasters/familytree/pkg/version"
)

func printVersionInfo() {
	logger.Infof("etcd-backup-restore Version: %s", ver.Version)
	logger.Infof("Git SHA: %s", ver.GitSHA)
	logger.Infof("Go Version: %s", runtime.Version())
	logger.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
}
