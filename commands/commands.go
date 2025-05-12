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
		"map": {
			name:        "map",
			description: "Areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Areas in the Pokemon world",
			Callback:    commandMapB,
		},
		"explore": {
			name: "explore",
			description: "get list of Pokemon located in area",
			Callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "catch a pokemon",
			Callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "inspect a pokemon to see details",
			Callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "display caught pokemon",
			Callback: commandPokedex,
		},
    }

	return commands
}