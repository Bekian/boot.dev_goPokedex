package main

import (
	"bufio"
	"fmt"
	"os"
	pokecache "pokedex/pokeCache"
	"pokedex/pokeapi"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      config
}

type config struct {
	option string
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// command wrapper for the GetLocation function
func commandMap(pageCount *int) error {
	locations := pokeapi.GetLocation(pageCount)
	if locations.Count <= 0 {
		fmt.Println("No location with the given page found for page: ", *pageCount)
		// should return error here
		return nil
	}
	fmt.Println(locations)
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

var cmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		config:      config{"none"},
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandExit, // this isnt used
		config:      config{"none"},
	},
	"map": {
		name:        "map",
		description: "Displays the next 20 names of location areas in the Pokemon world.",
		callback:    commandExit,
		config:      config{""},
	},
	"mapb": {
		name:        "map",
		description: "Displays the previous 20 names of location areas in the Pokemon world.",
		callback:    commandExit,
		config:      config{""},
	},
}

func main() {
	cache := pokecache.NewCache(5)
	cache.Add("hello", []byte{})
	mapCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) <= 0 {
			cmds["exit"].callback()
		}
		command := strings.Fields(strings.ToLower(input))

		fmt.Println("Your command was :", command[0])

		switch command[0] {
		case "exit", "\n":
			err := cmds["exit"].callback()
			Check(err)
		case "help":
			for _, cmd := range cmds {
				fmt.Println(cmd.name, ":", cmd.description)
			}
		case "map":
			commandMap(&mapCount)
			mapCount++
		case "mapb":
			mapCount--
			mapCount--
			if mapCount >= 0 {
				commandMap(&mapCount)
			} else {
				fmt.Println("No previous locations available")
				mapCount++
				mapCount++
			}
		default:
			fmt.Println("Unknown command: ", command[0])
		}

	}

}
