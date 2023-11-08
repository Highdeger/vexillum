package vexillum

// namedList represents an array of named flags.
type namedList []*named

// static private methods

// newNamedList returns a new namedList.
func newNamedList() *namedList {
	return &namedList{}
}

// non-static private methods

// add appends a named flag to the namedList.
func (r *namedList) add(flag *named) {
	*r = append(*r, flag)
}

// remove deletes a named flag from the namedList.
func (r *namedList) remove(index int) {
	*r = append((*r)[:index], (*r)[index+1:]...)
}

// len returns the length of the namedList.
func (r *namedList) len() int {
	return len(*r)
}

// list returns namedList as an array.
func (r *namedList) list() []*named {
	return *r
}

// findByShort finds and returns a named flag by its short name.
// returns nil if not found.
func (r *namedList) findByShort(short rune) *named {
	for _, v := range *r {
		if v.short == short {
			return v
		}
	}

	return nil
}

// findByLong finds and returns a named flag by its long name.
// returns nil if not found.
func (r *namedList) findByLong(long string) *named {
	for _, v := range *r {
		if v.long == long {
			return v
		}
	}

	return nil
}

// findByShortAndLong finds and returns a named flag by both its short and long names.
// returns nil if not found.
func (r *namedList) findByShortAndLong(short rune, long string) *named {
	for _, v := range *r {
		if v.short == short && v.long == long {
			return v
		}
	}

	return nil
}

// maxIdLength returns the length of longest flag id in the namedList.
func (r *namedList) maxIdLength() int {
	m := 0

	for _, v := range *r {
		if len(v.id()) > m {
			m = len(v.id())
		}
	}

	return m
}
