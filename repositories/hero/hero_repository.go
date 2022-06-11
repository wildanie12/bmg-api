package hero

import (
	"encoding/json"
	"gin-bmg-restful/entities/web"
	"net/http"
	"strings"
)

// Repository main struct
type Repository struct {
	baseURL string
	client *http.Client
}

// NewRepository constructor
func NewRepository() *Repository {
	return &Repository{
		baseURL: "https://ddragon.leagueoflegends.com/cdn/6.24.1/data/en_US/champion.json",
		client: &http.Client{},
	}
}

// FindAllHeroes returns list of hero correlated with provided
// nameSearch parameter
func (repo Repository) FindAllHeroes(nameSearch string) ([]web.Hero, error) {

	request, err := http.NewRequest(http.MethodGet, repo.baseURL, nil)
	if err != nil {
		return []web.Hero{}, web.Error{ Code: 500, Message: "Server Error: cannot create request" }
	}
	response, err := repo.client.Do(request)
	if err != nil {
		return []web.Hero{}, web.Error{ Code: 500, Message: "Server Error: request failed" }
	}

	heroAPIResponse := web.HeroAPIResponse{}
	err = json.NewDecoder(response.Body).Decode(&heroAPIResponse)
	if err != nil {
		return []web.Hero{}, web.Error{ Code: 500, Message: "Server Error: response parsing failed" }
	}
	heroes := []web.Hero{}
	for _, hero := range heroAPIResponse.Data {
		if strings.Contains(hero.Name, nameSearch) {
			heroes = append(heroes, hero)
		}
	}
	return heroes, nil
}