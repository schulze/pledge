package pledge

//#include <unistd.h>
import "C"

func Pledge(promises string, paths []string) error {
	// TODO: Should also pass paths, but path whitelisting is not implemented yet, so no hurry.
	_, err := C.pledge(C.CString(promises), nil)
	return err
}


