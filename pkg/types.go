// This package implements an http client for resource and pokemon APIs from pokeAPI
package pokemon

// Resource is the resources available for an endpoint
type Resource struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

// Result is a single API resource
type Result struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

// Pokemon holds a single pokemon data
type Pokemon struct {
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	BaseExperience         int               `json:"base_experience"`
	Height                 int               `json:"height"`
	IsDefault              bool              `json:"is_default"`
	Order                  int               `json:"order"`
	Weight                 int               `json:"weight"`
	Abilities              []PokemonAbility  `json:"abilities"`
	Forms                  []Form            `json:"forms"`
	GameIndices            []GameIndex       `json:"game_indices"`
	HeldItems              []PokemonHeldItem `json:"held_items"`
	LocationAreaEncounters string            `json:"location_area_encounters"`
	Moves                  []PokemonMove     `json:"moves"`
	PastTypes              []PokemonTypePast `json:"past_types"`
	Sprites                PokemonSprites    `json:"sprites"`
	Cries                  PokemonCries      `json:"cries"`
	Species                Species           `json:"species"`
	Stats                  []PokemonStat     `json:"stats"`
	Types                  []PokemonType     `json:"types"`
}

// PokemonAbility is a single pokemon ability
type PokemonAbility struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}

// Ability is an API resource
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Form is a named API resource
type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GameIndex is a single game index
type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

// Version is a named API resource
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Item is a named API resource
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonHeldItem is a single Pokemon item
type PokemonHeldItem struct {
	Item           Item                   `json:"item"`
	VersionDetails PokemonHeldItemVersion `json:"version_details"`
}

// PokemonHeldItemVersion hold the information about the held item version
type PokemonHeldItemVersion struct {
	Version Version `json:"version"`
	Rarity  int     `json:"rarity"`
}

// MoveLearnMethod is a named API resource
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VersionGroup is a named API resource
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonMoveVersion holds the details about the pokemon move versions
type PokemonMoveVersion struct {
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
	LevelLearnedAt  int             `json:"level_learned_at"`
}

// Move is a named API resource
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonMove is a single pokemon move
type PokemonMove struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

// Type is a named API resource
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonType is a single pokemon type
type PokemonType struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// Generation is a named API resource
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonTypePast is a single pokemon past type
type PokemonTypePast struct {
	Generation Generation    `json:"generation"`
	Types      []PokemonType `json:"types"`
}

// PokemonSprites holds the details about single pokemon sprites
type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

// PokemonCries holds the details about single pokemon cries
type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// Species is a named API resource
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Stat is a named API resource
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonStat holds the details about single pokemon stats
type PokemonStat struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
}
