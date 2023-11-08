package vexillum

// flagType represents the type of flag.
type flagType int

const (
	Short flagType = iota // Short represents a short named flag, e.g. "-f".
	Long                  // Long represents a long named flag, e.g. "--flag".
	Wild                  // Wild represents an ordered flag without a name.
)
