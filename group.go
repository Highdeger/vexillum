package vexillum

import (
	"fmt"
	"os"
	"strings"
)

// App represents a group of flags specifics to a single app.
type App struct {
	app            string
	version        string
	namedList      *namedList
	wildList       *wildList
	groupList      []*App
	parseIndex     int
	parseIndexWild int
	showWarnings   bool
	onBareRun      func()
	onError        func()
	onHelp         func()
	textUsage      string
	textNamedFlags string
	textWildFlags  string
	parentApp      *App
}

// static private methods

// newApp returns a new App.
func newApp(app, version string) *App {
	g := &App{
		app:            app,
		version:        version,
		namedList:      newNamedList(),
		wildList:       newWildList(),
		parseIndex:     0,
		parseIndexWild: 0,
		showWarnings:   false,
		onBareRun:      func() {},
		onError:        func() {},
		onHelp:         func() {},
		textUsage:      "usage:",
		textNamedFlags: "named flags:",
		textWildFlags:  "wild flags:",
		parentApp:      nil,
	}

	g.onBareRun = func() {
		fmt.Println("app ran without any arguments")
		fmt.Println()
		g.PrintUsage()
		fmt.Println()
		os.Exit(0)
	}
	g.onError = func() {
		fmt.Println("incorrect app usage")
		fmt.Println()
		g.PrintUsage()
		fmt.Println()
		os.Exit(1)
	}
	g.onHelp = func() {
		g.PrintUsage()
		fmt.Println()
		os.Exit(0)
	}

	g.Bool('h', "help", "show the help", false)

	return g
}

// addNamedFlag adds a named flag to an app and returns a pointer to its value.
func addNamedFlag[T string | int | float64 | bool](g *App, short rune, long, help string, defaultValue T, validator func(T) error) *T {
	if foundFlag := g.namedList.findByShort(short); foundFlag != nil {
		panic(fmt.Sprintf("flag '-%s' already exists", string(short)))
	}

	if foundFlag := g.namedList.findByLong(long); foundFlag != nil {
		panic(fmt.Sprintf("flag '--%s' already exists", long))
	}

	f, v := newNamedFlag(short, long, help, defaultValue, validator)
	g.namedList.add(f)

	return v
}

// addWildFlag adds a wild flag to an app and returns a pointer to its value.
func addWildFlag[T string | int | float64](g *App, placeholder, help string, defaultValue T, validator func(T) error) *T {
	if foundFlag := g.wildList.findByPlaceholder(placeholder); foundFlag != nil {
		panic(fmt.Sprintf("flag with placeholder '%s' already exists", placeholder))
	}

	f, v := newWildFlag(g.wildList.len(), placeholder, help, defaultValue, validator)
	g.wildList.add(f)

	return v
}

// static public methods

// SetApp sets the app name.
func (r *App) SetApp(name string) {
	r.app = name
}

// SetVersion sets the app version.
func (r *App) SetVersion(version string) {
	r.version = version
}

// Name returns the app name and version together.
func (r *App) Name() string {
	return fmt.Sprintf("%s %s", r.app, r.version)
}

// ShowWarnings sets whether to print warnings or not.
func (r *App) ShowWarnings(show bool) {
	r.showWarnings = show
}

// OnBareRun sets a function to be called when the app is run without any arguments.
func (r *App) OnBareRun(f func()) {
	r.onBareRun = f
}

// OnError sets a function to be called when an error occurs.
func (r *App) OnError(f func()) {
	r.onError = f
}

// OnHelp sets a function to be called when the app is run with -h or --help.
func (r *App) OnHelp(f func()) {
	r.onHelp = f
}

// PrintUsage prints the usage of the app.
func (r *App) PrintUsage() {
	b := strings.Builder{}

	name := r.Name()
	app := r

	for app.parentApp != nil {
		name = fmt.Sprintf("%s %s", app.parentApp.app, name)
		app = app.parentApp
	}

	b.WriteString(name)
	b.WriteString("\n")

	if r.namedList.len() != 0 || r.wildList.len() != 0 {
		b.WriteString(r.textUsage)

		cmdBuilder := strings.Builder{}
		cmdBuilder.WriteString("  cli> ")

		filepath := strings.Split(os.Args[0], string(os.PathSeparator))
		cmdBuilder.WriteString(filepath[len(filepath)-1])

		if r.namedList.len() != 0 {
			cmdBuilder.WriteString(" [named flags]")
		}

		for _, f := range r.wildList.list() {
			cmdBuilder.WriteString(fmt.Sprintf(" [%s]", f.placeholder))
		}

		b.WriteString("\n")
		b.WriteString(cmdBuilder.String())

		if r.namedList.len() != 0 {
			b.WriteString("\n")
			b.WriteString(fmt.Sprintf("  %s", r.textNamedFlags))

			for _, f := range r.namedList.list() {
				b.WriteString("\n")

				defaultValue := f.def
				if f.kind == typeString {
					defaultValue = fmt.Sprintf("\"%s\"", defaultValue)
				}

				b.WriteString(fmt.Sprintf("    %s: (type: %s, default: %v)", f.name(r.namedList.maxIdLength()), f.kind, defaultValue))

				if f.help != "" {
					b.WriteString("\n")
					b.WriteString(f.helpBlock("      ", width))
				}
			}
		}

		if r.wildList.len() != 0 {
			b.WriteString("\n")
			b.WriteString(fmt.Sprintf("  %s", r.textWildFlags))

			for _, f := range r.wildList.list() {
				b.WriteString("\n")

				defaultValue := f.def
				if f.kind == typeString {
					defaultValue = fmt.Sprintf("\"%s\"", defaultValue)
				}

				b.WriteString(fmt.Sprintf("    %s: (type: %s, default: %v)", f.name(r.wildList.maxIdLength()), f.kind, defaultValue))

				if f.help != "" {
					b.WriteString("\n")
					b.WriteString(f.helpBlock("      ", width))
				}
			}
		}
	}

	fmt.Println(b.String())
}

// StringValidated adds a string named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) StringValidated(short rune, long, help string, defaultValue string, validator func(string) error) *string {
	return addNamedFlag(r, short, long, help, defaultValue, validator)
}

// IntValidated adds an int named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) IntValidated(short rune, long, help string, defaultValue int, validator func(int) error) *int {
	return addNamedFlag(r, short, long, help, defaultValue, validator)
}

// Float64Validated adds a float64 named flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) Float64Validated(short rune, long, help string, defaultValue float64, validator func(float64) error) *float64 {
	return addNamedFlag(r, short, long, help, defaultValue, validator)
}

// String adds a string named flag to the app and returns a pointer to its value.
func (r *App) String(short rune, long, help string, defaultValue string) *string {
	return addNamedFlag(r, short, long, help, defaultValue, nil)
}

// Int adds an int named flag to the app and returns a pointer to its value.
func (r *App) Int(short rune, long, help string, defaultValue int) *int {
	return addNamedFlag(r, short, long, help, defaultValue, nil)
}

// Float64 adds a float64 named flag to the app and returns a pointer to its value.
func (r *App) Float64(short rune, long, help string, defaultValue float64) *float64 {
	return addNamedFlag(r, short, long, help, defaultValue, nil)
}

// Bool adds a bool named flag to the app and returns a pointer to its value.
func (r *App) Bool(short rune, long, help string, defaultValue bool) *bool {
	return addNamedFlag(r, short, long, help, defaultValue, nil)
}

// WildStringValidator adds a string wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) WildStringValidator(placeholder, help string, defaultValue string, validator func(string) error) *string {
	return addWildFlag(r, placeholder, help, defaultValue, validator)
}

// WildIntValidator adds an int wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) WildIntValidator(placeholder, help string, defaultValue int, validator func(int) error) *int {
	return addWildFlag(r, placeholder, help, defaultValue, validator)
}

// WildFloat64Validator adds a float64 wild flag to the app and returns a pointer to its value.
// it gets a validator function to validate the value before setting it.
func (r *App) WildFloat64Validator(placeholder, help string, defaultValue float64, validator func(float64) error) *float64 {
	return addWildFlag(r, placeholder, help, defaultValue, validator)
}

// WildString adds a string wild flag to the app and returns a pointer to its value.
func (r *App) WildString(placeholder, help string, defaultValue string) *string {
	return addWildFlag(r, placeholder, help, defaultValue, nil)
}

// WildInt adds an int wild flag to the app and returns a pointer to its value.
func (r *App) WildInt(placeholder, help string, defaultValue int) *int {
	return addWildFlag(r, placeholder, help, defaultValue, nil)
}

// WildFloat64 adds a float64 wild flag to the app and returns a pointer to its value.
func (r *App) WildFloat64(placeholder, help string, defaultValue float64) *float64 {
	return addWildFlag(r, placeholder, help, defaultValue, nil)
}

// Parse parses the arguments, and set all the values.
func (r *App) Parse(args ...string) {
	if len(args) == 1 {
		r.onBareRun()
	} else if len(args) > 1 {
		r.parseIndex = 0
		r.parseIndexWild = 0
		args2 := args[1:]

		for _, g := range r.groupList {
			if g.app == args2[0] {
				current = g
				g.Parse(args2...)
				return
			}
		}

		for r.parseIndex < len(args2) {
			f, fType := detectFlag(args2[r.parseIndex])
			switch fType {
			case Short:
				r.parseShort(f, &args2)
			case Long:
				r.parseLong(f, &args2)
			case Wild:
				r.parseWild(f)
			}

			r.parseIndex++
		}

		for i, f := range r.namedList.list() {
			if !f.referred && i != r.helpIndex() {
				logWarningValueNotReferred(r, f.id())
			}
		}
		for _, f := range r.wildList.list() {
			if !f.referred {
				logWarningValueNotReferred(r, f.id())
			}
		}

		if r.helpIndex() > -1 && r.helpTriggered() {
			r.onHelp()
		}
	}
}

// NoHelpFlag disables the help flag.
func (r *App) NoHelpFlag() {
	i := r.helpIndex()
	if i != -1 {
		r.namedList.remove(i)
	}
}

// NewApp adds a new app to the app.
func (r *App) NewApp(app, version string) *App {
	for _, g := range r.groupList {
		if g.app == app {
			r.logError("app '%s' already exists in the group '%s'", app, g.Name())
			return nil
		}
	}

	g := newApp(app, version)
	g.parentApp = r
	r.groupList = append(r.groupList, g)

	return g
}

// non-static private methods

// logWarning logs a warning if App.showWarnings is true.
func (r *App) logWarning(format string, a ...any) {
	if r.showWarnings {
		loggerWarning.Printf(format, a...)
	}
}

// logError logs an error. runs App.onError().
func (r *App) logError(format string, a ...any) {
	loggerError.Printf(format, a...)
	r.onError()
}

// parseShort parses a short flag, e.g. "-h".
func (r *App) parseShort(f string, args *[]string) {
	shorts := make([]rune, 0)
	for _, r := range f {
		shorts = append(shorts, r)
	}

	if len(shorts) == 1 { // single flag
		flag := r.namedList.findByShort(shorts[0])
		if flag != nil {
			var (
				valueSet             bool
				valueValidationError error
				setBool              = true
			)

			if r.parseIndex != len(*args)-1 { // non-last flag
				nextFlag, nextFlagType := detectFlag((*args)[r.parseIndex+1])
				if nextFlagType == Wild {
					valueSet = true
					r.parseIndex++
					valueValidationError = flagParse(&flag.core, nextFlag)
					setBool = false
				}
			}

			if setBool { // last flag OR non-wild next flag
				if flag.kind == typeBool {
					valueSet = true
					flagSetValue(&flag.core, true)
				}
			}

			flag.core.referred = true

			if !valueSet {
				logWarningValueMissing(r, flag.id())
			} else if valueValidationError != nil {
				logWarningValueInvalid(r, flag.id(), valueValidationError.Error())
			}
		} else {
			logErrorNotExist(r, "-"+string(shorts[0]))
		}
	} else { // flag group
		for i, sh := range shorts {
			var (
				valueSet bool
				setBool  = true
			)

			flag := r.namedList.findByShort(sh)
			if flag != nil {
				if i != len(shorts)-1 { // non-last short flag in a group
					if flag.kind != typeBool {
						r.logError("flag '-%s' should be boolean because it's inside a group of flags and it's not the last flag", string(sh))
					}
				} else { // last short flag in a group
					if r.parseIndex != len(*args)-1 { // non-last flag
						nextFlag, nextFlagType := detectFlag((*args)[r.parseIndex+1])
						if nextFlagType == Wild {
							valueSet = true
							setBool = false
							r.parseIndex++

							err := flagParse(&flag.core, nextFlag)
							if err != nil {
								r.logError(err.Error())
							}
						}
					}
				}

				if setBool { // last flag OR non-wild next flag
					valueSet = true
					flagSetValue(&flag.core, true)
				}

				flag.core.referred = true

				if !valueSet {
					logWarningValueMissing(r, flag.id())
				}
			} else {
				logErrorNotExist(r, "-"+string(sh))
			}
		}
	}
}

// parseLong parses a long flag, e.g. "--help".
func (r *App) parseLong(f string, args *[]string) {
	flag := r.namedList.findByLong(f)
	if flag != nil {
		var (
			valueSet             bool
			valueValidationError error
			setBool              = true
		)

		if r.parseIndex != len(*args)-1 { // non-last flag
			nextFlag, nextFlagType := detectFlag((*args)[r.parseIndex+1])
			if nextFlagType == Wild {
				valueSet = true
				r.parseIndex++
				valueValidationError = flagParse(&flag.core, nextFlag)
				setBool = false
			}
		}

		if setBool { // last flag OR non-wild next flag
			if flag.kind == typeBool {
				valueSet = true
				flagSetValue(&flag.core, true)
			}
		}

		flag.core.referred = true

		if !valueSet {
			logWarningValueMissing(r, flag.id())
		} else if valueValidationError != nil {
			logWarningValueInvalid(r, flag.id(), valueValidationError.Error())
		}
	} else {
		logErrorNotExist(r, "--"+f)
	}
}

// parseWild parses a wild flag.
func (r *App) parseWild(f string) {
	var valueValidationError error

	flag := r.wildList.findByIndex(r.parseIndexWild)
	if flag != nil {
		flag.core.referred = true
		valueValidationError = flagParse(&flag.core, f)
	} else {
		remainingArgs = append(remainingArgs, f)
	}

	if valueValidationError != nil {
		logWarningValueInvalid(r, flag.id(), valueValidationError.Error())
	}

	r.parseIndexWild++
}

// helpTriggered returns true if the help flag is triggered.
func (r *App) helpTriggered() bool {
	f := r.namedList.findByShortAndLong('h', "help")
	if f != nil {
		return flagGetValue[bool](&f.core)
	}

	return false
}

// helpIndex returns the index of the help flag.
func (r *App) helpIndex() int {
	for i, v := range r.namedList.list() {
		if v.id() == "-h --help" {
			return i
		}
	}

	return -1
}
