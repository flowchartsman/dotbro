package main

import (
	"strings"

	. "github.com/logrusorgru/aurora"
)

func runInitCommands(config *Configuration, outputer Outputer) error {
	if err := prepInit(outputer); err != nil {
		return err
	}
	for _, section := range []string{"common", currentOS, "after"} {
		commands, has := config.Init[section]
		if !has {
			return nil
		}
		outputer.OutInfo("--> Running [%s] init commands...", section)

		for _, command := range commands {
			// replace any occurrence of %DOTFILEDIR with dotfile directory
			command = strings.Replace(command, "%DOTFILEDIR", config.Directories.Dotfiles, -1)
			if dry {
				outputer.OutInfo("  %s would run: '%s'", Blue("❯"), command)
				continue
			}
			if err := runCommand(command, outputer); err != nil {
				//maybe?
				return err
			}
		}
	}
	return nil
}
