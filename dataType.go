package vexillum

// dataType represents the data type of the value of flag.
type dataType string

const (
	typeString  dataType = "string"  // typeString represents a string flag.
	typeInt     dataType = "integer" // typeInt represents an integer flag.
	typeFloat64 dataType = "decimal" // typeFloat64 represents a decimal flag.
	typeBool    dataType = "boolean" // typeBool represents a boolean flag.s
)
