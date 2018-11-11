package main

import (
	. "github.com/logrusorgru/aurora"

	"fmt"
	"strings"
)

func runInitCommands(config *Configuration, outputer Outputer) error {
	if err := prepInit(outputer); err != nil {
		return err
	}

	//prep replacer
	var macroReplacer *strings.Replacer
	{
		rs := make([]string, 0, len(commandMacros)*2)
		for macro, value := range commandMacros {
			rs = append(rs, macro, value)
		}
		macroReplacer = strings.NewReplacer(rs...)
	}

	for _, section := range []string{"common", currentOS, currentDistro, "after"} {
		if section == "" {
			continue
		}
		commands, has := config.Init[section]
		if !has {
			continue
		}
		outputer.OutInfo("--> Running [%s] init commands...", section)

	commandLoop:
		for _, command := range commands {
			if missing := missingMacros(command); missing != nil {
				outputer.OutWarn("command: %s refers to non existent macros %q, skipping", command, missing)
				continue commandLoop
			}
			command = macroReplacer.Replace(command)
			if dry {
				outputer.OutInfo("  %s would run: '%s'", Blue("❯"), command)
				continue
			}
			if err := runCommand(command, outputer); err != nil {
				return fmt.Errorf(`error running "%s": %s`, command, err)
			}
		}
	}
	return nil
}
