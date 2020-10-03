package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2"

var cachedDesc = make(map[string]string)

type pokemonInfo struct {
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
}

type speciesInfo struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
}

func getAPI(resource string) (*http.Response, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", apiurl, resource))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Pokeapi respone code %v: %s", resp.StatusCode, body)
	}
	return resp, nil
}

func getPokemon(name string) (*pokemonInfo, error) {
	resp, err := getAPI(fmt.Sprintf("pokemon/%s", name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var info pokemonInfo
	json.NewDecoder(resp.Body).Decode(&info)
	return &info, nil
}

func getPokemonSpecies(name string) (*speciesInfo, error) {
	resp, err := getAPI(fmt.Sprintf("pokemon-species/%s", name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var info speciesInfo
	json.NewDecoder(resp.Body).Decode(&info)
	return &info, nil
}

// GetPokemonDesc returns description of pokemon
func GetPokemonDesc(name string) (string, error) {
	// This API provided for free, so trying not to use to much
	if d, ok := cachedDesc[name]; ok {
		return d, nil
	}

	p, err := getPokemon(name)
	if err != nil {
		return "", fmt.Errorf("GetPokemonDesc pokemon %v error: %w", name, err)
	}

	ps, err := getPokemonSpecies(p.Species.Name)
	if err != nil {
		return "", fmt.Errorf("GetPokemonDesc pokemon species %v error: %w", p.Species.Name, err)
	}
	for _, v := range ps.FlavorTextEntries {
		if v.Language.Name == "en" {
			cachedDesc[name] = v.FlavorText
			return v.FlavorText, nil
		}
	}
	return "", fmt.Errorf("GetPokemonDesc failed to get english description of %v", p.Species.Name)
}
