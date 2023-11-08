package vexillum

// wildList represents an array of wild flags.
type wildList []*wild

// static private methods

// newWildList returns a new wildList.
func newWildList() *wildList {
	return &wildList{}
}

// non-static private methods

// add appends a wild flag to the wildList.
func (r *wildList) add(flag *wild) {
	*r = append(*r, flag)
}

// remove deletes a wild flag from the wildList.
func (r *wildList) remove(index int) {
	*r = append((*r)[:index], (*r)[index+1:]...)
}

// len returns the length of the wildList.
func (r *wildList) len() int {
	return len(*r)
}

// list returns wildList as an array.
func (r *wildList) list() []*wild {
	return *r
}

// findByIndex finds and returns a wild flag by its index.
// returns nil if not found.
func (r *wildList) findByIndex(index int) *wild {
	for _, v := range *r {
		if v.index == index {
			return v
		}
	}

	return nil
}

// findByPlaceholder finds and returns a wild flag by its placeholder.
// returns nil if not found.
func (r *wildList) findByPlaceholder(placeholder string) *wild {
	for _, v := range *r {
		if v.placeholder == placeholder {
			return v
		}
	}

	return nil
}

// findByIndexAndPlaceholder finds and returns a wild flag by its index and placeholder.
// returns nil if not found.
func (r *wildList) findByIndexAndPlaceholder(index int, placeholder string) *wild {
	for _, v := range *r {
		if v.index == index && v.placeholder == placeholder {
			return v
		}
	}

	return nil
}

// maxIdLength returns the length of the longest id of a wildList.
func (r *wildList) maxIdLength() int {
	m := 0

	for _, v := range *r {
		if len(v.id()) > m {
			m = len(v.id())
		}
	}

	return m
}
