package codes

import (
	"math"
	"runtime"

	"github.com/amsokol/go-eccodes/debug"
	"github.com/amsokol/go-eccodes/native"
)

type Message interface {
	isOpen() bool

	GetString(key string) (string, error)

	GetLong(key string) (int64, error)
	SetLong(key string, value int64) error

	GetDouble(key string) (float64, error)
	SetDouble(key string, value float64) error

	Iterator() (Iterator, error)

	Data() (latitudes []float64, longitudes []float64, values []float64, err error)
	DataUnsafe() (latitudes *Float64ArrayUnsafe, longitudes *Float64ArrayUnsafe, values *Float64ArrayUnsafe, err error)

	Close() error
}

type message struct {
	handle native.Ccodes_handle
}

func newMessage(h native.Ccodes_handle) Message {
	m := &message{handle: h}
	runtime.SetFinalizer(m, messageFinalizer)

	// set missing value to NaN
	m.SetDouble(parameterMissingValue, math.NaN())

	return m
}

func (m *message) isOpen() bool {
	return m.handle != nil
}

func (m *message) GetString(key string) (string, error) {
	return native.Ccodes_get_string(m.handle, key)
}

func (m *message) GetLong(key string) (int64, error) {
	return native.Ccodes_get_long(m.handle, key)
}

func (m *message) SetLong(key string, value int64) error {
	return native.Ccodes_set_long(m.handle, key, value)
}

func (m *message) GetDouble(key string) (float64, error) {
	return native.Ccodes_get_double(m.handle, key)
}

func (m *message) SetDouble(key string, value float64) error {
	return native.Ccodes_set_double(m.handle, key, value)
}

type iterator struct {
	iterator native.Ccodes_iterator
}

func (i *iterator) Next() (latitude float64, longitude float64, value float64, err error) {
	return native.Ccodes_grib_iterator_next(i.iterator)
}

func (i *iterator) Close() error {
	return native.Ccodes_grib_iterator_delete(i.iterator)
}

type Iterator interface {
	Next() (latitude float64, longitude float64, value float64, err error)
	Close() error
}

func (m *message) Iterator() (Iterator, error) {
	iteratorVal, err := native.Ccodes_grib_iterator_new(m.handle, 0)
	if err != nil {
		return nil, err
	}
	iteratorStruct := &iterator{
		iterator: iteratorVal,
	}
	return iteratorStruct, nil
}

func (m *message) Data() (latitudes []float64, longitudes []float64, values []float64, err error) {
	return native.Ccodes_grib_get_data(m.handle)
}

func (m *message) DataUnsafe() (latitudes *Float64ArrayUnsafe, longitudes *Float64ArrayUnsafe, values *Float64ArrayUnsafe, err error) {
	lats, lons, vals, err := native.Ccodes_grib_get_data_unsafe(m.handle)
	if err != nil {
		return nil, nil, nil, err
	}
	return newFloat64ArrayUnsafe(lats), newFloat64ArrayUnsafe(lons), newFloat64ArrayUnsafe(vals), nil
}

func (m *message) Close() error {
	defer func() { m.handle = nil }()
	return native.Ccodes_handle_delete(m.handle)
}

func messageFinalizer(m *message) {
	if m.isOpen() {
		debug.MemoryLeakLogger.Print("message is not closed")
		m.Close()
	}
}
