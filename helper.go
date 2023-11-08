package vexillum

import (
	"strings"
)

// detectFlag returns a flag and its type.
// it trims "-" or "--" from the beginning of the named flags.
func detectFlag(f string) (string, flagType) {
	if strings.HasPrefix(f, "--") {
		return strings.TrimPrefix(f, "--"), Long
	} else if strings.HasPrefix(f, "-") {
		return strings.TrimPrefix(f, "-"), Short
	} else {
		return f, Wild
	}
}

// logErrorNotExist logs an error when a flag does not exist.
func logErrorNotExist(g *App, flag string) {
	g.logError("'%s' does not exist", flag)
}

// logWarningValueMissing logs a warning when a flag value is missing.
func logWarningValueMissing(g *App, flag string) {
	g.logWarning("'%s' set to default because the value is missing", flag)
}

// logWarningValueInvalid logs a warning when a flag value is invalid.
func logWarningValueInvalid(g *App, flag string, err string) {
	g.logWarning("'%s' set to default because the value is invalid (validation error: %s)", flag, err)
}

// logWarningValueNotReferred logs a warning when a flag is not referred.
func logWarningValueNotReferred(g *App, flag string) {
	g.logWarning("'%s' set to default because it's not referred", flag)
}
