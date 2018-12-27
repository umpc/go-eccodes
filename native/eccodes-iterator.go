package native

/*
#include <eccodes.h>
*/
import "C"
import (
	"unsafe"

	"github.com/amsokol/go-errors"
)

func Ccodes_grib_iterator_new(handle Ccodes_handle, flags int) (Ccodes_iterator, error) {
  res := C.int(0)
  ccodes_iterator := unsafe.Pointer(C.codes_grib_iterator_new((*C.codes_handle)(handle), C.ulong(Culong(flags)), (*C.int)(&res)))
  if int(res) != 0 {
    return ccodes_iterator, errors.New(Cgrib_get_error_message(int(res)))
  }
  return ccodes_iterator, nil
}

func Ccodes_grib_iterator_next(kiter Ccodes_iterator) (latitude float64, longitude float64, value float64, err error) {
	res := int(C.codes_grib_iterator_next(
		(*C.codes_iterator)(kiter),
		(*C.double)(unsafe.Pointer(&latitude)),
		(*C.double)(unsafe.Pointer(&longitude)),
		(*C.double)(unsafe.Pointer(&value)),
	))
	if res != 0 {
		return 0, 0, 0, errors.New(Cgrib_get_error_message(int(res)))
	}
	return latitude, longitude, value, nil
}

func Ccodes_grib_iterator_delete(kiter Ccodes_iterator) error {
	res := C.codes_grib_iterator_delete((*C.codes_iterator)(kiter))
	if res != 0 {
		return errors.New(Cgrib_get_error_message(int(res)))
	}
	return nil
}
