# vexillum

[![Release](https://img.shields.io/badge/release-v1.0.2-blue)](https://github.com/highdeger/vexillum/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/highdeger/vexillum.svg)](https://pkg.go.dev/github.com/highdeger/vexillum)
[![Go Report Card](https://goreportcard.com/badge/github.com/highdeger/vexillum)](https://goreportcard.com/report/github.com/highdeger/vexillum)

comprehensive solution for managing cli arguments in pure go.

# Example
```go
package main

import (
	"fmt"
	"github.com/highdeger/vexillum"
)

var (
	encryptionTypeHelp = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus et libero in nisl maximus hendrerit.
Morbi vel dignissim neque.
Cras ut nunc eget ante vulputate porttitor at scelerisque nisl. Suspendisse vestibulum mollis tortor. Suspendisse eu augue vestibulum, imperdiet metus accumsan, euismod enim. Quisque auctor dignissim ornare. Nulla ullamcorper, erat id sodales cursus, risus lorem congue orci, sed ultrices arcu orci et orci. In aliquet dapibus commodo. Sed a ligula nibh. Sed at velit vel odio maximus commodo ut a arcu. Aliquam sit amet sem est. Integer id mattis justo. Fusce nec porta erat, eget lobortis dolor. Integer non velit id ipsum aliquam luctus at ac diam. Donec maximus venenatis auctor.
`
	subTypeHelp    = "sub type if applicable like 'cbc' for AES-CBC"
	keyLengthHelp  = "length of the key for encryption in bits, e.g. 128 for AES-128"
	randomSeedHelp = "a decimal number between 0 and 1 to be used as seed in random number generation"

	inputText      = vexillum.String('i', "input-text", "input text to be encrypted", "")
	encryptionType = vexillum.String('t', "type", encryptionTypeHelp, "aes")
	subType        = vexillum.String('s', "sub-type", subTypeHelp, "")
	keyLength      = vexillum.Int('k', "key-length", keyLengthHelp, 128)
	inputFile      = vexillum.WildString("input-file", "the file to be encrypted", "")
	verbose        = vexillum.Bool('v', "verbose", "turn on verbose printing", true)
	randomSeed     = vexillum.Float64Validated('r', "random-seed", randomSeedHelp, 0.384526,
		func(d float64) error {
			if d <= 0 || d >= 1 {
				return fmt.Errorf("random seed must be between 0 and 1, both exclusive")
			}

			return nil
		})

	hashApp          = vexillum.NewApp("hash", "v0.0.1")
	hashAppAlgorithm = hashApp.String('a', "algorithm", "the algorithm for hashing", "md5")
	hashAppFile      = hashApp.WildString("file", "the file to be hashed", "")

	checksumApp          = vexillum.NewApp("checksum", "v0.0.3")
	checksumAppAlgorithm = checksumApp.String('a', "algorithm", "the algorithm", "md5")
	checksumAppFile      = checksumApp.WildString("file", "the file for checksum", "")
	checksumAppExpected  = checksumApp.WildString("expected", "the expected checksum", "")
)

func main() {
	vexillum.SetApp("encryptor")
	vexillum.SetVersion("v0.0.2")
	vexillum.Parse()

	switch vexillum.CurrentApp() {
	case hashApp:
		fmt.Printf("hash app is running\n")
		fmt.Printf("  algorithm: %s\n", *hashAppAlgorithm)
		fmt.Printf("  file: %s\n", *hashAppFile)
	case checksumApp:
		fmt.Printf("checksum app is running\n")
		fmt.Printf("  algorithm: %s\n", *checksumAppAlgorithm)
		fmt.Printf("  file: %s\n", *checksumAppFile)
		fmt.Printf("  expected: %s\n", *checksumAppExpected)
	default:
		fmt.Printf("encryptor app is running\n")
		fmt.Printf("  input text: %s\n", *inputText)
		fmt.Printf("  type: %s\n", *encryptionType)
		fmt.Printf("  sub type: %s\n", *subType)
		fmt.Printf("  key length: %d\n", *keyLength)
		fmt.Printf("  input file: %s\n", *inputFile)
		fmt.Printf("  verbose: %t\n", *verbose)
		fmt.Printf("  random seed: %f\n", *randomSeed)
	}

	fmt.Printf("remaining arguments: %+v\n", vexillum.Remaining())
}
```

---
output of `app-exe -h`:
```
encryptor v0.0.2
usage:
  cli> build_example.exe [named flags] [input-file]
  named flags:
    -h        --help: (type: boolean, default: false)
      show the help
    -i  --input-text: (type: string, default: "")
      input text to be encrypted
    -t        --type: (type: string, default: "aes")
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus et libero in nisl maximus
      hendrerit.
      Morbi vel dignissim neque.
      Cras ut nunc eget ante vulputate porttitor at scelerisque nisl. Suspendisse vestibulum
      mollis tortor. Suspendisse eu augue vestibulum, imperdiet metus accumsan, euismod enim.
      Quisque auctor dignissim ornare. Nulla ullamcorper, erat id sodales cursus, risus lorem
      congue orci, sed ultrices arcu orci et orci. In aliquet dapibus commodo. Sed a ligula nibh.
      Sed at velit vel odio maximus commodo ut a arcu. Aliquam sit amet sem est. Integer id mattis
      justo. Fusce nec porta erat, eget lobortis dolor. Integer non velit id ipsum aliquam luctus
      at ac diam. Donec maximus venenatis auctor.
    -s    --sub-type: (type: string, default: "")
      sub type if applicable like 'cbc' for AES-CBC
    -k  --key-length: (type: integer, default: 128)
      length of the key for encryption in bits, e.g. 128 for AES-128
    -v     --verbose: (type: boolean, default: false)
      turn on verbose printing
    -r --random-seed: (type: decimal, default: 0.384526)
      a decimal number between 0 and 1 to be used as seed in random number generation
  wild flags:
    [0] input-file: (type: string, default: "")
      the file to be encrypted
```

---
output of `app-exe hash -h`:
```
encryptor hash v0.0.1
usage:
  cli> build_example.exe [named flags] [file]
  named flags:
    -h      --help: (type: boolean, default: false)
      show the help
    -a --algorithm: (type: string, default: "md5")
      the algorithm for hashing
  wild flags:
    [0] file: (type: string, default: "")
      the file to be hashed
```

---
output of `app-exe checksum -h`:
```
encryptor checksum v0.0.3
usage:
  cli> build_example.exe [named flags] [file] [expected]
  named flags:
    -h      --help: (type: boolean, default: false)
      show the help
    -a --algorithm: (type: string, default: "md5")
      the algorithm
  wild flags:
    [0]     file: (type: string, default: "")
      the file for checksum
    [1] expected: (type: string, default: "")
      the expected checksum
```

---
output of `app-exe -i "Hello World!!!" file.txt -v -s cbc file2.png --key-length 256`:
```
encryptor app is running
  input text: Hello World!!!
  type: aes
  sub type: cbc
  key length: 256
  input file: file.txt
  verbose: true
  random seed: 0.384526
remaining arguments: [file2.png]
```
in the prior example `-r --random-seed` and `-t --type` is set to default because no value is provided for them.
you can see a warning log for that if you set `vexillum.ShowWarnings(true)` before `vexillum.Parse()`.

warnings which led flags fall back to their default values:
  - when a name flag is not referred.
  - when a name flag is referred but there is no value provided for it.
  - when the value provided for a flag is invalid and validation function is done with error.
errors which led application to exit with code 1:
  - when a name flag is used, but it is not defined in the app.

---
output of `app-exe -i "Hello World!!!" file.txt -v -s cbc file2.png -t --key-length 256` when warnings are shown:
```
flag warning: '-t --type' set to default because the value is missing
flag warning: '-r --random-seed' set to default because it's not referred
encryptor app is running
  input text: Hello World!!!
  type: aes
  sub type: cbc
  key length: 256
  input file: file.txt
  verbose: true
  random seed: 0.384526
remaining arguments: [file2.png]
```

---
output of `app-exe hash f1 f2 -a md4 f3 f4` without show warnings:
```
hash app is running
  algorithm: md4
  file: f1
remaining arguments: [f2 f3 f4]
```

---
short flags can be packed into groups like `-xyz`:
  - non-last short flags in a group can be only booleans, for example:
    - in `-xyz arg2`, flags `x` and `y` have to be booleans.
  - last short flag in a group can be any type, for example:
    - in `-xyz arg2`, `arg2` will be parsed as value for flag `z`.
    - in `-xyz --verbose`, because there is no value passed for `z` it should be boolean.