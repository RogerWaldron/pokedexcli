package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*ApiConfig) error
}