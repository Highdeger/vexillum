package vexillum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// core is the base for all flags.
type core struct {
	help      string
	kind      dataType
	pointer   any
	def       any
	validator any
	referred  bool
}

// static private methods

// flagGetPointer returns a pointer to the value of a flag.
func flagGetPointer[T string | int | float64 | bool](flag *core) *T {
	return flag.pointer.(*T)
}

// flagGetValue returns the value of a flag.
func flagGetValue[T string | int | float64 | bool](flag *core) T {
	return *flagGetPointer[T](flag)
}

// flagSetValue sets the value of a flag.
func flagSetValue[T string | int | float64 | bool](flag *core, v T) {
	*flagGetPointer[T](flag) = v
}

// flagValidate validates a flag.
// it returns an error if it is invalid.
func flagValidate[T string | int | float64](flag *core, v T) error {
	_ = flagGetPointer[T](flag)

	if flag.validator == nil {
		return nil
	}

	return flag.validator.(func(T) error)(v)
}

// flagParse validates and sets a flag value based on its type.
// it returns an error if it is invalid or if the value is missing.
func flagParse(flag *core, v string) error {
	switch flag.kind {
	case typeString:
		err := flagValidate(flag, v)
		if err != nil {
			return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not valid: %s\n", flag.id(), err.Error()))
		} else {
			flagSetValue(flag, v)
		}
	case typeInt:
		n, e := strconv.ParseInt(v, 10, 0)
		if e != nil {
			return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not an integer number\n", flag.id()))
		} else {
			err := flagValidate(flag, int(n))
			if err != nil {
				return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not valid: %s\n", flag.id(), err.Error()))
			} else {
				flagSetValue(flag, int(n))
			}
		}
	case typeFloat64:
		n, e := strconv.ParseFloat(v, 64)
		if e != nil {
			return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not a decimal number\n", flag.id()))
		} else {
			err := flagValidate(flag, n)
			if err != nil {
				return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not valid: %s\n", flag.id(), err.Error()))
			} else {
				flagSetValue(flag, n)
			}
		}
	case typeBool:
		b, e := strconv.ParseBool(v)
		if e != nil {
			return errors.New(fmt.Sprintf("flag '%s' stayed default because the provided value is not a boolean\n", flag.id()))
		} else {
			flagSetValue(flag, b)
		}
	}

	return nil
}

// non-static private methods

// name returns the name of a flag.
// it can be lengthened to a certain max length.
func (r *core) name(length int) string {
	return ""
}

// id returns the unique id of a flag.
func (r *core) id() string {
	return ""
}

// helpBlock returns the help of a flag in a block of text with a certain indentation and width.
func (r *core) helpBlock(indent string, width int) string {
	breakLine := func(line string, length int) (string, string) {
		line = strings.Trim(line, "\n\r\t ")

		for i := length; i < len(line); i++ {
			if line[i] == ' ' {
				return strings.Trim(line[:i], "\n\r\t "), strings.Trim(line[i:], "\n\r\t ")
			}
		}

		return line, ""
	}
	lineToArray := func(line string, length int) []string {
		l := line
		arr := make([]string, 0)

		for {
			part1, part2 := breakLine(l, length)
			arr = append(arr, part1)

			if part2 == "" {
				break
			}

			l = part2
		}

		return arr
	}
	textToArray := func(text string, length int) []string {
		lines := strings.Split(text, "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] = strings.Trim(lines[i], "\n\r\t ")
		}

		lines2 := make([]string, 0)
		for _, line := range lines {
			lines2 = append(lines2, lineToArray(line, length)...)
		}

		return lines2
	}

	s := strings.Builder{}
	for i, line := range textToArray(r.help, width+len(indent)) {
		if i != 0 {
			s.WriteString("\n")
		}

		s.WriteString(indent + line)
	}

	return s.String()
}
