package native

/*
#include <eccodes.h>
*/
import "C"
import (
	"unsafe"

	"github.com/amsokol/go-errors"
)

func Ccodes_grib_iterator_new(handle Ccodes_handle, flags int) (Ccodes_grib_iterator, error) {
  var err C.int

  ccodes_grib_iterator := unsafe.Pointer(C.codes_grib_iterator_new((*C.codes_handle)(handle), C.ulong(Culong(flags)), (*C.int)(&err)))
  if int(err) != 0 {
    return ccodes_grib_iterator, errors.New(Cgrib_get_error_message(int(err)))
  }

  return ccodes_grib_iterator, nil
}

func Ccodes_grib_iterator_next(kiter Ccodes_grib_iterator) (latitude float64, longitude float64, value float64, err error) {
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

func Ccodes_grib_iterator_delete(kiter Ccodes_grib_iterator) error {
	err := C.codes_grib_iterator_delete((*C.codes_iterator)(kiter))
	if err != 0 {
		return errors.New(Cgrib_get_error_message(int(err)))
	}
	return nil
}
