package pokemon

type Resource struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}
type Result struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

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
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonAbility struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}
type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type PokemonHeldItem struct {
	Item           Item                   `json:"item"`
	VersionDetails PokemonHeldItemVersion `json:"version_details"`
}
type PokemonHeldItemVersion struct {
	Version Version `json:"version"`
	Rarity  int     `json:"rarity"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonMoveVersion struct {
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
	LevelLearnedAt  int             `json:"level_learned_at"`
}
type PokemonMove struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonTypePast struct {
	Generation Generation    `json:"generation"`
	Types      []PokemonType `json:"types"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
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
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonStat struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonType struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
