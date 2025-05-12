package commands

import (
	"fmt"
)

func commandInspect(cfg *ApiConfig) error {
	if len(cfg.Text) < 2 {
		return nil
	}

	caught, ok := cfg.Captured[cfg.Text[1]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	
	fmt.Println("Name: ", caught.Name)
	fmt.Println("Height: ", caught.Height)
	fmt.Println("Weight: ", caught.Weight)
	fmt.Println("Stats:")
	for _, item := range caught.Stats {
		fmt.Printf("-%s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range caught.Types {
		fmt.Printf("- %s\n", item.Type.Name)
	}

	return nil
}