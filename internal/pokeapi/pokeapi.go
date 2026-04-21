package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/farkasstev/pokedex/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
	}
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates,omitempty"`
	GameIndex            int                    `json:"game_index,omitempty"`
	Id                   int                    `json:"id,omitempty"`
	Location             Location               `json:"location,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Names                []Names                `json:"names,omitempty"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters,omitempty"`
}

type EncounterMethod struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Version struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionDetails struct {
	Rate    int     `json:"rate,omitempty"`
	Version Version `json:"version,omitempty"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method,omitempty"`
	VersionDetails  []VersionDetails `json:"version_details,omitempty"`
}

type Location struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Language struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Names struct {
	Language Language `json:"language,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type Pokemon struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type PokemonInfo struct {
	Abilities              []Abilities     `json:"abilities"`
	BaseExperience         int             `json:"base_experience"`
	Cries                  Cries           `json:"cries"`
	Forms                  []Forms         `json:"forms"`
	GameIndices            []GameIndices   `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              []HeldItems     `json:"held_items"`
	ID                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []Moves         `json:"moves"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	PastAbilities          []PastAbilities `json:"past_abilities"`
	PastStats              []PastStats     `json:"past_stats"`
	PastTypes              []any           `json:"past_types"`
	Species                Species         `json:"species"`
	Sprites                Sprites         `json:"sprites"`
	Stats                  []Stats         `json:"stats"`
	Types                  []Types         `json:"types"`
	Weight                 int             `json:"weight"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type HeldItems struct {
	Item           Item             `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type Move struct {
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
type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	Order           any             `json:"order"`
	VersionGroup    VersionGroup    `json:"version_group"`
}
type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PastAbilities struct {
	Abilities  []Abilities `json:"abilities"`
	Generation Generation  `json:"generation"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type PastStats struct {
	Generation Generation `json:"generation"`
	Stats      []Stats    `json:"stats"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type Home struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Showdown struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}
type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type Yellow struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}
type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}
type Gold struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type Silver struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}
type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}
type DiamondPearl struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type HeartgoldSoulsilver struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type Platinum struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}
type ScarletViolet struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type GenerationIx struct {
	ScarletViolet ScarletViolet `json:"scarlet-violet"`
}
type Animated struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type BlackWhite struct {
	Animated         Animated `json:"animated"`
	BackDefault      string   `json:"back_default"`
	BackFemale       string   `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  string   `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      string   `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale string   `json:"front_shiny_female"`
}
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}
type OmegarubyAlphasapphire struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type XY struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}
type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type UltraSunUltraMoon struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}
type BrilliantDiamondShiningPearl struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type GenerationViii struct {
	BrilliantDiamondShiningPearl BrilliantDiamondShiningPearl `json:"brilliant-diamond-shining-pearl"`
	Icons                        Icons                        `json:"icons"`
}
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationIx   GenerationIx   `json:"generation-ix"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}
type Sprites struct {
	BackDefault      string   `json:"back_default"`
	BackFemale       string   `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  string   `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      string   `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale string   `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

func (c *Client) GetLocationAreas(url string) (NamedAPIResourceList, error) {
	bytes, present := c.cache.Get(url)
	if !present {
		res, err := http.Get(url)
		if err != nil {
			return NamedAPIResourceList{}, err
		}

		bytes, err = io.ReadAll(res.Body)

		defer res.Body.Close()

		if res.StatusCode > 299 {
			return NamedAPIResourceList{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, bytes)
		}
		if err != nil {
			return NamedAPIResourceList{}, fmt.Errorf("Response failed with error: %v", err)
		}
		c.cache.Add(url, bytes)
	}
	namedAPIResources := NamedAPIResourceList{}
	err := json.Unmarshal(bytes, &namedAPIResources)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	return namedAPIResources, nil
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name
	bytes, present := c.cache.Get(url)
	if !present {
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}

		bytes, err = io.ReadAll(res.Body)

		defer res.Body.Close()

		if res.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, bytes)
		}
		if err != nil {
			return LocationArea{}, fmt.Errorf("Response failed with error: %v", err)
		}
		c.cache.Add(url, bytes)
	}

	locationArea := LocationArea{}
	err := json.Unmarshal(bytes, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}

func (c *Client) GetPokemonInfo(name string) (PokemonInfo, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	bytes, present := c.cache.Get(url)
	if !present {
		res, err := http.Get(url)
		if err != nil {
			return PokemonInfo{}, err
		}

		bytes, err = io.ReadAll(res.Body)

		defer res.Body.Close()

		if res.StatusCode > 299 {
			return PokemonInfo{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, bytes)
		}
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("Response failed with error: %v", err)
		}
		c.cache.Add(url, bytes)
	}

	pokemonInfo := PokemonInfo{}
	err := json.Unmarshal(bytes, &pokemonInfo)
	if err != nil {
		return pokemonInfo, err
	}

	return pokemonInfo, nil

}
