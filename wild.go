package vexillum

import (
	"fmt"
	"reflect"
	"strings"
)

// wild represents a flag which has index and placeholder.
// this flag can be set with referring it withing the right order of index,
// without any name, e.g. "app-exe hello 2 3.14".
type wild struct {
	core
	index       int
	placeholder string
}

// static private methods

// newWildFlag returns a new wild flag.
func newWildFlag[T string | int | float64](index int, placeholder string, help string, defaultValue T, validator func(T) error) (*wild, *T) {
	v := defaultValue

	kind := typeString
	switch reflect.ValueOf(defaultValue).Kind() {
	case reflect.Int:
		kind = typeInt
	case reflect.Float64:
		kind = typeFloat64
	case reflect.Bool:
		kind = typeBool
	}

	validator2 := func(T) error {
		return nil
	}
	if validator != nil {
		validator2 = validator
	}

	return &wild{
		core: core{
			help:      help,
			pointer:   &v,
			def:       defaultValue,
			validator: validator2,
			kind:      kind,
			referred:  false,
		},
		index:       index,
		placeholder: placeholder,
	}, &v
}

// non-static private methods

// name returns the name of the wild flag.
// it can be lengthened to a certain max length.
// e.g. "[0]     input-text"
func (r *wild) name(length int) string {
	space := strings.Builder{}

	for i := 0; i < length-len(fmt.Sprintf("[%d]%s", r.index, r.placeholder)); i++ {
		space.WriteString(" ")
	}

	return fmt.Sprintf("[%d]%s%s", r.index, space.String(), r.placeholder)
}

// id returns the unique id of the named flag.
// e.g. "[0] input-text"
func (r *wild) id() string {
	return fmt.Sprintf("[%d] %s", r.index, r.placeholder)
}
