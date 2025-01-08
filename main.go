package main

import (
	"bufio"
	"fmt"
	"os"
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
	cmd         string
	description string
	callback    func() error
	config      config
}

type config struct {
	option string
}

func noArgProvided(args []string, msg string) bool {
	if len(args) < 2 {
		fmt.Printf("\nNo argument \"%s\" provided, try again\n", msg)
		return true
	}
	return false
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

// wrapper for the Pokemon function
func commandPokemon(query string) error {
	pokemon := pokeapi.QueryPokemon(query)
	if pokemon.Id == 0 {
		return nil // pokemon not found
	}
	fmt.Println("pokemon found:", pokemon.Name, " ID:", pokemon.Id) // might need some other formatting
	return nil
}

// wrapper for the types function
func commandTypes(query string) error {
	types := pokeapi.QueryTypes(query)
	if types.Id == 0 {
		return nil // type not found
	}
	fmt.Println("type found:", types.Name, " ID:", types.Id, " Types:", types.Damage_relations)
	return nil
}

func commandLocation(query string) error {
	location := pokeapi.GetLocationAreaName(&query)
	if location.Id == 0 {
		return nil // location not found
	}
	fmt.Println("Location found:", location.Name, " ID:", location.Id)
	return nil
}

func handleMap(back bool, mapCount *int) {
	mapNumber := *mapCount
	if !back {
		mapNumber--
		mapNumber--
		if mapNumber >= 0 {
			commandMap(&mapNumber)
		} else {
			fmt.Println("No previous locations available")
			mapNumber++
			mapNumber++
		}
		return
	}
	commandMap(&mapNumber)
	mapNumber++
}

var cmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		cmd:         "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		config:      config{"none"},
	},
	"help": {
		name:        "help",
		cmd:         "help",
		description: "Displays a help message",
		callback:    commandExit, // this isnt used
		config:      config{"none"},
	},
	"map": {
		name:        "map",
		cmd:         "map",
		description: "Displays the next 20 names of location areas in the Pokemon world.",
		callback:    commandExit,
		config:      config{""},
	},
	"mapb": {
		name:        "map back",
		cmd:         "mapb",
		description: "Displays the previous 20 names of location areas in the Pokemon world.",
		callback:    commandExit,
		config:      config{""},
	},
	"loc": {
		name:        "locaction",
		cmd:         "loc",
		description: "Displays the information of a location, given by a location's name or ID number",
		callback:    commandExit,
		config:      config{""},
	},
	"pokemon": {
		name:        "pokemon",
		cmd:         "pok",
		description: "Displays the name and ID number of a pokemon, given by a pokemons name or ID number",
		callback:    commandExit, // not used
		config:      config{""},  // not used
	},
	"types": {
		name:        "types",
		cmd:         "types",
		description: "Displays the type relations for a given type or type ID",
		callback:    commandExit, // not used
		config:      config{""},  // not used
	},
}

func main() {
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
				fmt.Printf("-%s: %s usage - %s", cmd.name, cmd.description, cmd.cmd)
			}
		case "map":
			handleMap(false, &mapCount)
		case "mapb":
			handleMap(true, &mapCount)
		case "loc":
			if noArgProvided(command, cmds["location"].name) {
				break
			}
			commandLocation(command[1])
		case "pokemon":
			if noArgProvided(command, cmds["pokemon"].name) {
				break
			}
			commandPokemon(command[1])
		case "types":
			if noArgProvided(command, cmds["types"].name) {
				break
			}
			commandTypes(command[1])
		default:
			fmt.Println("Unknown command: ", command[0])
		}

	}

}
