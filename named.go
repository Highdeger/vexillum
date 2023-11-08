package vexillum

import (
	"fmt"
	"reflect"
	"strings"
)

// named represents a flag which has short and/or long names.
// this flag can be set with referring by one of its short or long names,
// e.g. "app-exe -f value --flag value".
type named struct {
	core
	short rune
	long  string
}

// static private methods

// newNamedFlag returns a new named flag.
func newNamedFlag[T string | int | float64 | bool](short rune, long, help string, defaultValue T, validator func(T) error) (*named, *T) {
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

	return &named{
		core: core{
			help:      strings.Trim(help, "\n\t\r "),
			pointer:   &v,
			def:       defaultValue,
			validator: validator2,
			kind:      kind,
			referred:  false,
		},
		short: short,
		long:  long,
	}, &v
}

// non-static private methods

// name returns the name of the named flag.
// it can be lengthened to a certain max length.
// e.g. "-h     --help".
func (r *named) name(length int) string {

	space := strings.Builder{}

	for i := 0; i < length-len(fmt.Sprintf("-%s--%s", string(r.short), r.long)); i++ {
		space.WriteString(" ")
	}

	if r.long == "" && string(r.short) != "" {
		return space.String() + string(r.short)
	} else if string(r.short) == "" && r.long != "" {
		return space.String() + r.long
	}

	return fmt.Sprintf("-%s%s--%s", string(r.short), space.String(), r.long)
}

// id returns the unique id of the named flag.
// e.g. "-h --help".
func (r *named) id() string {
	if r.long == "" && string(r.short) != "" {
		return string(r.short)
	} else if string(r.short) == "" && r.long != "" {
		return r.long
	}

	return fmt.Sprintf("-%s --%s", string(r.short), r.long)
}
