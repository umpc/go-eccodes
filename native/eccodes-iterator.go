package native

/*
#include <eccodes.h>
*/
import "C"
import (
	"unsafe"

	"github.com/amsokol/go-errors"
)

func Ccodes_grib_iterator_new(handle Ccodes_handle, flags int, err error) (Ccodes_grib_iterator, error) {
  var errInt C.int
  defer C.free(unsafe.Pointer(&errInt))

  ccodes_grib_iterator := unsafe.Pointer(C.codes_grib_iterator_new((*C.codes_handle)(handle), C.ulong(Culong(flags)), &errInt))
  if int(errInt) == 0 {
    return ccodes_grib_iterator, nil
  }
  return ccodes_grib_iterator, errors.New("int(errInt) != 0")
}

func Ccodes_grib_iterator_next(kiter Ccodes_grib_iterator) int {
	return int(C.codes_grib_iterator_next((*C.codes_grib_iterator)(kiter)))
}

func Ccodes_grib_iterator_get_name(kiter Ccodes_grib_iterator) string {
	return C.GoString(C.codes_grib_iterator_get_name((*C.codes_grib_iterator)(kiter)))
}

func Ccodes_grib_iterator_delete(kiter Ccodes_grib_iterator) error {
	err := C.codes_grib_iterator_delete((*C.codes_grib_iterator)(kiter))
	if err != 0 {
		return errors.New(Cgrib_get_error_message(int(err)))
	}
	return nil
}