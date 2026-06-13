package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type parsedInfo struct {
	command string
	params  []string
	flags   map[string]string
}

var validCommands = [...]string{"add"}

func parse(args []string) (parsedInfo, error) {
	command := args[0]
	isValidCommand := slices.Contains(validCommands[:], command)
	if !isValidCommand {
		return parsedInfo{}, fmt.Errorf("Command not valid")
	}

	var params []string
	var flags = make(map[string]string)

	args = args[1:]
	for i := 0; i < len(args); i++ {
		word := args[i]

		if strings.TrimSpace(word) == "" {
			return parsedInfo{}, fmt.Errorf("Expecting value, recieved empty")
		}

		if word[0] == '-' {
			if _, isInFlagMap := flags[word]; isInFlagMap {
				return parsedInfo{}, fmt.Errorf("Duplicated flag %s", word)
			}

			if i+1 > len(args)-1 { // ending command input with flag and no value
				return parsedInfo{}, fmt.Errorf("Wrong flag use")
			}

			flagValue := args[i+1]

			if strings.TrimSpace(flagValue) == "" {
				return parsedInfo{}, fmt.Errorf("Expecting value, recieved empty")
			}

			if flagValue[0] == '-' {
				return parsedInfo{}, fmt.Errorf("Wrong flag use")
			}

			flags[word] = flagValue
			i++
			continue
		}

		params = append(params, word)
	}

	return parsedInfo{command: command, params: params, flags: flags}, nil
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("\n                         YACT")
		fmt.Println("                 Yet Another CLI Todo")
		return
	}

	data, err := parse(args[1:])

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
