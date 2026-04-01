package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_value_urls"`
	Chance          int                `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

func GetLocationArea(url string) (NamedAPIResourceList, error) {
	res, err := http.Get(url)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return NamedAPIResourceList{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return NamedAPIResourceList{}, fmt.Errorf("Response failed with error: %v", err)
	}

	namedAPIResources := NamedAPIResourceList{}
	err = json.Unmarshal(body, &namedAPIResources)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	return namedAPIResources, nil
}
