package pokeapi

// // Simple Types provided by the api docs

type ApiResource struct {
	Url string `json:"url"`
}

type Description struct {
	Description string           `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type Effect struct {
	Effect   string           `json:"effect"`
	Language NamedAPIResource `json:"language"`
}

type Encounter struct {
	Min_level        int                `json:"min_level"`
	Max_level        int                `json:"max_level"`
	Condition_values []NamedAPIResource `json:"condition_values"`
	Chance           int                `json:"chance"`
	Method           NamedAPIResource   `json:"method"`
}

type FlavorText struct {
	Flavor_text string           `json:"flavor_text"`
	Language    NamedAPIResource `json:"language"`
	Version     NamedAPIResource `json:"version"`
}

type GenerationGameIndex struct {
	Game_index int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type MachineVersionDetail struct {
	Machine       ApiResource      `json:"machine"`
	Version_group NamedAPIResource `json:"version_group"`
}

type Name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

// this generic struct is used as a field in various endpoint types via composition
type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VerboseEffect struct {
	Effect       string           `json:"effect"`
	Short_effect string           `json:"short_effect"`
	Language     NamedAPIResource `json:"language"`
}

type VersionEncounterDetail struct {
	Version           NamedAPIResource `json:"version"`
	Max_chance        int              `json:"max_chance"`
	Encounter_details []Encounter      `json:"encounter_details"`
}

type VersionGameIndex struct {
	Game_index int              `json:"game_index"`
	Version    NamedAPIResource `json:"version"`
}

type VersionGroupFlavorText struct {
	Text          string           `json:"text"`
	Language      NamedAPIResource `json:"language"`
	Version_group NamedAPIResource `json:"version_group"`
}

type PokemonAbility struct {
	Is_hidden bool             `json:"is_hidden"`
	Slot      int              `json:"slot"`
	Ability   NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonFormType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

type PokemonHeldItem struct {
	Item            NamedAPIResource         `json:"item"`
	Version_details []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonMove struct {
	Move                  NamedAPIResource     `json:"move"`
	Version_group_details []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	Move_learn_method NamedAPIResource `json:"move_learn_method"`
	VersionGroup      NamedAPIResource `json:"version_group"`
	Level_learned_at  int              `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat      NamedAPIResource `json:"stat"`
	Effort    int              `json:"effort"` // ev values
	Base_stat int              `json:"base_stat"`
}

type PokemonSprites struct {
	Front_default      string `json:"front_default"`
	Front_shiny        string `json:"front_shiny"`
	Front_female       string `json:"front_female"`
	Front_shiny_female string `json:"front_shiny_female"`
	Back_default       string `json:"back_default"`
	Back_shiny         string `json:"back_shiny"`
	Back_female        string `json:"back_female"`
	Back_shiny_female  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type TypePokemon struct {
	Slot    int              `json:"slot"`
	Pokemon NamedAPIResource `json:"pokemon"`
}

type TypeRelations struct {
	No_damage_to       []NamedAPIResource `json:"no_damage_to"`
	Half_damage_to     []NamedAPIResource `json:"half_damage_to"`
	Double_damage_to   []NamedAPIResource `json:"double_damage_to"`
	No_damage_from     []NamedAPIResource `json:"no_damage_from"`
	Half_damage_from   []NamedAPIResource `json:"half_damage_from"`
	Double_damage_from []NamedAPIResource `json:"double_damage_from"`
}

type TypeRelationsPast struct {
	Generation       NamedAPIResource `json:"generation"`
	Damage_relations TypeRelations    `json:"damage_relations"`
}

type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          `json:"encounter_method"`
	Version_details  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         `json:"pokemon"`
	Version_details []VersionEncounterDetail `json:"version_details"`
}

// // Complex response models

// this genric response is for all endpoints without a query, such as a name or ID
// its a paginated list of all resources for an endpoint
type NamedAPIResourceList struct {
	Count    int                `json:"count"`    // count of resources // S/N all json field names must be uppercase
	Next     *string            `json:"next"`     // next page
	Previous *string            `json:"previous"` // previous page
	Results  []NamedAPIResource `json:"results"`  // named resources provided @ the current page
}

// generic response for unnamed api resources
type UnnamedAPIResourceList struct {
	Count    int           `json:"count"`    // same as above
	Next     *string       `json:"next"`     // ...
	Previous *string       `json:"previous"` // ...
	Results  []ApiResource `json:"results"`  // unnamed resources provided @ the current page
}

// response provided by `location/{id or name}/`
type LocationIDResponse struct {
	Areas        []NamedAPIResource    `json:"areas"`
	Game_indices []GenerationGameIndex `json:"game_indices"`
	Id           int                   `json:"id"`
	Name         string                `json:"name"`
	Names        []Name                `json:"names"`
	Region       NamedAPIResource      `json:"region"`
}

// response from `location-area/{id or name}/`
type LocationAreaResponse struct {
	Id                     int                   `json:"id"`
	Name                   string                `json:"name"`
	Game_index             int                   `json:"game_index"`
	Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
	Location               NamedAPIResource      `json:"location"`
	Names                  []Name                `json:"names"`
	Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
}

type Pokemon struct {
	Id                       int                   `json:"id"`
	Name                     string                `json:"name"`
	Base_experience          int                   `json:"base_experience"`
	Height                   int                   `json:"height"`
	Is_default               bool                  `json:"is_default"` // Set for exactly one Pokémon used as the default for each species.
	Order                    int                   `json:"order"`      // Order for sorting. Almost national order, except families are grouped together.
	Weight                   int                   `json:"weight"`
	Abilities                []PokemonAbility      `json:"abilities"`
	Forms                    []NamedAPIResource    `json:"forms"`        // A list of forms this Pokémon can take on.
	Game_indices             []GenerationGameIndex `json:"game_indices"` // A list of game indices relevent to Pokémon item by generation.
	Held_items               []PokemonHeldItem     `json:"held_items"`
	Location_area_encounters string                `json:"location_area_encounters"`
	Moves                    []PokemonMove         `json:"moves"`
	Past_types               []PokemonTypePast     `json:"past_types"` // A list of details showing types this pokémon had in previous generations
	Sprites                  PokemonSprites        `json:"sprites"`    // A set of sprites used to depict this Pokémon in the game.
	Cries                    PokemonCries          `json:"cries"`
	Species                  NamedAPIResource      `json:"species"`
	Stats                    []PokemonStat         `json:"stats"`
	Types                    []PokemonType         `json:"types"`
}

type Type struct {
	Id                    int                   `json:"id"`
	Name                  string                `json:"name"`
	Damage_relations      TypeRelations         `json:"damage_relations"`
	Past_damage_relations []TypeRelationsPast   `json:"past_damage_relations"`
	Game_indices          []GenerationGameIndex `json:"game_indices"`
	Generation            NamedAPIResource      `json:"generation"`
	Move_damage_class     NamedAPIResource      `json:"move_damage_class"`
	Names                 []Name                `json:"names"`
	Pokemon               []TypePokemon         `json:"pokemon"`
	Moves                 []NamedAPIResource    `json:"moves"`
}
