package commands

func GetCommands() map[string]cliCommand {
	commands := map[string]cliCommand {
        "help":  {
			name: "Help",
			description: "Help",
			Callback: commandHelp,
		},
        "clear": {
			name: "Clear Screen",
			description: "Clear the screen",
			Callback: commandClear,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"mapnext": {
			name:        "map",
			description: "Areas in the Pokemon world",
			Callback:    commandMapNext,
		},
		"mapprev": {
			name:        "map",
			description: "Areas in the Pokemon world",
			Callback:    commandMapPrev,
		},
    }

	return commands
}