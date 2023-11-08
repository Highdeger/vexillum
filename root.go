package vexillum

import "os"

var (
	// width is the count of characters inside which the printed text is wrapped horizontally.
	// texts will be break into lines if the width is exceeded,
	// but not from in the middle of a word. default is 80.
	width         int
	root          *App
	current       *App
	remainingArgs []string
)

func init() {
	width = 80
	root = newApp("App", "v1.0.0")
	current = root
	remainingArgs = make([]string, 0)
}

// CurrentApp returns the current app.
// best to be used in switch, for example:
//
//	switch root.CurrentApp() {
//		case app1:
//			if "app-exe app1 ..." was run.
//		case app2:
//			if "app-exe app2 ..." was run.
//		default:
//			if "app-exe ..." was run.
//	}
func CurrentApp() *App {
	return current
}

// PrintWidth sets the count of characters inside which the printed text is wrapped horizontally.
// texts will be break into lines if the width is exceeded,
// but not from in the middle of a word.
// default value is 80 characters.
func PrintWidth(w int) {
	width = w
}

// SetApp sets the app name.
func SetApp(name string) {
	root.SetApp(name)
}

// SetVersion sets the app version.
func SetVersion(version string) {
	root.SetVersion(version)
}

// Name returns the app name and version together.
func Name() string {
	return root.Name()
}

// ShowWarnings sets whether to print warnings or not.
func ShowWarnings(show bool) {
	root.ShowWarnings(show)
}

// OnBareRun sets a function to be called when the app is run without any arguments.
func OnBareRun(f func()) {
	root.OnBareRun(f)
}

// OnError sets a function to be called when an error occurs.
func OnError(f func()) {
	root.OnError(f)
}

// OnHelp sets a function to be called when the app is run with -h or --help.
func OnHelp(f func()) {
	root.OnHelp(f)
}

// PrintUsage prints the usage of the app.
func PrintUsage() {
	root.PrintUsage()
}

// StringValidated adds a string named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func StringValidated(short rune, long, help string, defaultValue string, validator func(string) error) *string {
	return root.StringValidated(short, long, help, defaultValue, validator)
}

// IntValidated adds an int named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func IntValidated(short rune, long, help string, defaultValue int, validator func(int) error) *int {
	return root.IntValidated(short, long, help, defaultValue, validator)
}

// Float64Validated adds a float64 named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func Float64Validated(short rune, long, help string, defaultValue float64, validator func(float64) error) *float64 {
	return root.Float64Validated(short, long, help, defaultValue, validator)
}

// String adds a string named flag to the app and returns a pointer to its value.
func String(short rune, long, help string, defaultValue string) *string {
	return root.String(short, long, help, defaultValue)
}

// Int adds an int named flag to the app and returns a pointer to its value.
func Int(short rune, long, help string, defaultValue int) *int {
	return root.Int(short, long, help, defaultValue)
}

// Float64 adds a float64 named flag to the app and returns a pointer to its value.
func Float64(short rune, long, help string, defaultValue float64) *float64 {
	return root.Float64(short, long, help, defaultValue)
}

// Bool adds a bool named flag to the app and returns a pointer to its value.
func Bool(short rune, long, help string, defaultValue bool) *bool {
	return root.Bool(short, long, help, defaultValue)
}

// WildStringValidator adds a string wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func WildStringValidator(placeholder, help string, defaultValue string, validator func(string) error) *string {
	return root.WildStringValidator(placeholder, help, defaultValue, validator)
}

// WildIntValidator adds an int wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func WildIntValidator(placeholder, help string, defaultValue int, validator func(int) error) *int {
	return root.WildIntValidator(placeholder, help, defaultValue, validator)
}

// WildFloat64Validator adds a float64 wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func WildFloat64Validator(placeholder, help string, defaultValue float64, validator func(float64) error) *float64 {
	return root.WildFloat64Validator(placeholder, help, defaultValue, validator)
}

// WildString adds a string wild flag to the app and returns a pointer to its value.
func WildString(placeholder, help string, defaultValue string) *string {
	return root.WildString(placeholder, help, defaultValue)
}

// WildInt adds an int wild flag to the app and returns a pointer to its value.
func WildInt(placeholder, help string, defaultValue int) *int {
	return root.WildInt(placeholder, help, defaultValue)
}

// WildFloat64 adds a float64 wild flag to the app and returns a pointer to its value.
func WildFloat64(placeholder, help string, defaultValue float64) *float64 {
	return root.WildFloat64(placeholder, help, defaultValue)
}

// Remaining returns the remaining arguments which are not defined as flags,
// and left out at the end of parsing.
func Remaining() []string {
	return remainingArgs
}

// Parse parses the arguments, and set all the values.
func Parse() {
	root.Parse(os.Args...)
}

// NoHelpFlag disables the help flag.
func NoHelpFlag() {
	root.NoHelpFlag()
}

// NewApp adds a new app to the app.
func NewApp(app, version string) *App {
	return root.NewApp(app, version)
}
