package commands

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	Next     string
	Previous string
}