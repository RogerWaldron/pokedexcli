package commands

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand {
        "help":  {
			name: "Help",
			description: "Help",
			callback: commandHelp,
		},
        "clear": {
			name: "Clear Screen",
			description: "Clear the screen",
			callback: commandClear,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Areas in the Pokemon world",
			callback:    commandMap,
		},
    }

	return commands
}