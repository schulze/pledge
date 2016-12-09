// +build !openbsd

package pledge

func init() {
	panic("pledge() is currently only supported on OpenBSD.")
}
