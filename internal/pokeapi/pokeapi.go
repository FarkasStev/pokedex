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
