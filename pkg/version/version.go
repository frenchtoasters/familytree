package version

var (
	// These variables typically come from -ldflags settings in build

	// Version shows the etcd-backup-restore binary version.
	Version string
	// GitSHA shows the etcd-backup-restore binary code commit SHA on git.
	GitSHA string
)
