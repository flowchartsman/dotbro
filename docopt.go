package main

import "github.com/docopt/docopt-go"

const version = "0.2.0"

// ParseArguments parses arguments, that were passed to the dotbro, by docopt.
func ParseArguments(argv []string) (map[string]interface{}, error) {
	usage := `dotbro - simple yet effective dotfiles manager.

Usage:
  dotbro [options]
  dotbro init [options]
  dotbro add [options] <filename>
  dotbro -h | --help
  dotbro --version

Common options:
  -c --config=<filepath>  Dotbro's configuration file in JSON or TOML format.
						  [default: dotbro.toml]
  -q --quiet              Quiet mode. Do not print any output, except warnings
                          and errors.
  -d --dry                Dry mode. Do not take any action, instead print what
						  WOULD happen (suppresses --quiet).
  -v --verbose            Verbose mode. Detailed output.

Add options:
  <filename>              File to add.

Other options:
  -h --help               Show this helpful info.
  -V --version            Show version.
`

	return docopt.Parse(usage, argv, true, "dotbro "+version, false)
}
