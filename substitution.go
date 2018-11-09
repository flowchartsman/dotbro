package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
)

var (
	currentOS     string
	currentDistro string
	currentHost   string
)

var (
	reGoodMacro     = regexp.MustCompile(`^[[:upper:]]+$`)
	reExtractMacros = regexp.MustCompile(`(%[[:upper:]]+%)`)
)

var (
	commandMacros = map[string]string{}
)

func registerCommandMacro(macroName string, value string) {
	if macroName == "" || !reGoodMacro.MatchString(macroName) {
		panic(fmt.Sprintf(`Invalid macro provided, name must match /%s/. Name given: "%s"`, reGoodMacro, macroName))
	}
	macroName = `%` + macroName + `%`
	if _, ok := commandMacros[macroName]; ok {
		panic("attempt to re-register macro " + macroName)
	}
	commandMacros[macroName] = value
}

func missingMacros(command string) []string {
	foundMacros := reExtractMacros.FindAllString(command, -1)
	missingMap := map[string]struct{}{}
	for _, macro := range foundMacros {
		if _, ok := commandMacros[macro]; !ok {
			missingMap[macro] = struct{}{}
		}
	}
	if len(missingMap) == 0 {
		return nil
	}
	missing := make([]string, 0, len(missingMap))
	for m := range missingMap {
		missing = append(missing, m)
	}
	return missing
}

func init() {
	var err error
	currentOS = runtime.GOOS
	currentHost, err = os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Unable to determine hostname: %s", err)
	}
	registerCommandMacro("OS", currentOS)
	registerCommandMacro("HOSTNAME", currentHost)
}
