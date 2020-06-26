// +build !windows,!plan9

package git

func isSymlinkWindowsNonAdmin(err error) bool {
	return false
}
